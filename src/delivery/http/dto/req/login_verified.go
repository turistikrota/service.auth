package req

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.shared/auth/session"
)

type LoginVerifiedRequest struct {
	DeviceUUID string
	UserUUID   string
	Device     *session.Device
}

func (r *LoginVerifiedRequest) ToCommand() command.LoginVerifiedCommand {
	return command.LoginVerifiedCommand{
		UserUUID:   r.UserUUID,
		Device:     r.Device,
		DeviceUUID: r.DeviceUUID,
	}
}

func (r *request) LoginVerified() *LoginVerifiedRequest {
	return &LoginVerifiedRequest{}
}
