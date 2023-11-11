package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/account"
)

type AccountCreateCommand struct {
	UserUUID    string
	AccountName string
}
type AccountCreateResult struct{}

type AccountCreateHandler decorator.CommandHandler[AccountCreateCommand, *AccountCreateResult]

type accountCreateHandler struct {
	repo account.Repository
}

type AccountCreateHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountCreateHandler(config AccountCreateHandlerConfig) AccountCreateHandler {
	return decorator.ApplyCommandDecorators[AccountCreateCommand, *AccountCreateResult](
		&accountCreateHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountCreateHandler) Handle(ctx context.Context, cmd AccountCreateCommand) (*AccountCreateResult, *i18np.Error) {
	_ = h.repo.Create(ctx, &account.Entity{
		UserUUID:   cmd.UserUUID,
		UserName:   cmd.AccountName,
		IsActive:   true,
		IsDeleted:  false,
		IsVerified: false,
		BirthDate:  nil,
	})
	return &AccountCreateResult{}, nil
}
