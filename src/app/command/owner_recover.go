package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.auth/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnerRecoverCommand struct {
	NickName string
}

type OwnerRecoverResult struct{}

type OwnerRecoverHandler decorator.CommandHandler[OwnerRecoverCommand, *OwnerRecoverResult]

type ownerRecoverHandler struct {
	repo owner.Repository
}

type OwnerRecoverHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerRecoverHandler(config OwnerRecoverHandlerConfig) OwnerRecoverHandler {
	return decorator.ApplyCommandDecorators[OwnerRecoverCommand, *OwnerRecoverResult](
		ownerRecoverHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerRecoverHandler) Handle(ctx context.Context, command OwnerRecoverCommand) (*OwnerRecoverResult, *i18np.Error) {
	err := h.repo.Recover(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &OwnerRecoverResult{}, nil
}
