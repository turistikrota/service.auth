package command

import (
	"context"

	"api.turistikrota.com/auth/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountCreateCommand struct {
	UserUUID    string
	AccountName string
	AccountCode string
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
		UserCode:   cmd.AccountCode,
		IsActive:   false,
		IsDeleted:  false,
		IsVerified: false,
		BirthDate:  nil,
	})
	return &AccountCreateResult{}, nil
}
