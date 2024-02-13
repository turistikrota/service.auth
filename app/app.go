package app

import "github.com/turistikrota/service.auth/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ChangePassword command.ChangePasswordHandler
}

type Queries struct{}
