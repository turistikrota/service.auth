package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessRemoveUserPermissionCommand struct {
	NickName   string
	User       business.UserDetail
	Permission string
}

type BusinessRemoveUserPermissionResult struct{}

type BusinessRemoveUserPermissionHandler decorator.CommandHandler[BusinessRemoveUserPermissionCommand, *BusinessRemoveUserPermissionResult]

type businessRemoveUserPermissionHandler struct {
	repo business.Repository
}

type BusinessRemoveUserPermissionHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessRemoveUserPermissionHandler(config BusinessRemoveUserPermissionHandlerConfig) BusinessRemoveUserPermissionHandler {
	return decorator.ApplyCommandDecorators[BusinessRemoveUserPermissionCommand, *BusinessRemoveUserPermissionResult](
		businessRemoveUserPermissionHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessRemoveUserPermissionHandler) Handle(ctx context.Context, command BusinessRemoveUserPermissionCommand) (*BusinessRemoveUserPermissionResult, *i18np.Error) {
	err := h.repo.RemoveUserPermission(ctx, command.NickName, command.User, command.Permission)
	if err != nil {
		return nil, err
	}
	return &BusinessRemoveUserPermissionResult{}, nil
}
