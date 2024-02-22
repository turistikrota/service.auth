package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionDestroyAllCmd struct {
	UserUUID string `json:"-"`
}

type SessionDestroyAllRes struct {
}

type SessionDestroyAllHandler cqrs.HandlerFunc[SessionDestroyAllCmd, *SessionDestroyAllRes]

func NewSessionDestroyAllHandler(sessionSrv session.Service) SessionDestroyAllHandler {
	return func(ctx context.Context, cmd SessionDestroyAllCmd) (*SessionDestroyAllRes, *i18np.Error) {
		err := sessionSrv.DestroyAll(ctx, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		return &SessionDestroyAllRes{}, nil
	}
}
