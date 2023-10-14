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
	Login                     command.LoginHandler
	LoginVerified             command.LoginVerifiedHandler
	Register                  command.RegisterHandler
	RefreshToken              command.RefreshTokenHandler
	Logout                    command.LogoutHandler
	UserUpdated               command.UserUpdatedHandler
	Verify                    command.VerifyHandler
	ReSendVerification        command.ReSendVerificationHandler
	AccountCreate             command.AccountCreateHandler
	AccountUpdate             command.AccountUpdateHandler
	AccountDelete             command.AccountDeleteHandler
	AccountEnable             command.AccountEnableHandler
	AccountDisable            command.AccountDisableHandler
	OwnerCreate               command.OwnerCreateHandler
	OwnerAddUser              command.OwnerAddUserHandler
	OwnerRemoveUser           command.OwnerRemoveUserHandler
	OwnerAddUserPermission    command.OwnerAddUserPermissionHandler
	OwnerRemoveUserPermission command.OwnerRemoveUserPermissionHandler
	OwnerEnable               command.OwnerEnableHandler
	OwnerDisable              command.OwnerDisableHandler
	OwnerDelete               command.OwnerDeleteHandler
	OwnerRecover              command.OwnerRecoverHandler
	OwnerVerify               command.OwnerVerifyHandler
	UserRolesAdd              command.UserRolesAddHandler
	UserRolesRemove           command.UserRolesRemoveHandler
	UserDelete                command.UserDeleteHandler
	SessionDestroy            command.SessionDestroyHandler
	SessionDestroyOthers      command.SessionDestroyOthersHandler
	SessionDestroyAll         command.SessionDestroyAllHandler
}

type Queries struct {
	CheckEmail  query.CheckEmailHandler
	UserList    query.UserListHandler
	SessionList query.SessionListHandler
}
