package dto

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessPermissionEvent struct {
	NickName   string          `json:"nickName"`
	User       UserDetailEvent `json:"user"`
	Permission string          `json:"permission"`
}

type UserDetailEvent struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (e *BusinessPermissionEvent) ToAddCommand() command.BusinessAddUserPermissionCommand {
	return command.BusinessAddUserPermissionCommand{
		NickName:   e.NickName,
		Permission: e.Permission,
		User: business.UserDetail{
			UUID: e.User.UUID,
			Name: e.User.Name,
			Code: e.User.Code,
		},
	}
}

func (e *BusinessPermissionEvent) ToRemoveCommand() command.BusinessRemoveUserPermissionCommand {
	return command.BusinessRemoveUserPermissionCommand{
		NickName:   e.NickName,
		Permission: e.Permission,
		User: business.UserDetail{
			UUID: e.User.UUID,
			Name: e.User.Name,
			Code: e.User.Code,
		},
	}
}
