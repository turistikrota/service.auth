package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/user"
)

type UserRolesAddCommand struct {
	UserUUID    string
	Permissions []string
}

type UserRolesAddResult struct{}

type UserRolesAddHandler decorator.CommandHandler[UserRolesAddCommand, *UserRolesAddResult]

type userRolesAddHandler struct {
	repo user.Repository
}

type UserRolesAddHandlerConfig struct {
	Repo     user.Repository
	CqrsBase decorator.Base
}

func NewUserRolesAddHandler(config UserRolesAddHandlerConfig) UserRolesAddHandler {
	return decorator.ApplyCommandDecorators[UserRolesAddCommand, *UserRolesAddResult](
		userRolesAddHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h userRolesAddHandler) Handle(ctx context.Context, command UserRolesAddCommand) (*UserRolesAddResult, *i18np.Error) {
	err := h.repo.AddRoles(ctx, command.UserUUID, command.Permissions)
	if err != nil {
		return nil, err
	}
	return &UserRolesAddResult{}, nil
}
