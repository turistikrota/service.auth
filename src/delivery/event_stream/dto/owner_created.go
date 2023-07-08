package dto

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/domain/owner"
)

type OwnerCreated struct {
	Owner *owner.Entity `json:"owner"`
}

func (e *OwnerCreated) ToCommand() command.OwnerCreateCommand {
	return command.OwnerCreateCommand{
		Entity: e.Owner,
	}
}
