package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessRemoveUserCommand struct {
	NickName string
	User     business.UserDetail
}

type BusinessRemoveUserResult struct{}

type BusinessRemoveUserHandler decorator.CommandHandler[BusinessRemoveUserCommand, *BusinessRemoveUserResult]

type businessRemoveUserHandler struct {
	repo business.Repository
}

type BusinessRemoveUserHandlerConfig struct {
	Repo     business.Repository
	CqrsBase decorator.Base
}

func NewBusinessRemoveUserHandler(config BusinessRemoveUserHandlerConfig) BusinessRemoveUserHandler {
	return decorator.ApplyCommandDecorators[BusinessRemoveUserCommand, *BusinessRemoveUserResult](
		businessRemoveUserHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h businessRemoveUserHandler) Handle(ctx context.Context, command BusinessRemoveUserCommand) (*BusinessRemoveUserResult, *i18np.Error) {
	err := h.repo.RemoveUser(ctx, command.NickName, command.User)
	if err != nil {
		return nil, err
	}
	return &BusinessRemoveUserResult{}, nil
}
