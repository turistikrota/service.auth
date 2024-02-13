package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/domains/user"
)

type UserRolesRemoveCmd struct {
	UserUUID    string   `json:"uuid"`
	Permissions []string `json:"permissions"`
}

type UserRolesRemoveRes struct{}

type UserRolesRemoveHandler cqrs.HandlerFunc[UserRolesRemoveCmd, *UserRolesRemoveRes]

func NewUserRolesRemoveHandler(repo user.Repo) UserRolesRemoveHandler {
	return func(ctx context.Context, cmd UserRolesRemoveCmd) (*UserRolesRemoveRes, *i18np.Error) {
		err := repo.RemoveRoles(ctx, cmd.UserUUID, cmd.Permissions)
		if err != nil {
			return nil, err
		}
		return &UserRolesRemoveRes{}, nil
	}
}
