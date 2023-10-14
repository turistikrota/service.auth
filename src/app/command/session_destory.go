package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionDestroyCommand struct {
	UserUUID   string
	DeviceUUID string
}

type SessionDestroyResult struct {
	Sessions []session.Entity
}

type SessionDestroyHandler decorator.CommandHandler[SessionDestroyCommand, *SessionDestroyResult]

type sessionDestroyHandler struct {
	sessionSrv session.Service
}

type SessionDestroyHandlerConfig struct {
	SessionSrv session.Service
	CqrsBase   decorator.Base
}

func NewSessionDestroyHandler(config SessionDestroyHandlerConfig) SessionDestroyHandler {
	return decorator.ApplyQueryDecorators[SessionDestroyCommand, *SessionDestroyResult](
		sessionDestroyHandler{
			sessionSrv: config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h sessionDestroyHandler) Handle(ctx context.Context, cmd SessionDestroyCommand) (*SessionDestroyResult, *i18np.Error) {
	err := h.sessionSrv.Destroy(session.DestroyCommand{
		UserUUID:   cmd.UserUUID,
		DeviceUUID: cmd.DeviceUUID,
	})
	if err != nil {
		return nil, err
	}
	return &SessionDestroyResult{}, nil
}
