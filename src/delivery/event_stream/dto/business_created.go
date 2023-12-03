package dto

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/domain/business"
)

type BusinessCreated struct {
	Business *business.Entity `json:"business"`
}

func (e *BusinessCreated) ToCommand() command.BusinessCreateCommand {
	return command.BusinessCreateCommand{
		Entity: e.Business,
	}
}
