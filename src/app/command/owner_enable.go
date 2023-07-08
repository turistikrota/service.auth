package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/owner"
)

type OwnerEnableCommand struct {
	NickName string
}

type OwnerEnableResult struct{}

type OwnerEnableHandler decorator.CommandHandler[OwnerEnableCommand, *OwnerEnableResult]

type ownerEnableHandler struct {
	repo owner.Repository
}

type OwnerEnableHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerEnableHandler(config OwnerEnableHandlerConfig) OwnerEnableHandler {
	return decorator.ApplyCommandDecorators[OwnerEnableCommand, *OwnerEnableResult](
		ownerEnableHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerEnableHandler) Handle(ctx context.Context, command OwnerEnableCommand) (*OwnerEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &OwnerEnableResult{}, nil
}
