package command

import (
	"context"

	"api.turistikrota.com/auth/src/domain/user"
	"api.turistikrota.com/shared/decorator"
	"github.com/mixarchitecture/i18np"
)

type UserUpdatedCommand struct {
	UserUUID         string
	TwoFactorEnabled bool
}

type UserUpdatedResult struct{}

type UserUpdatedHandler decorator.CommandHandler[UserUpdatedCommand, *UserUpdatedResult]

type userUpdatedHandler struct {
	repo user.Repository
}

type UserUpdatedHandlerConfig struct {
	Repo     user.Repository
	CqrsBase decorator.Base
}

func NewUserUpdatedHandler(config UserUpdatedHandlerConfig) UserUpdatedHandler {
	return decorator.ApplyCommandDecorators[UserUpdatedCommand, *UserUpdatedResult](
		userUpdatedHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h userUpdatedHandler) Handle(ctx context.Context, cmd UserUpdatedCommand) (*UserUpdatedResult, *i18np.Error) {
	u, err := h.repo.GetByUUID(ctx, cmd.UserUUID)
	if err != nil {
		return nil, err
	}
	u.TwoFactorEnabled = cmd.TwoFactorEnabled
	_, err = h.repo.UpdateByUUID(ctx, u)
	return nil, err
}
