package command

import (
	"context"

	"api.turistikrota.com/auth/src/config"
	"api.turistikrota.com/auth/src/domain/account"
	"api.turistikrota.com/auth/src/domain/owner"
	"api.turistikrota.com/auth/src/domain/user"
	"api.turistikrota.com/shared/auth/session"
	"api.turistikrota.com/shared/auth/token"
	"api.turistikrota.com/shared/cipher"
	"api.turistikrota.com/shared/decorator"
	"api.turistikrota.com/shared/events"
	"api.turistikrota.com/shared/helper"
	"api.turistikrota.com/shared/jwt"
	"github.com/mixarchitecture/i18np"
)

type LoginCommand struct {
	Email      string
	Password   string
	DeviceUUID string
	Device     *session.Device
}

type LoginResult struct {
	AccessToken  string
	RefreshToken string
	Verify       bool
	TempToken    string
}

type VerifyRequiredError struct {
	VerifyRequired bool `json:"verifyRequired"`
}

type LoginHandler decorator.CommandHandler[LoginCommand, *LoginResult]

type loginHandler struct {
	userRepo     user.Repository
	accountRepo  account.Repository
	ownerRepo    owner.Repository
	authTopics   config.AuthTopics
	verifyTopics config.VerifyTopics
	publisher    events.Publisher
	errors       user.Errors
	tokenSrv     token.Service
	sessionSrv   session.Service
}

type LoginHandlerConfig struct {
	UserRepo     user.Repository
	AccountRepo  account.Repository
	OwnerRepo    owner.Repository
	AuthTopics   config.AuthTopics
	VerifyTopics config.VerifyTopics
	Publisher    events.Publisher
	TokenSrv     token.Service
	SessionSrv   session.Service
	Errors       user.Errors
	CqrsBase     decorator.Base
}

func NewLoginHandler(config LoginHandlerConfig) LoginHandler {
	return decorator.ApplyCommandDecorators[LoginCommand, *LoginResult](
		loginHandler{
			userRepo:     config.UserRepo,
			accountRepo:  config.AccountRepo,
			ownerRepo:    config.OwnerRepo,
			authTopics:   config.AuthTopics,
			verifyTopics: config.VerifyTopics,
			publisher:    config.Publisher,
			errors:       config.Errors,
			tokenSrv:     config.TokenSrv,
			sessionSrv:   config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h loginHandler) Handle(ctx context.Context, cmd LoginCommand) (*LoginResult, *i18np.Error) {
	user, err := h.userRepo.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsVerified {
		return nil, h.errors.NotVerified(VerifyRequiredError{
			VerifyRequired: true,
		})
	}
	if err := cipher.Compare(user.Password, cmd.Password); err != nil {
		_ = h.publisher.Publish(h.authTopics.LoginFailed, user)
		return nil, h.errors.InvalidPassword()
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
	if user.TwoFactorEnabled {
		return h.start2FA(&Login2FAConfig{
			User:    user,
			Command: cmd,
		}, ses)
	}
	return h.login(user, cmd, ses)
}

type Login2FAConfig struct {
	User    *user.User   `json:"user"`
	Command LoginCommand `json:"command"`
}

func (h loginHandler) start2FA(config *Login2FAConfig, ses *session.SessionUser) (*LoginResult, *i18np.Error) {
	h.publisher.CheckSubAndPublish(h.verifyTopics.Start2FA, helper.Verify.Start2FA(helper.Start2FAConfig{
		UserUUID:   config.User.UUID,
		DeviceUUID: config.Command.DeviceUUID,
		Redirect: helper.Start2FARedirect{
			WebURL:     "https://api.turistikrota.com/auth/check-2fa",
			Stream:     h.authTopics.LoginVerified,
			BaseStream: h.authTopics.Base,
			StreamData: config,
		},
	}))
	token, err := h.sessionSrv.New2FA(session.NewCommand{
		UserUUID:   config.User.UUID,
		DeviceUUID: config.Command.DeviceUUID,
		Device:     config.Command.Device,
		User:       ses,
	})
	if err != nil {
		return nil, h.errors.Failed("token")
	}
	return &LoginResult{
		Verify:    true,
		TempToken: token,
	}, nil
}

func (h loginHandler) login(user *user.User, cmd LoginCommand, ses *session.SessionUser) (*LoginResult, *i18np.Error) {
	tokens, error := h.sessionSrv.New(session.NewCommand{
		UserUUID:   user.UUID,
		DeviceUUID: cmd.DeviceUUID,
		Device:     cmd.Device,
		User:       ses,
	})
	if error != nil {
		return nil, h.errors.Failed("token")
	}
	_ = h.publisher.Publish(h.authTopics.LoggedIn, user)
	return &LoginResult{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		Verify:       false,
	}, nil
}

func (h loginHandler) getUserRelations(ctx context.Context, userUUID string) ([]jwt.UserClaimAccount, []jwt.UserClaimOwner, *i18np.Error) {
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
