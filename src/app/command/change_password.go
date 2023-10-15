package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/user"
	"github.com/turistikrota/service.shared/cipher"
)

type ChangePasswordCommand struct {
	OldPassword string
	NewPassword string
	UserUUID    string
}

type ChangePasswordResult struct {
}

type ChangePasswordHandler decorator.CommandHandler[ChangePasswordCommand, *ChangePasswordResult]

type changePasswordHandler struct {
	userRepo user.Repository
	errors   user.Errors
}

type ChangePasswordHandlerConfig struct {
	UserRepo user.Repository
	Errors   user.Errors
	CqrsBase decorator.Base
}

func NewChangePasswordHandler(config ChangePasswordHandlerConfig) ChangePasswordHandler {
	return decorator.ApplyCommandDecorators[ChangePasswordCommand, *ChangePasswordResult](
		changePasswordHandler{
			userRepo: config.UserRepo,
			errors:   config.Errors,
		},
		config.CqrsBase,
	)
}

func (h changePasswordHandler) Handle(ctx context.Context, cmd ChangePasswordCommand) (*ChangePasswordResult, *i18np.Error) {
	user, err := h.userRepo.GetByUUID(ctx, cmd.UserUUID)
	if err != nil {
		return nil, err
	}
	if err := cipher.Compare(user.Password, cmd.OldPassword); err != nil {
		return nil, h.errors.InvalidPassword()
	}
	err = h.userRepo.SetPassword(ctx, cmd.UserUUID, []byte(cmd.NewPassword))
	if err != nil {
		return nil, err
	}
	return &ChangePasswordResult{}, nil
}
