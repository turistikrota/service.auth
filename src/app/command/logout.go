package command

import (
	"context"

	"api.turistikrota.com/auth/src/config"
	"api.turistikrota.com/auth/src/domain/user"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/decorator"
	"github.com/turistikrota/service.shared/events"
)

type LogoutCommand struct {
	UserUUID   string
	DeviceUUID string
}

type LogoutResult struct{}

type LogoutHandler decorator.CommandHandler[LogoutCommand, *LogoutResult]

type logoutHandler struct {
	authTopics config.AuthTopics
	publisher  events.Publisher
	tokenSrv   token.Service
	sessionSrv session.Service
	errors     user.Errors
}

type LogoutHandlerConfig struct {
	AuthTopics config.AuthTopics
	Publisher  events.Publisher
	TokenSrv   token.Service
	SessionSrv session.Service
	Errors     user.Errors
	CqrsBase   decorator.Base
}

func NewLogoutHandler(config LogoutHandlerConfig) LogoutHandler {
	return decorator.ApplyCommandDecorators[LogoutCommand, *LogoutResult](
		logoutHandler{
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			tokenSrv:   config.TokenSrv,
			sessionSrv: config.SessionSrv,
			errors:     config.Errors,
		},
		config.CqrsBase,
	)
}

func (h logoutHandler) Handle(ctx context.Context, command LogoutCommand) (*LogoutResult, *i18np.Error) {
	err := h.sessionSrv.Destroy(session.DestroyCommand{
		UserUUID:   command.UserUUID,
		DeviceUUID: command.DeviceUUID,
	})
	if err != nil {
		return nil, err
	}
	return &LogoutResult{}, nil
}
