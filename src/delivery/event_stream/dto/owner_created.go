package dto

import (
	"api.turistikrota.com/auth/src/app/command"
	"api.turistikrota.com/auth/src/domain/owner"
)

type OwnerCreated struct {
	Owner *owner.Entity `json:"owner"`
}

func (e *OwnerCreated) ToCommand() command.OwnerCreateCommand {
	return command.OwnerCreateCommand{
		Entity: e.Owner,
	}
}
