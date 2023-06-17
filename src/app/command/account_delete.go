package command

import (
	"context"

	"api.turistikrota.com/auth/src/domain/account"
	"api.turistikrota.com/shared/decorator"
	"github.com/mixarchitecture/i18np"
)

type AccountDeleteCommand struct {
	UserUUID string
	Name     string
	Code     string
}
type AccountDeleteResult struct{}

type AccountDeleteHandler decorator.CommandHandler[AccountDeleteCommand, *AccountDeleteResult]

type accountDeleteHandler struct {
	repo account.Repository
}

type AccountDeleteHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountDeleteHandler(config AccountDeleteHandlerConfig) AccountDeleteHandler {
	return decorator.ApplyCommandDecorators[AccountDeleteCommand, *AccountDeleteResult](
		&accountDeleteHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountDeleteHandler) Handle(ctx context.Context, cmd AccountDeleteCommand) (*AccountDeleteResult, *i18np.Error) {
	_ = h.repo.Delete(ctx, account.UserUnique{
		UserUUID: cmd.UserUUID,
		Name:     cmd.Name,
		Code:     cmd.Code,
	})
	return &AccountDeleteResult{}, nil
}
