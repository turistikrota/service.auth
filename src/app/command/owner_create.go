package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/owner"
)

type OwnerCreateCommand struct {
	Entity *owner.Entity
}

type OwnerCreateResult struct{}

type OwnerCreateHandler decorator.CommandHandler[OwnerCreateCommand, *OwnerCreateResult]

type ownerCreateHandler struct {
	repo    owner.Repository
	factory owner.Factory
}

type OwnerCreateHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	CqrsBase decorator.Base
}

func NewOwnerCreateHandler(config OwnerCreateHandlerConfig) OwnerCreateHandler {
	return decorator.ApplyCommandDecorators[OwnerCreateCommand, *OwnerCreateResult](
		ownerCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h ownerCreateHandler) Handle(ctx context.Context, command OwnerCreateCommand) (*OwnerCreateResult, *i18np.Error) {
	err := h.repo.Create(ctx, command.Entity)
	if err != nil {
		return nil, err
	}
	return &OwnerCreateResult{}, nil
}
