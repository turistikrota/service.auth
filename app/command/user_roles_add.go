package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/domains/user"
)

type UserRolesAddCmd struct {
	UserUUID    string   `json:"uuid"`
	Permissions []string `json:"permissions"`
}

type UserRolesAddRes struct{}

type UserRolesAddHandler cqrs.HandlerFunc[UserRolesAddCmd, *UserRolesAddRes]

func NewUserRolesAddHandler(repo user.Repo) UserRolesAddHandler {
	return func(ctx context.Context, cmd UserRolesAddCmd) (*UserRolesAddRes, *i18np.Error) {
		err := repo.AddRoles(ctx, cmd.UserUUID, cmd.Permissions)
		if err != nil {
			return nil, err
		}
		return &UserRolesAddRes{}, nil
	}
}
