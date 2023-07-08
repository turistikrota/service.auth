package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/owner"
)

type OwnerRemoveUserCommand struct {
	NickName string
	User     owner.UserDetail
}

type OwnerRemoveUserResult struct{}

type OwnerRemoveUserHandler decorator.CommandHandler[OwnerRemoveUserCommand, *OwnerRemoveUserResult]

type ownerRemoveUserHandler struct {
	repo owner.Repository
}

type OwnerRemoveUserHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerRemoveUserHandler(config OwnerRemoveUserHandlerConfig) OwnerRemoveUserHandler {
	return decorator.ApplyCommandDecorators[OwnerRemoveUserCommand, *OwnerRemoveUserResult](
		ownerRemoveUserHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerRemoveUserHandler) Handle(ctx context.Context, command OwnerRemoveUserCommand) (*OwnerRemoveUserResult, *i18np.Error) {
	err := h.repo.RemoveUser(ctx, command.NickName, command.User)
	if err != nil {
		return nil, err
	}
	return &OwnerRemoveUserResult{}, nil
}
