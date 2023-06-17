package dto

import "api.turistikrota.com/auth/src/app/command"

type AccountDisabled struct {
	UserUUID    string `json:"user_uuid"`
	AccountName string `json:"name"`
	AccountCode string `json:"code"`
}

func (e *AccountDisabled) ToCommand() command.AccountDisableCommand {
	return command.AccountDisableCommand{
		UserUUID: e.UserUUID,
		Name:     e.AccountName,
		Code:     e.AccountCode,
	}
}
