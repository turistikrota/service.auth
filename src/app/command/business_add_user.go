package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessAddUserCommand struct {
	NickName string
	User     *business.User
}

type BusinessAddUserResult struct{}

type BusinessAddUserHandler decorator.CommandHandler[BusinessAddUserCommand, *BusinessAddUserResult]

type businessAddUserHandler struct {
	repo business.Repository
}

type BusinessAddUserHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessAddUserHandler(config BusinessAddUserHandlerConfig) BusinessAddUserHandler {
	return decorator.ApplyCommandDecorators[BusinessAddUserCommand, *BusinessAddUserResult](
		businessAddUserHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessAddUserHandler) Handle(ctx context.Context, command BusinessAddUserCommand) (*BusinessAddUserResult, *i18np.Error) {
	err := h.repo.AddUser(ctx, command.NickName, command.User)
	if err != nil {
		return nil, err
	}
	return &BusinessAddUserResult{}, nil
}
