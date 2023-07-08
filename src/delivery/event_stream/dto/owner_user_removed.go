package dto

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/domain/owner"
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
