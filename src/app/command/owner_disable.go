package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.auth/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnerDisableCommand struct {
	NickName string
}

type OwnerDisableResult struct{}

type OwnerDisableHandler decorator.CommandHandler[OwnerDisableCommand, *OwnerDisableResult]

type ownerDisableHandler struct {
	repo owner.Repository
}

type OwnerDisableHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerDisableHandler(config OwnerDisableHandlerConfig) OwnerDisableHandler {
	return decorator.ApplyCommandDecorators[OwnerDisableCommand, *OwnerDisableResult](
		ownerDisableHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerDisableHandler) Handle(ctx context.Context, command OwnerDisableCommand) (*OwnerDisableResult, *i18np.Error) {
	err := h.repo.Disable(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &OwnerDisableResult{}, nil
}
