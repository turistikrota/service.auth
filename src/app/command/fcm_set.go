package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.shared/auth/session"
)

type FcmSetCommand struct {
	UserUUID   string
	DeviceUUID string
	FcmToken   string
}

type FcmSetResult struct {
	Sessions []session.Entity
}

type FcmSetHandler decorator.CommandHandler[FcmSetCommand, *FcmSetResult]

type fcmSetHandler struct {
	sessionSrv session.Service
}

type FcmSetHandlerConfig struct {
	SessionSrv session.Service
	CqrsBase   decorator.Base
}

func NewFcmSetHandler(config FcmSetHandlerConfig) FcmSetHandler {
	return decorator.ApplyQueryDecorators[FcmSetCommand, *FcmSetResult](
		fcmSetHandler{
			sessionSrv: config.SessionSrv,
		},
		config.CqrsBase,
	)
}

func (h fcmSetHandler) Handle(ctx context.Context, cmd FcmSetCommand) (*FcmSetResult, *i18np.Error) {
	err := h.sessionSrv.SetFcmToken(cmd.UserUUID, cmd.DeviceUUID, cmd.FcmToken)
	if err != nil {
		return nil, err
	}
	return &FcmSetResult{}, nil
}
