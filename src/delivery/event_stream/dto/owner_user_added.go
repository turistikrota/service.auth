package dto

import (
	"api.turistikrota.com/auth/src/app/command"
	"api.turistikrota.com/auth/src/domain/owner"
)

type OwnerUserAdded struct {
	NickName string      `json:"nickName"`
	User     *owner.User `json:"user"`
}

type UserAddType struct{}

func (e *OwnerUserAdded) ToCommand() command.OwnerAddUserCommand {
	return command.OwnerAddUserCommand{
		NickName: e.NickName,
		User:     e.User,
	}
}
