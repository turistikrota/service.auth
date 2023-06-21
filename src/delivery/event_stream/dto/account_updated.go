package dto

import (
	"time"

	"api.turistikrota.com/auth/src/app/command"
)

type AccountUpdated struct {
	UserUUID    string              `json:"user_uuid"`
	AccountName string              `json:"name"`
	Entity      AccountUpdateEntity `json:"entity"`
}

type AccountUpdateEntity struct {
	UserName  string     `json:"user_name"`
	BirthDate *time.Time `json:"birth_date"`
}

func (e *AccountUpdated) ToCommand() command.AccountUpdateCommand {
	return command.AccountUpdateCommand{
		UserUUID:    e.UserUUID,
		CurrentName: e.AccountName,
		NewName:     e.Entity.UserName,
		BirthDate:   e.Entity.BirthDate,
	}
}
