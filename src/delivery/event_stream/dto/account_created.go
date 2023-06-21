package dto

import (
	"api.turistikrota.com/auth/src/app/command"
)

type AccountCreated struct {
	UserUUID    string `json:"user_uuid"`
	AccountName string `json:"name"`
}

func (e *AccountCreated) ToCommand() command.AccountCreateCommand {
	return command.AccountCreateCommand{
		UserUUID:    e.UserUUID,
		AccountName: e.AccountName,
	}
}
