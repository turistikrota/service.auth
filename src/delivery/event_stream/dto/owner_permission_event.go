package dto

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/domain/owner"
)

type OwnerPermissionEvent struct {
	NickName   string          `json:"nickName"`
	User       UserDetailEvent `json:"user"`
	Permission string          `json:"permission"`
}

type UserDetailEvent struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (e *OwnerPermissionEvent) ToAddCommand() command.OwnerAddUserPermissionCommand {
	return command.OwnerAddUserPermissionCommand{
		NickName:   e.NickName,
		Permission: e.Permission,
		User: owner.UserDetail{
			UUID: e.User.UUID,
			Name: e.User.Name,
			Code: e.User.Code,
		},
	}
}

func (e *OwnerPermissionEvent) ToRemoveCommand() command.OwnerRemoveUserPermissionCommand {
	return command.OwnerRemoveUserPermissionCommand{
		NickName:   e.NickName,
		Permission: e.Permission,
		User: owner.UserDetail{
			UUID: e.User.UUID,
			Name: e.User.Name,
			Code: e.User.Code,
		},
	}
}
