package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
	"github.com/turistikrota/service.auth/domains/user"
)

type ReSendVerificationCodeCmd struct {
	Email string `json:"email" validate:"required,email"`
	Lang  string `json:"-"`
}

type ReSendVerificationCodeRes struct{}

type ReSendVerificationCodeHandler cqrs.HandlerFunc[ReSendVerificationCodeCmd, *ReSendVerificationCodeRes]

func NewReSendVerificationCodeHandler(repo user.Repo, factory user.Factory, events user.Events) ReSendVerificationCodeHandler {
	return func(ctx context.Context, cmd ReSendVerificationCodeCmd) (*ReSendVerificationCodeRes, *i18np.Error) {
		u, err := repo.GetByEmail(ctx, cmd.Email)
		if err != nil {
			return nil, err
		}
		if u.IsVerified {
			return nil, factory.Errors.AlreadyVerified()
		}
		if time.Since(u.UpdatedAt) <= 24*time.Hour {
			return nil, factory.Errors.TokenNotExpired()
		}
		token := uuid.New().String()
		err = repo.SetToken(ctx, cmd.Email, token)
		if err != nil {
			return nil, err
		}
		events.SendVerification(user.SendVerificationEvent{
			Email: u.Email,
			Token: token,
			Lang:  cmd.Lang,
		})
		return &ReSendVerificationCodeRes{}, nil
	}
}
