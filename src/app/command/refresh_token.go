package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.auth/src/config"
	"github.com/turistikrota/service.auth/src/domain/account"
	"github.com/turistikrota/service.auth/src/domain/owner"
	"github.com/turistikrota/service.auth/src/domain/user"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/jwt"
)

type RefreshTokenCommand struct {
	RefreshToken string
	AccessToken  string
	DeviceUUID   string
	UserUUID     string
	IpAddress    string
}

type RefreshTokenResult struct {
	AccessToken  string
	RefreshToken string
}

type RefreshTokenHandler decorator.CommandHandler[RefreshTokenCommand, *RefreshTokenResult]

type refreshTokenHandler struct {
	userRepo    user.Repository
	accountRepo account.Repository
	ownerRepo   owner.Repository
	authTopics  config.AuthTopics
	publisher   events.Publisher
	tokenSrv    token.Service
	sessionSrv  session.Service
	errors      user.Errors
}

type RefreshTokenHandlerConfig struct {
	UserRepo    user.Repository
	AccountRepo account.Repository
	OwnerRepo   owner.Repository
	AuthTopics  config.AuthTopics
	Publisher   events.Publisher
	TokenSrv    token.Service
	SessionSrv  session.Service
	Errors      user.Errors
	CqrsBase    decorator.Base
}

func NewRefreshTokenHandler(config RefreshTokenHandlerConfig) RefreshTokenHandler {
	return decorator.ApplyCommandDecorators[RefreshTokenCommand, *RefreshTokenResult](
		refreshTokenHandler{
			userRepo:    config.UserRepo,
			accountRepo: config.AccountRepo,
			ownerRepo:   config.OwnerRepo,
			authTopics:  config.AuthTopics,
			publisher:   config.Publisher,
			tokenSrv:    config.TokenSrv,
			sessionSrv:  config.SessionSrv,
			errors:      config.Errors,
		},
		config.CqrsBase,
	)
}

func (h refreshTokenHandler) Handle(ctx context.Context, cmd RefreshTokenCommand) (*RefreshTokenResult, *i18np.Error) {
	available := h.sessionSrv.IsRefreshAvailable(session.IsRefreshAvailableCommand{
		UserUUID:     cmd.UserUUID,
		DeviceUUID:   cmd.DeviceUUID,
		AccessToken:  cmd.AccessToken,
		RefreshToken: cmd.RefreshToken,
	})
	if !available {
		return nil, h.errors.RefreshTokenNotAvailable()
	}
	user, err := h.userRepo.GetByUUID(ctx, cmd.UserUUID)
	if err != nil {
		return nil, h.errors.Failed(err.Error())
	}
	accounts, owner, error := h.getUserRelations(ctx, user.UUID)
	if error != nil {
		return nil, error
	}
	ses := &session.SessionUser{
		UUID:     user.UUID,
		Email:    user.Email,
		Roles:    user.Roles,
		Accounts: accounts,
		Owners:   owner,
	}
	tokens, err := h.sessionSrv.Refresh(session.RefreshCommand{
		UserUUID:     cmd.UserUUID,
		DeviceUUID:   cmd.DeviceUUID,
		RefreshToken: cmd.RefreshToken,
		AccessToken:  cmd.AccessToken,
		User:         ses,
		IpAddress:    cmd.IpAddress,
	})
	if err != nil {
		return nil, h.errors.Failed(err.Error())
	}
	return &RefreshTokenResult{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h refreshTokenHandler) getUserRelations(ctx context.Context, userUUID string) ([]jwt.UserClaimAccount, []jwt.UserClaimOwner, *i18np.Error) {
	accounts, err := h.accountRepo.ListAsClaims(ctx, userUUID)
	if err != nil {
		return nil, nil, err
	}
	owners, err := h.ownerRepo.GetAllAsClaim(ctx, userUUID)
	if err != nil {
		return nil, nil, err
	}
	return accounts, owners, nil
}
