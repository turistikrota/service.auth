package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessDeleteCommand struct {
	NickName string
}

type BusinessDeleteResult struct{}

type BusinessDeleteHandler decorator.CommandHandler[BusinessDeleteCommand, *BusinessDeleteResult]

type businessDeleteHandler struct {
	repo business.Repository
}

type BusinessDeleteHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessDeleteHandler(config BusinessDeleteHandlerConfig) BusinessDeleteHandler {
	return decorator.ApplyCommandDecorators[BusinessDeleteCommand, *BusinessDeleteResult](
		businessDeleteHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessDeleteHandler) Handle(ctx context.Context, command BusinessDeleteCommand) (*BusinessDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &BusinessDeleteResult{}, nil
}
