package app

import "github.com/turistikrota/service.auth/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ChangePassword command.ChangePasswordHandler
	SetFcmToken    command.SetFcmTokenHandler
	Login          command.LoginHandler
}

type Queries struct{}
