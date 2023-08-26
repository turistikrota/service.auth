package dto

import "github.com/turistikrota/service.auth/src/app/command"

type UserRolesEvent struct {
	UUID        string   `json:"uuid"`
	Permissions []string `json:"permissions"`
}

func (e *UserRolesEvent) ToAddCommand() command.UserRolesAddCommand {
	return command.UserRolesAddCommand{
		UserUUID:    e.UUID,
		Permissions: e.Permissions,
	}
}

func (e *UserRolesEvent) ToRemoveCommand() command.UserRolesRemoveCommand {
	return command.UserRolesRemoveCommand{
		UserUUID:    e.UUID,
		Permissions: e.Permissions,
	}
}
