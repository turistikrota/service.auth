package app

import (
	"github.com/turistikrota/service.auth/app/command"
	"github.com/turistikrota/service.auth/app/query"
)

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
	UserRolesAdd           command.UserRolesAddHandler
	UserRolesRemove        command.UserRolesRemoveHandler
}

type Queries struct {
	CheckEmail  query.CheckEmailHandler
	SessionList query.SessionListHandler
	UserList    query.UserListHandler
}
