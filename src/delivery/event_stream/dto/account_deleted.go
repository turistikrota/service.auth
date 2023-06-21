package dto

import "api.turistikrota.com/auth/src/app/command"

type AccountDeleted struct {
	UserUUID    string `json:"user_uuid"`
	AccountName string `json:"name"`
}

func (e *AccountDeleted) ToCommand() command.AccountDeleteCommand {
	return command.AccountDeleteCommand{
		UserUUID: e.UserUUID,
		Name:     e.AccountName,
	}
}
