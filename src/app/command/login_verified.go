package command

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.auth/src/config"
	"github.com/turistikrota/service.auth/src/domain/account"
	"github.com/turistikrota/service.auth/src/domain/business"
	"github.com/turistikrota/service.auth/src/domain/user"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/jwt"
)

type LoginVerifiedCommand struct {
	UserUUID   string
	DeviceUUID string
	Device     *session.Device
}

type LoginVerifiedResult struct {
	AccessToken  string
	RefreshToken string
}

type LoginVerifiedHandler decorator.CommandHandler[LoginVerifiedCommand, *LoginVerifiedResult]

type loginVerifiedHandler struct {
	repo         user.Repository
	accountRepo  account.Repository
	businessRepo business.Repository
	authTopics   config.AuthTopics
	publisher    events.Publisher
	errors       user.Errors
	tokenSrv     token.Service
	sessionSrv   session.Service
}

type LoginVerifiedHandlerConfig struct {
	Repo         user.Repository
	AccountRepo  account.Repository
	BusinessRepo business.Repository
	AuthTopics   config.AuthTopics
	Publisher    events.Publisher
	TokenSrv     token.Service
	SessionSrv   session.Service
	Errors       user.Errors
	CqrsBase     decorator.Base
}

func NewLoginVerifiedHandler(config LoginVerifiedHandlerConfig) LoginVerifiedHandler {
	return decorator.ApplyCommandDecorators[LoginVerifiedCommand, *LoginVerifiedResult](
		loginVerifiedHandler{
			repo:         config.Repo,
			accountRepo:  config.AccountRepo,
			businessRepo: config.BusinessRepo,
			authTopics:   config.AuthTopics,
			publisher:    config.Publisher,
			errors:       config.Errors,
			tokenSrv:     config.TokenSrv,
			sessionSrv:   config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h loginVerifiedHandler) Handle(ctx context.Context, cmd LoginVerifiedCommand) (*LoginVerifiedResult, *i18np.Error) {
	u, err := h.repo.GetByUUID(ctx, cmd.UserUUID)
	if err != nil {
		return nil, h.errors.Failed(cmd.UserUUID)
	}
	accounts, bus, error := h.getUserRelations(ctx, u.UUID)
	if error != nil {
		return nil, error
	}
	tokens, error := h.sessionSrv.New(session.NewCommand{
		UserUUID:   u.UUID,
		DeviceUUID: cmd.DeviceUUID,
		Device:     cmd.Device,
		User: &session.SessionUser{
			UUID:       u.UUID,
			Email:      u.Email,
			Roles:      u.Roles,
			Accounts:   accounts,
			Businesses: bus,
		},
	})
	if error != nil {
		return nil, h.errors.Failed("token")
	}
	if u.IsDeleted && time.Since(u.DeletedAt) > 30*time.Hour*24 {
		error = h.repo.Recover(ctx, u.UUID)
		if error != nil {
			return nil, error
		}
	}
	_ = h.publisher.Publish(h.authTopics.LoggedIn, u)
	return &LoginVerifiedResult{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h loginVerifiedHandler) getUserRelations(ctx context.Context, userUUID string) ([]jwt.UserClaimAccount, []jwt.UserClaimBusiness, *i18np.Error) {
	accounts, err := h.accountRepo.ListAsClaims(ctx, userUUID)
	if err != nil {
		return nil, nil, err
	}
	businesses, err := h.businessRepo.GetAllAsClaim(ctx, userUUID)
	if err != nil {
		return nil, nil, err
	}
	return accounts, businesses, nil
}
