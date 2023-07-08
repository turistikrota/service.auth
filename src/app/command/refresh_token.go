package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.auth/src/config"
	"github.com/turistikrota/service.auth/src/domain/user"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
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
	userRepo   user.Repository
	authTopics config.AuthTopics
	publisher  events.Publisher
	tokenSrv   token.Service
	sessionSrv session.Service
	errors     user.Errors
}

type RefreshTokenHandlerConfig struct {
	UserRepo   user.Repository
	AuthTopics config.AuthTopics
	Publisher  events.Publisher
	TokenSrv   token.Service
	SessionSrv session.Service
	Errors     user.Errors
	CqrsBase   decorator.Base
}

func NewRefreshTokenHandler(config RefreshTokenHandlerConfig) RefreshTokenHandler {
	return decorator.ApplyCommandDecorators[RefreshTokenCommand, *RefreshTokenResult](
		refreshTokenHandler{
			userRepo:   config.UserRepo,
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			tokenSrv:   config.TokenSrv,
			sessionSrv: config.SessionSrv,
			errors:     config.Errors,
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
	tokens, err := h.sessionSrv.Refresh(session.RefreshCommand{
		UserUUID:     cmd.UserUUID,
		DeviceUUID:   cmd.DeviceUUID,
		RefreshToken: cmd.RefreshToken,
		AccessToken:  cmd.AccessToken,
		User:         user.ToSession(),
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
