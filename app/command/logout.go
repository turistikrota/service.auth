package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.shared/auth/session"
)

type LogoutCmd struct {
	UserUUID   string `json:"-"`
	DeviceUUID string `json:"-"`
}

type LogoutRes struct{}

type LogoutHandler cqrs.HandlerFunc[LogoutCmd, *LogoutRes]

func NewLogoutHandler(sessionSrv session.Service) LogoutHandler {
	return func(ctx context.Context, cmd LogoutCmd) (*LogoutRes, *i18np.Error) {
		err := sessionSrv.Destroy(ctx, session.DestroyCommand{
			UserUUID:   cmd.UserUUID,
			DeviceUUID: cmd.DeviceUUID,
		})
		if err != nil {
			return nil, err
		}
		return &LogoutRes{}, nil
	}
}
