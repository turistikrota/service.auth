package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
	"github.com/turistikrota/service.auth/domains/user"
	"github.com/turistikrota/service.shared/cipher"
)

type RegisterCmd struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Privacy  *bool  `json:"privacy" validate:"required,eq=true"`
	Lang     string `json:"-"`
}

type RegisterRes struct{}

type RegisterHandler cqrs.HandlerFunc[RegisterCmd, *RegisterRes]

func NewRegisterHandler(repo user.Repo, factory user.Factory, events user.Events) RegisterHandler {
	return func(ctx context.Context, cmd RegisterCmd) (*RegisterRes, *i18np.Error) {
		pw, error := cipher.Hash(cmd.Password)
		if error != nil {
			return nil, factory.Errors.Failed("hash")
		}
		token := uuid.New().String()
		u, err := repo.Create(ctx, factory.New(cmd.Email, pw, token))
		if err != nil {
			return nil, err
		}
		events.SendVerification(user.SendVerificationEvent{
			Email: u.Email,
			Token: token,
			Lang:  cmd.Lang,
		})
		return &RegisterRes{}, nil
	}
}
