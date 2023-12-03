package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessVerifyCommand struct {
	NickName string
}

type BusinessVerifyResult struct{}

type BusinessVerifyHandler decorator.CommandHandler[BusinessVerifyCommand, *BusinessVerifyResult]

type businessVerifyHandler struct {
	repo business.Repository
}

type BusinessVerifyHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessVerifyHandler(config BusinessVerifyHandlerConfig) BusinessVerifyHandler {
	return decorator.ApplyCommandDecorators[BusinessVerifyCommand, *BusinessVerifyResult](
		businessVerifyHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessVerifyHandler) Handle(ctx context.Context, command BusinessVerifyCommand) (*BusinessVerifyResult, *i18np.Error) {
	err := h.repo.Verify(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &BusinessVerifyResult{}, nil
}
