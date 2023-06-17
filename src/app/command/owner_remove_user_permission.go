package command

import (
	"context"

	"api.turistikrota.com/auth/src/domain/owner"
	"api.turistikrota.com/shared/decorator"
	"github.com/mixarchitecture/i18np"
)

type OwnerRemoveUserPermissionCommand struct {
	NickName   string
	User       owner.UserDetail
	Permission string
}

type OwnerRemoveUserPermissionResult struct{}

type OwnerRemoveUserPermissionHandler decorator.CommandHandler[OwnerRemoveUserPermissionCommand, *OwnerRemoveUserPermissionResult]

type ownerRemoveUserPermissionHandler struct {
	repo owner.Repository
}

type OwnerRemoveUserPermissionHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerRemoveUserPermissionHandler(config OwnerRemoveUserPermissionHandlerConfig) OwnerRemoveUserPermissionHandler {
	return decorator.ApplyCommandDecorators[OwnerRemoveUserPermissionCommand, *OwnerRemoveUserPermissionResult](
		ownerRemoveUserPermissionHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerRemoveUserPermissionHandler) Handle(ctx context.Context, command OwnerRemoveUserPermissionCommand) (*OwnerRemoveUserPermissionResult, *i18np.Error) {
	err := h.repo.RemoveUserPermission(ctx, command.NickName, command.User, command.Permission)
	if err != nil {
		return nil, err
	}
	return &OwnerRemoveUserPermissionResult{}, nil
}
