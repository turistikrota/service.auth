package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/user"
	"github.com/turistikrota/service.shared/auth/session"
)

type UserDeleteCommand struct {
	UserID     string
	DeviceUUID string
}

type UserDeleteResult struct{}

type UserDeleteHandler decorator.CommandHandler[UserDeleteCommand, *UserDeleteResult]

type userDeleteHandler struct {
	repo       user.Repository
	sessionSrv session.Service
}

type UserDeleteHandlerConfig struct {
	Repo       user.Repository
	SessionSrv session.Service
	CqrsBase   decorator.Base
}

func NewUserDeleteHandler(config UserDeleteHandlerConfig) UserDeleteHandler {
	return decorator.ApplyCommandDecorators[UserDeleteCommand, *UserDeleteResult](
		userDeleteHandler{
			repo:       config.Repo,
			sessionSrv: config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h userDeleteHandler) Handle(ctx context.Context, command UserDeleteCommand) (*UserDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, command.UserID)
	if err != nil {
		return nil, err
	}
	err = h.sessionSrv.Destroy(session.DestroyCommand{
		UserUUID:   command.UserID,
		DeviceUUID: command.DeviceUUID,
	})
	if err != nil {
		return nil, err
	}
	return &UserDeleteResult{}, nil
}
