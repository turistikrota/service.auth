package req

import (
	"api.turistikrota.com/auth/src/app/command"
	"api.turistikrota.com/shared/auth/session"
)

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,password"`
	DeviceUUID string
	Device     *session.Device
}

func (r *LoginRequest) ToCommand() command.LoginCommand {
	return command.LoginCommand{
		Email:      r.Email,
		Password:   r.Password,
		Device:     r.Device,
		DeviceUUID: r.DeviceUUID,
	}
}

func (r *request) Login() *LoginRequest {
	return &LoginRequest{}
}
