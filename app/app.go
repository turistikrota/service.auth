package app

import "github.com/turistikrota/service.auth/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ChangePassword         command.ChangePasswordHandler
	SetFcmToken            command.SetFcmTokenHandler
	Login                  command.LoginHandler
	Logout                 command.LogoutHandler
	ReSendVerificationCode command.ReSendVerificationCodeHandler
	RefreshToken           command.RefreshTokenHandler
	Register               command.RegisterHandler
	SessionDestroyAll      command.SessionDestroyAllHandler
	SessionDestroyOthers   command.SessionDestroyOthersHandler
	SessionDestroy         command.SessionDestroyHandler
	UserDelete             command.UserDeleteHandler
}

type Queries struct{}
