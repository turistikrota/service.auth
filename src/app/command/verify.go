package command

import (
	"context"
	"time"

	"api.turistikrota.com/auth/src/domain/user"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type VerifyCommand struct {
	Token string
}

type VerifyResult struct{}

type ExpiredError struct {
	Email  string `json:"email"`
	ReSend bool   `json:"reSendable"`
}

type VerifyHandler decorator.CommandHandler[VerifyCommand, *VerifyResult]

type verifyHandler struct {
	repo    user.Repository
	factory user.Factory
	events  user.Events
}

type VerifyHandlerConfig struct {
	Repo     user.Repository
	Factory  user.Factory
	Events   user.Events
	CqrsBase decorator.Base
}

func NewVerifyHandler(config VerifyHandlerConfig) VerifyHandler {
	return decorator.ApplyCommandDecorators[VerifyCommand, *VerifyResult](
		verifyHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h verifyHandler) Handle(ctx context.Context, command VerifyCommand) (*VerifyResult, *i18np.Error) {
	u, err := h.repo.GetByToken(ctx, command.Token)
	if err != nil {
		return nil, err
	}
	if time.Now().Sub(u.UpdatedAt) > 24*time.Hour {
		return nil, h.factory.Errors.TokenExpired(ExpiredError{
			Email:  u.Email,
			ReSend: true,
		})
	}
	if u.IsVerified {
		return nil, h.factory.Errors.AlreadyVerified()
	}
	err = h.repo.Verify(ctx, command.Token)
	if err != nil {
		return nil, err
	}
	h.events.UserVerified(user.UserVerifiedEvent{
		UserUUID: u.UUID,
		User:     *u,
	})
	return &VerifyResult{}, nil
}
