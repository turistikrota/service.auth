package req

import (
	"api.turistikrota.com/auth/src/app/command"
)

type RefreshTokenRequest struct {
	RefreshToken string
	UserUUID     string
	AccessToken  string
	DeviceUUID   string
	IpAddress    string
}

func (r *RefreshTokenRequest) ToCommand() command.RefreshTokenCommand {
	return command.RefreshTokenCommand{
		RefreshToken: r.RefreshToken,
		AccessToken:  r.AccessToken,
		DeviceUUID:   r.DeviceUUID,
		UserUUID:     r.UserUUID,
		IpAddress:    r.IpAddress,
	}
}

func (r *request) RefreshToken() *RefreshTokenRequest {
	return &RefreshTokenRequest{}
}
