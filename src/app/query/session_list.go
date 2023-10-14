package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionListQuery struct {
	UserUUID string
}

type SessionListResult struct {
	Sessions []session.Entity
}

type SessionListHandler decorator.QueryHandler[SessionListQuery, *SessionListResult]

type sessionListHandler struct {
	sessionSrv session.Service
}

type SessionListHandlerConfig struct {
	SessionSrv session.Service
	CqrsBase   decorator.Base
}

func NewSessionListHandler(config SessionListHandlerConfig) SessionListHandler {
	return decorator.ApplyQueryDecorators[SessionListQuery, *SessionListResult](
		sessionListHandler{
			sessionSrv: config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h sessionListHandler) Handle(ctx context.Context, query SessionListQuery) (*SessionListResult, *i18np.Error) {
	sessions, err := h.sessionSrv.GetAll(query.UserUUID)
	if err != nil {
		return nil, err
	}
	return &SessionListResult{
		Sessions: sessions,
	}, nil
}
