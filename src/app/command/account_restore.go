package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/account"
)

type AccountRestoreCommand struct {
	UserUUID string
	Name     string
}
type AccountRestoreResult struct{}

type AccountRestoreHandler decorator.CommandHandler[AccountRestoreCommand, *AccountRestoreResult]

type accountRestoreHandler struct {
	repo account.Repository
}

type AccountRestoreHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountRestoreHandler(config AccountRestoreHandlerConfig) AccountRestoreHandler {
	return decorator.ApplyCommandDecorators[AccountRestoreCommand, *AccountRestoreResult](
		&accountRestoreHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountRestoreHandler) Handle(ctx context.Context, cmd AccountRestoreCommand) (*AccountRestoreResult, *i18np.Error) {
	_ = h.repo.Restore(ctx, account.UserUnique{
		UserUUID: cmd.UserUUID,
		Name:     cmd.Name,
	})
	return &AccountRestoreResult{}, nil
}
