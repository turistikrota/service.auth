package command

import (
	"context"

	"api.turistikrota.com/auth/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountDisableCommand struct {
	UserUUID string
	Name     string
	Code     string
}
type AccountDisableResult struct{}

type AccountDisableHandler decorator.CommandHandler[AccountDisableCommand, *AccountDisableResult]

type accountDisableHandler struct {
	repo account.Repository
}

type AccountDisableHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountDisableHandler(config AccountDisableHandlerConfig) AccountDisableHandler {
	return decorator.ApplyCommandDecorators[AccountDisableCommand, *AccountDisableResult](
		&accountDisableHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountDisableHandler) Handle(ctx context.Context, cmd AccountDisableCommand) (*AccountDisableResult, *i18np.Error) {
	_ = h.repo.Disable(ctx, account.UserUnique{
		UserUUID: cmd.UserUUID,
		Name:     cmd.Name,
		Code:     cmd.Code,
	})
	return &AccountDisableResult{}, nil
}
