package command

import (
	"context"

	"api.turistikrota.com/auth/src/domain/owner"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnerAddUserPermissionCommand struct {
	NickName   string
	User       owner.UserDetail
	Permission string
}

type OwnerAddUserPermissionResult struct{}

type OwnerAddUserPermissionHandler decorator.CommandHandler[OwnerAddUserPermissionCommand, *OwnerAddUserPermissionResult]

type ownerAddUserPermissionHandler struct {
	repo owner.Repository
}

type OwnerAddUserPermissionHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerAddUserPermissionHandler(config OwnerAddUserPermissionHandlerConfig) OwnerAddUserPermissionHandler {
	return decorator.ApplyCommandDecorators[OwnerAddUserPermissionCommand, *OwnerAddUserPermissionResult](
		ownerAddUserPermissionHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerAddUserPermissionHandler) Handle(ctx context.Context, command OwnerAddUserPermissionCommand) (*OwnerAddUserPermissionResult, *i18np.Error) {
	err := h.repo.AddUserPermission(ctx, command.NickName, command.User, command.Permission)
	if err != nil {
		return nil, err
	}
	return &OwnerAddUserPermissionResult{}, nil
}
