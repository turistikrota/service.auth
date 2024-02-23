package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/domains/user"
)

type TwoFactorDisableCmd struct {
	UserUUID string `json:"-"`
}

type TwoFactorDisableRes struct{}

type TwoFactorDisableHandler cqrs.HandlerFunc[TwoFactorDisableCmd, *TwoFactorDisableRes]

func NewTwoFactorDisableHandler(repo user.Repo) TwoFactorDisableHandler {
	return func(ctx context.Context, cmd TwoFactorDisableCmd) (*TwoFactorDisableRes, *i18np.Error) {
		err := repo.DisableTwoFactor(ctx, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		return &TwoFactorDisableRes{}, nil
	}
}
