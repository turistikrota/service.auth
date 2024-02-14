package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionDestroyOthersCmd struct {
	UserUUID   string `json:"-"`
	DeviceUUID string `json:"-"`
}

type SessionDestroyOthersRes struct {
}

type SessionDestroyOthersHandler cqrs.HandlerFunc[SessionDestroyOthersCmd, *SessionDestroyOthersRes]

func NewSessionDestroyOthersHandler(sessionSrv session.Service) SessionDestroyOthersHandler {
	return func(ctx context.Context, cmd SessionDestroyOthersCmd) (*SessionDestroyOthersRes, *i18np.Error) {
		err := sessionSrv.DestroyOthers(cmd.UserUUID, cmd.DeviceUUID)
		if err != nil {
			return nil, err
		}
		return &SessionDestroyOthersRes{}, nil
	}
}
