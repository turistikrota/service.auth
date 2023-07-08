package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/owner"
)

type OwnerAddUserCommand struct {
	NickName string
	User     *owner.User
}

type OwnerAddUserResult struct{}

type OwnerAddUserHandler decorator.CommandHandler[OwnerAddUserCommand, *OwnerAddUserResult]

type ownerAddUserHandler struct {
	repo owner.Repository
}

type OwnerAddUserHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerAddUserHandler(config OwnerAddUserHandlerConfig) OwnerAddUserHandler {
	return decorator.ApplyCommandDecorators[OwnerAddUserCommand, *OwnerAddUserResult](
		ownerAddUserHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerAddUserHandler) Handle(ctx context.Context, command OwnerAddUserCommand) (*OwnerAddUserResult, *i18np.Error) {
	err := h.repo.AddUser(ctx, command.NickName, command.User)
	if err != nil {
		return nil, err
	}
	return &OwnerAddUserResult{}, nil
}
