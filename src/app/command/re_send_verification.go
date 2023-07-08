package command

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/user"
)

type ReSendVerificationCommand struct {
	Email string
	Lang  string
}

type ReSendVerificationResult struct{}

type ReSendVerificationHandler decorator.CommandHandler[ReSendVerificationCommand, *ReSendVerificationResult]

type reSendVerificationHandler struct {
	repo    user.Repository
	factory user.Factory
	events  user.Events
}

type ReSendVerificationHandlerConfig struct {
	Repo     user.Repository
	Factory  user.Factory
	Events   user.Events
	CqrsBase decorator.Base
}

func NewReSendVerificationHandler(config ReSendVerificationHandlerConfig) ReSendVerificationHandler {
	return decorator.ApplyCommandDecorators[ReSendVerificationCommand, *ReSendVerificationResult](
		reSendVerificationHandler{
			repo:    config.Repo,
			events:  config.Events,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h reSendVerificationHandler) Handle(ctx context.Context, cmd ReSendVerificationCommand) (*ReSendVerificationResult, *i18np.Error) {
	u, err := h.repo.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}
	if u.IsVerified {
		return nil, h.factory.Errors.AlreadyVerified()
	}
	if time.Now().Sub(u.UpdatedAt) <= 24*time.Hour {
		return nil, h.factory.Errors.TokenNotExpired()
	}
	token := uuid.New().String()
	err = h.repo.SetToken(ctx, cmd.Email, token)
	if err != nil {
		return nil, err
	}
	h.events.SendVerification(user.SendVerificationEvent{
		Email: u.Email,
		Token: token,
		Lang:  cmd.Lang,
	})
	return &ReSendVerificationResult{}, nil
}
