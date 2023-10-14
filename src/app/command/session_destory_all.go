package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionDestroyAllCommand struct {
	UserUUID string
}

type SessionDestroyAllResult struct {
	Sessions []session.Entity
}

type SessionDestroyAllHandler decorator.CommandHandler[SessionDestroyAllCommand, *SessionDestroyAllResult]

type sessionDestroyAllHandler struct {
	sessionSrv session.Service
}

type SessionDestroyAllHandlerConfig struct {
	SessionSrv session.Service
	CqrsBase   decorator.Base
}

func NewSessionDestroyAllHandler(config SessionDestroyAllHandlerConfig) SessionDestroyAllHandler {
	return decorator.ApplyQueryDecorators[SessionDestroyAllCommand, *SessionDestroyAllResult](
		sessionDestroyAllHandler{
			sessionSrv: config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h sessionDestroyAllHandler) Handle(ctx context.Context, cmd SessionDestroyAllCommand) (*SessionDestroyAllResult, *i18np.Error) {
	err := h.sessionSrv.DestroyAll(cmd.UserUUID)
	if err != nil {
		return nil, err
	}
	return &SessionDestroyAllResult{}, nil
}
