package command

import (
	"context"
	"time"

	"api.turistikrota.com/auth/src/domain/account"
	"github.com/turistikrota/service.shared/decorator"

	"github.com/mixarchitecture/i18np"
)

type AccountUpdateCommand struct {
	UserUUID    string
	CurrentName string
	NewName     string
	BirthDate   *time.Time
}
type AccountUpdateResult struct{}

type AccountUpdateHandler decorator.CommandHandler[AccountUpdateCommand, *AccountUpdateResult]

type accountUpdateHandler struct {
	repo account.Repository
}

type AccountUpdateHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountUpdateHandler(config AccountUpdateHandlerConfig) AccountUpdateHandler {
	return decorator.ApplyCommandDecorators[AccountUpdateCommand, *AccountUpdateResult](
		&accountUpdateHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountUpdateHandler) Handle(ctx context.Context, cmd AccountUpdateCommand) (*AccountUpdateResult, *i18np.Error) {
	_ = h.repo.Update(ctx, account.UserUnique{
		UserUUID: cmd.UserUUID,
		Name:     cmd.CurrentName,
	}, &account.Entity{
		UserName:  cmd.NewName,
		BirthDate: cmd.BirthDate,
	})
	return &AccountUpdateResult{}, nil
}
