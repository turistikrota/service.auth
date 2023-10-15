package req

import "github.com/turistikrota/service.auth/src/app/command"

type FcmRequest struct {
	Token string `json:"token" validate:"required"`
}

func (r *request) Fcm() *FcmRequest {
	return &FcmRequest{}
}

func (r *FcmRequest) ToCommand(userUUID string, deviceUUID string) command.FcmSetCommand {
	return command.FcmSetCommand{
		UserUUID:   userUUID,
		DeviceUUID: deviceUUID,
		FcmToken:   r.Token,
	}
}
