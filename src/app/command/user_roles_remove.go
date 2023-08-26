package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/user"
)

type UserRolesRemoveCommand struct {
	UserUUID    string
	Permissions []string
}

type UserRolesRemoveResult struct{}

type UserRolesRemoveHandler decorator.CommandHandler[UserRolesRemoveCommand, *UserRolesRemoveResult]

type userRolesRemoveHandler struct {
	repo user.Repository
}

type UserRolesRemoveHandlerConfig struct {
	Repo     user.Repository
	CqrsBase decorator.Base
}

func NewUserRolesRemoveHandler(config UserRolesRemoveHandlerConfig) UserRolesRemoveHandler {
	return decorator.ApplyCommandDecorators[UserRolesRemoveCommand, *UserRolesRemoveResult](
		userRolesRemoveHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h userRolesRemoveHandler) Handle(ctx context.Context, command UserRolesRemoveCommand) (*UserRolesRemoveResult, *i18np.Error) {
	err := h.repo.RemoveRoles(ctx, command.UserUUID, command.Permissions)
	if err != nil {
		return nil, err
	}
	return &UserRolesRemoveResult{}, nil
}
