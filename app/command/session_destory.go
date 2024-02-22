package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionDestroyCmd struct {
	UserUUID   string `json:"-"`
	DeviceUUID string `params:"device_uuid" validate:"required,uuid"`
}

type SessionDestroyRes struct {
}

type SessionDestroyHandler cqrs.HandlerFunc[SessionDestroyCmd, *SessionDestroyRes]

func NewSessionDestroyHandler(sessionSrv session.Service) SessionDestroyHandler {
	return func(ctx context.Context, cmd SessionDestroyCmd) (*SessionDestroyRes, *i18np.Error) {
		err := sessionSrv.Destroy(ctx, session.DestroyCommand{
			UserUUID:   cmd.UserUUID,
			DeviceUUID: cmd.DeviceUUID,
		})
		if err != nil {
			return nil, err
		}
		return &SessionDestroyRes{}, nil
	}
}
