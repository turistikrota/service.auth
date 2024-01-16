package dto

import "github.com/turistikrota/service.auth/src/app/command"

type AccountRestored struct {
	UserUUID    string `json:"user_uuid"`
	AccountName string `json:"name"`
}

func (e *AccountRestored) ToCommand() command.AccountRestoreCommand {
	return command.AccountRestoreCommand{
		UserUUID: e.UserUUID,
		Name:     e.AccountName,
	}
}
