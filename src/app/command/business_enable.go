package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessEnableCommand struct {
	NickName string
}

type BusinessEnableResult struct{}

type BusinessEnableHandler decorator.CommandHandler[BusinessEnableCommand, *BusinessEnableResult]

type businessEnableHandler struct {
	repo business.Repository
}

type BusinessEnableHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessEnableHandler(config BusinessEnableHandlerConfig) BusinessEnableHandler {
	return decorator.ApplyCommandDecorators[BusinessEnableCommand, *BusinessEnableResult](
		businessEnableHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessEnableHandler) Handle(ctx context.Context, command BusinessEnableCommand) (*BusinessEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &BusinessEnableResult{}, nil
}
