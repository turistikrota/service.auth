package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessCreateCommand struct {
	Entity *business.Entity
}

type BusinessCreateResult struct{}

type BusinessCreateHandler decorator.CommandHandler[BusinessCreateCommand, *BusinessCreateResult]

type businessCreateHandler struct {
	repo    business.Repository
	factory business.Factory
}

type BusinessCreateHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	CqrsBase decorator.Base
}

func NewBusinessCreateHandler(config BusinessCreateHandlerConfig) BusinessCreateHandler {
	return decorator.ApplyCommandDecorators[BusinessCreateCommand, *BusinessCreateResult](
		businessCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h businessCreateHandler) Handle(ctx context.Context, command BusinessCreateCommand) (*BusinessCreateResult, *i18np.Error) {
	err := h.repo.Create(ctx, command.Entity)
	if err != nil {
		return nil, err
	}
	return &BusinessCreateResult{}, nil
}
