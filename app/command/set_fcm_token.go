package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.shared/auth/session"
)

type SetFcmTokenCmd struct {
	FcmToken   string `json:"token" validate:"required"`
	UserUUID   string `json:"-"`
	DeviceUUID string `json:"-"`
}

type SetFcmTokenRes struct{}

type SetFcmTokenHandler cqrs.HandlerFunc[SetFcmTokenCmd, *SetFcmTokenRes]

func NewSetFcmTokenHandler(sessionSrv session.Service) SetFcmTokenHandler {
	return func(ctx context.Context, cmd SetFcmTokenCmd) (*SetFcmTokenRes, *i18np.Error) {
		err := sessionSrv.SetFcmToken(ctx, cmd.UserUUID, cmd.DeviceUUID, cmd.FcmToken)
		if err != nil {
			return nil, err
		}
		return &SetFcmTokenRes{}, nil
	}
}
