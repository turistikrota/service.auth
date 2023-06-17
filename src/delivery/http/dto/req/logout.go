package req

import "api.turistikrota.com/auth/src/app/command"

type LogoutRequest struct {
	UserUUID   string
	DeviceUUID string
}

func (r *LogoutRequest) ToCommand() command.LogoutCommand {
	return command.LogoutCommand{
		UserUUID:   r.UserUUID,
		DeviceUUID: r.DeviceUUID,
	}
}

func (r *request) Logout() *LogoutRequest {
	return &LogoutRequest{}
}
