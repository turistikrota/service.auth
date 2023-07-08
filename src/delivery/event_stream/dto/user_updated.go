package dto

import "github.com/turistikrota/service.auth/src/app/command"

type UserUpdated struct {
	UserUUID         string `json:"userUUID"`
	TwoFactorEnabled bool   `json:"twoFactorEnabled"`
}

func (u *UserUpdated) ToCommand() command.UserUpdatedCommand {
	return command.UserUpdatedCommand{
		UserUUID:         u.UserUUID,
		TwoFactorEnabled: u.TwoFactorEnabled,
	}
}
