package app

import (
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Login                        command.LoginHandler
	LoginVerified                command.LoginVerifiedHandler
	Register                     command.RegisterHandler
	RefreshToken                 command.RefreshTokenHandler
	Logout                       command.LogoutHandler
	UserUpdated                  command.UserUpdatedHandler
	Verify                       command.VerifyHandler
	ReSendVerification           command.ReSendVerificationHandler
	AccountCreate                command.AccountCreateHandler
	AccountUpdate                command.AccountUpdateHandler
	AccountDelete                command.AccountDeleteHandler
	AccountEnable                command.AccountEnableHandler
	AccountDisable               command.AccountDisableHandler
	BusinessCreate               command.BusinessCreateHandler
	BusinessAddUser              command.BusinessAddUserHandler
	BusinessRemoveUser           command.BusinessRemoveUserHandler
	BusinessAddUserPermission    command.BusinessAddUserPermissionHandler
	BusinessRemoveUserPermission command.BusinessRemoveUserPermissionHandler
	BusinessEnable               command.BusinessEnableHandler
	BusinessDisable              command.BusinessDisableHandler
	BusinessDelete               command.BusinessDeleteHandler
	BusinessRecover              command.BusinessRecoverHandler
	BusinessVerify               command.BusinessVerifyHandler
	UserRolesAdd                 command.UserRolesAddHandler
	UserRolesRemove              command.UserRolesRemoveHandler
	UserDelete                   command.UserDeleteHandler
	SessionDestroy               command.SessionDestroyHandler
	SessionDestroyOthers         command.SessionDestroyOthersHandler
	SessionDestroyAll            command.SessionDestroyAllHandler
	FcmSet                       command.FcmSetHandler
	ChangePassword               command.ChangePasswordHandler
}

type Queries struct {
	CheckEmail  query.CheckEmailHandler
	UserList    query.UserListHandler
	SessionList query.SessionListHandler
}
