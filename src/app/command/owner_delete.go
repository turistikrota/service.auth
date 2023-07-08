package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.auth/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnerDeleteCommand struct {
	NickName string
}

type OwnerDeleteResult struct{}

type OwnerDeleteHandler decorator.CommandHandler[OwnerDeleteCommand, *OwnerDeleteResult]

type ownerDeleteHandler struct {
	repo owner.Repository
}

type OwnerDeleteHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerDeleteHandler(config OwnerDeleteHandlerConfig) OwnerDeleteHandler {
	return decorator.ApplyCommandDecorators[OwnerDeleteCommand, *OwnerDeleteResult](
		ownerDeleteHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerDeleteHandler) Handle(ctx context.Context, command OwnerDeleteCommand) (*OwnerDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &OwnerDeleteResult{}, nil
}
