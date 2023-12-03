package dto

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessUserRemoved struct {
	NickName string          `json:"nickName"`
	User     UserDetailEvent `json:"user"`
}

func (e *BusinessUserRemoved) ToCommand() command.BusinessRemoveUserCommand {
	return command.BusinessRemoveUserCommand{
		NickName: e.NickName,
		User: business.UserDetail{
			UUID: e.User.UUID,
			Name: e.User.Name,
			Code: e.User.Code,
		},
	}
}
