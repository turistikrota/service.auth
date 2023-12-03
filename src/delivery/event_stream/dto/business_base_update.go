package dto

import "github.com/turistikrota/service.auth/src/app/command"

type BusinessUpdated struct {
	NickName string `json:"nickName"`
}

func (e *BusinessUpdated) ToDisableCommand() command.BusinessDisableCommand {
	return command.BusinessDisableCommand{
		NickName: e.NickName,
	}
}

func (e *BusinessUpdated) ToEnableCommand() command.BusinessEnableCommand {
	return command.BusinessEnableCommand{
		NickName: e.NickName,
	}
}

func (e *BusinessUpdated) ToDeleteCommand() command.BusinessDeleteCommand {
	return command.BusinessDeleteCommand{
		NickName: e.NickName,
	}
}

func (e *BusinessUpdated) ToRecoverCommand() command.BusinessRecoverCommand {
	return command.BusinessRecoverCommand{
		NickName: e.NickName,
	}
}

func (e *BusinessUpdated) ToVerifyCommand() command.BusinessVerifyCommand {
	return command.BusinessVerifyCommand{
		NickName: e.NickName,
	}
}
