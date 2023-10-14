package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionDestroyOthersCommand struct {
	UserUUID   string
	DeviceUUID string
}

type SessionDestroyOthersResult struct {
	Sessions []session.Entity
}

type SessionDestroyOthersHandler decorator.CommandHandler[SessionDestroyOthersCommand, *SessionDestroyOthersResult]

type sessionDestroyOthersHandler struct {
	sessionSrv session.Service
}

type SessionDestroyOthersHandlerConfig struct {
	SessionSrv session.Service
	CqrsBase   decorator.Base
}

func NewSessionDestroyOthersHandler(config SessionDestroyOthersHandlerConfig) SessionDestroyOthersHandler {
	return decorator.ApplyQueryDecorators[SessionDestroyOthersCommand, *SessionDestroyOthersResult](
		sessionDestroyOthersHandler{
			sessionSrv: config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h sessionDestroyOthersHandler) Handle(ctx context.Context, cmd SessionDestroyOthersCommand) (*SessionDestroyOthersResult, *i18np.Error) {
	err := h.sessionSrv.DestroyOthers(cmd.UserUUID, cmd.DeviceUUID)
	if err != nil {
		return nil, err
	}
	return &SessionDestroyOthersResult{}, nil
}
