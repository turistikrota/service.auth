package command

import (
	"context"

	"api.turistikrota.com/auth/src/config"
	"api.turistikrota.com/auth/src/domain/account"
	"api.turistikrota.com/auth/src/domain/owner"
	"api.turistikrota.com/auth/src/domain/user"
	"api.turistikrota.com/shared/auth/session"
	"api.turistikrota.com/shared/auth/token"
	"api.turistikrota.com/shared/decorator"
	"api.turistikrota.com/shared/events"
	"api.turistikrota.com/shared/jwt"
	"github.com/mixarchitecture/i18np"
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
	repo        user.Repository
	accountRepo account.Repository
	ownerRepo   owner.Repository
	authTopics  config.AuthTopics
	publisher   events.Publisher
	errors      user.Errors
	tokenSrv    token.Service
	sessionSrv  session.Service
}

type LoginVerifiedHandlerConfig struct {
	Repo        user.Repository
	AccountRepo account.Repository
	OwnerRepo   owner.Repository
	AuthTopics  config.AuthTopics
	Publisher   events.Publisher
	TokenSrv    token.Service
	SessionSrv  session.Service
	Errors      user.Errors
	CqrsBase    decorator.Base
}

func NewLoginVerifiedHandler(config LoginVerifiedHandlerConfig) LoginVerifiedHandler {
	return decorator.ApplyCommandDecorators[LoginVerifiedCommand, *LoginVerifiedResult](
		loginVerifiedHandler{
			repo:        config.Repo,
			accountRepo: config.AccountRepo,
			ownerRepo:   config.OwnerRepo,
			authTopics:  config.AuthTopics,
			publisher:   config.Publisher,
			errors:      config.Errors,
			tokenSrv:    config.TokenSrv,
			sessionSrv:  config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h loginVerifiedHandler) Handle(ctx context.Context, cmd LoginVerifiedCommand) (*LoginVerifiedResult, *i18np.Error) {
	u, err := h.repo.GetByUUID(ctx, cmd.UserUUID)
	if err != nil {
		return nil, h.errors.Failed(cmd.UserUUID)
	}
	accounts, owner, error := h.getUserRelations(ctx, u.UUID)
	if error != nil {
		return nil, error
	}
	tokens, error := h.sessionSrv.New(session.NewCommand{
		UserUUID:   u.UUID,
		DeviceUUID: cmd.DeviceUUID,
		Device:     cmd.Device,
		User: &session.SessionUser{
			UUID:     u.UUID,
			Email:    u.Email,
			Roles:    u.Roles,
			Accounts: accounts,
			Owners:   owner,
		},
	})
	if error != nil {
		return nil, h.errors.Failed("token")
	}
	_ = h.publisher.Publish(h.authTopics.LoggedIn, u)
	return &LoginVerifiedResult{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h loginVerifiedHandler) getUserRelations(ctx context.Context, userUUID string) ([]jwt.UserClaimAccount, []jwt.UserClaimOwner, *i18np.Error) {
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
