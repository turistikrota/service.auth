package dto

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/domain/owner"
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
