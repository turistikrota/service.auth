package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/domains/user"
)

type VerifyCmd struct {
	Token string `params:"token" validate:"required,uuid"`
}

type VerifyRes struct{}

type ExpiredError struct {
	Email  string `json:"email"`
	ReSend bool   `json:"reSendable"`
}

type VerifyHandler cqrs.HandlerFunc[VerifyCmd, *VerifyRes]

func NewVerifyHandler(repo user.Repo, factory user.Factory) VerifyHandler {
	return func(ctx context.Context, cmd VerifyCmd) (*VerifyRes, *i18np.Error) {
		u, err := repo.GetByToken(ctx, cmd.Token)
		if err != nil {
			return nil, err
		}
		if time.Since(u.UpdatedAt) > 24*time.Hour {
			return nil, factory.Errors.TokenExpired(i18np.P{
				"email":      u.Email,
				"reSendable": true,
			})
		}
		if u.IsVerified {
			return nil, factory.Errors.AlreadyVerified()
		}
		err = repo.Verify(ctx, cmd.Token)
		if err != nil {
			return nil, err
		}
		return &VerifyRes{}, nil
	}
}
