package dto

import (
	"api.turistikrota.com/auth/src/app/command"
	"api.turistikrota.com/auth/src/domain/owner"
)

type OwnerUserRemoved struct {
	NickName string          `json:"nickName"`
	User     UserDetailEvent `json:"user"`
}

func (e *OwnerUserRemoved) ToCommand() command.OwnerRemoveUserCommand {
	return command.OwnerRemoveUserCommand{
		NickName: e.NickName,
		User: owner.UserDetail{
			UUID: e.User.UUID,
			Name: e.User.Name,
			Code: e.User.Code,
		},
	}
}
