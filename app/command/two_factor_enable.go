package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/domains/user"
)

type TwoFactorEnableCmd struct {
	UserUUID string `json:"-"`
}

type TwoFactorEnableRes struct{}

type TwoFactorEnableHandler cqrs.HandlerFunc[TwoFactorEnableCmd, *TwoFactorEnableRes]

func NewTwoFactorEnableHandler(repo user.Repo) TwoFactorEnableHandler {
	return func(ctx context.Context, cmd TwoFactorEnableCmd) (*TwoFactorEnableRes, *i18np.Error) {
		err := repo.EnableTwoFactor(ctx, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		return &TwoFactorEnableRes{}, nil
	}
}
