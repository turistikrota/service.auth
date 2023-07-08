package dto

import "github.com/turistikrota/service.auth/src/app/command"

type OwnerUpdated struct {
	NickName string `json:"nickName"`
}

func (e *OwnerUpdated) ToDisableCommand() command.OwnerDisableCommand {
	return command.OwnerDisableCommand{
		NickName: e.NickName,
	}
}

func (e *OwnerUpdated) ToEnableCommand() command.OwnerEnableCommand {
	return command.OwnerEnableCommand{
		NickName: e.NickName,
	}
}

func (e *OwnerUpdated) ToDeleteCommand() command.OwnerDeleteCommand {
	return command.OwnerDeleteCommand{
		NickName: e.NickName,
	}
}

func (e *OwnerUpdated) ToRecoverCommand() command.OwnerRecoverCommand {
	return command.OwnerRecoverCommand{
		NickName: e.NickName,
	}
}

func (e *OwnerUpdated) ToVerifyCommand() command.OwnerVerifyCommand {
	return command.OwnerVerifyCommand{
		NickName: e.NickName,
	}
}
