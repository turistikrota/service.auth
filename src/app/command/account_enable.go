package command

import (
	"context"

	"api.turistikrota.com/auth/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountEnableCommand struct {
	UserUUID string
	Name     string
}
type AccountEnableResult struct{}

type AccountEnableHandler decorator.CommandHandler[AccountEnableCommand, *AccountEnableResult]

type accountEnableHandler struct {
	repo account.Repository
}

type AccountEnableHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountEnableHandler(config AccountEnableHandlerConfig) AccountEnableHandler {
	return decorator.ApplyCommandDecorators[AccountEnableCommand, *AccountEnableResult](
		&accountEnableHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountEnableHandler) Handle(ctx context.Context, cmd AccountEnableCommand) (*AccountEnableResult, *i18np.Error) {
	_ = h.repo.Enable(ctx, account.UserUnique{
		UserUUID: cmd.UserUUID,
		Name:     cmd.Name,
	})
	return &AccountEnableResult{}, nil
}
