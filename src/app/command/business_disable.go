package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessDisableCommand struct {
	NickName string
}

type BusinessDisableResult struct{}

type BusinessDisableHandler decorator.CommandHandler[BusinessDisableCommand, *BusinessDisableResult]

type businessDisableHandler struct {
	repo business.Repository
}

type BusinessDisableHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessDisableHandler(config BusinessDisableHandlerConfig) BusinessDisableHandler {
	return decorator.ApplyCommandDecorators[BusinessDisableCommand, *BusinessDisableResult](
		businessDisableHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessDisableHandler) Handle(ctx context.Context, command BusinessDisableCommand) (*BusinessDisableResult, *i18np.Error) {
	err := h.repo.Disable(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &BusinessDisableResult{}, nil
}
