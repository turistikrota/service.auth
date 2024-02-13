package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/domains/user"
	"github.com/turistikrota/service.shared/auth/session"
)

type UserDeleteCmd struct {
	UserUUID   string `json:"-"`
	DeviceUUID string `json:"-"`
}

type UserDeleteRes struct{}

type UserDeleteHandler cqrs.HandlerFunc[UserDeleteCmd, *UserDeleteRes]

func NewUserDeleteHandler(sessionSrv session.Service, repo user.Repo) UserDeleteHandler {
	return func(ctx context.Context, cmd UserDeleteCmd) (*UserDeleteRes, *i18np.Error) {
		err := repo.Delete(ctx, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		err = sessionSrv.Destroy(session.DestroyCommand{
			UserUUID:   cmd.UserUUID,
			DeviceUUID: cmd.DeviceUUID,
		})
		if err != nil {
			return nil, err
		}
		return &UserDeleteRes{}, nil
	}
}
