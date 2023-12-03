package dto

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessUserAdded struct {
	NickName string         `json:"nickName"`
	User     *business.User `json:"user"`
}

type UserAddType struct{}

func (e *BusinessUserAdded) ToCommand() command.BusinessAddUserCommand {
	return command.BusinessAddUserCommand{
		NickName: e.NickName,
		User:     e.User,
	}
}
