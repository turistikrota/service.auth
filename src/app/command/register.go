package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.auth/src/domain/user"
	"github.com/turistikrota/service.shared/cipher"
	"github.com/turistikrota/service.shared/decorator"
)

type RegisterCommand struct {
	Email    string
	Password string
	Lang     string
}

type RegisterResult struct{}

type RegisterHandler decorator.CommandHandler[RegisterCommand, *RegisterResult]

type registerHandler struct {
	repo    user.Repository
	events  user.Events
	factory user.Factory
}

type RegisterHandlerConfig struct {
	Repo     user.Repository
	Events   user.Events
	Factory  user.Factory
	CqrsBase decorator.Base
}

func NewRegisterHandler(config RegisterHandlerConfig) RegisterHandler {
	return decorator.ApplyCommandDecorators[RegisterCommand, *RegisterResult](
		registerHandler{
			repo:    config.Repo,
			events:  config.Events,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h registerHandler) Handle(ctx context.Context, cmd RegisterCommand) (*RegisterResult, *i18np.Error) {
	pw, error := cipher.Hash(cmd.Password)
	if error != nil {
		return nil, h.factory.Errors.Failed("hash")
	}
	token := uuid.New().String()
	u, err := h.repo.Create(ctx, cmd.Email, pw, token)
	if err != nil {
		return nil, err
	}
	h.events.SendVerification(user.SendVerificationEvent{
		Email: u.Email,
		Token: token,
		Lang:  cmd.Lang,
	})
	return &RegisterResult{}, err
}
