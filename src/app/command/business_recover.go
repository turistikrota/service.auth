package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessRecoverCommand struct {
	NickName string
}

type BusinessRecoverResult struct{}

type BusinessRecoverHandler decorator.CommandHandler[BusinessRecoverCommand, *BusinessRecoverResult]

type businessRecoverHandler struct {
	repo business.Repository
}

type BusinessRecoverHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessRecoverHandler(config BusinessRecoverHandlerConfig) BusinessRecoverHandler {
	return decorator.ApplyCommandDecorators[BusinessRecoverCommand, *BusinessRecoverResult](
		businessRecoverHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessRecoverHandler) Handle(ctx context.Context, command BusinessRecoverCommand) (*BusinessRecoverResult, *i18np.Error) {
	err := h.repo.Recover(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &BusinessRecoverResult{}, nil
}
