package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessAddUserPermissionCommand struct {
	NickName   string
	User       business.UserDetail
	Permission string
}

type BusinessAddUserPermissionResult struct{}

type BusinessAddUserPermissionHandler decorator.CommandHandler[BusinessAddUserPermissionCommand, *BusinessAddUserPermissionResult]

type businessAddUserPermissionHandler struct {
	repo business.Repository
}

type BusinessAddUserPermissionHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessAddUserPermissionHandler(config BusinessAddUserPermissionHandlerConfig) BusinessAddUserPermissionHandler {
	return decorator.ApplyCommandDecorators[BusinessAddUserPermissionCommand, *BusinessAddUserPermissionResult](
		businessAddUserPermissionHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessAddUserPermissionHandler) Handle(ctx context.Context, command BusinessAddUserPermissionCommand) (*BusinessAddUserPermissionResult, *i18np.Error) {
	err := h.repo.AddUserPermission(ctx, command.NickName, command.User, command.Permission)
	if err != nil {
		return nil, err
	}
	return &BusinessAddUserPermissionResult{}, nil
}
