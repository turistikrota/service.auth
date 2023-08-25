package service

import (
	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.auth/src/adapters"
	"github.com/turistikrota/service.auth/src/app"
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/app/query"
	"github.com/turistikrota/service.auth/src/config"
	"github.com/turistikrota/service.auth/src/domain/account"
	"github.com/turistikrota/service.auth/src/domain/owner"
	"github.com/turistikrota/service.auth/src/domain/user"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	TokenSrv    token.Service
	SessionSrv  session.Service
	Mongo       *mongo.DB
	I18n        *i18np.I18n
	CacheSrv    cache.Service
}

func NewApplication(c Config) app.Application {
	userFactory := user.NewFactory()
	userRepo := adapters.Mongo.NewUser(userFactory, c.Mongo.GetCollection(c.App.DB.Auth.Collection))
	userEvents := user.NewEvents(user.EventConfig{
		Publisher: c.EventEngine,
		Topics:    c.App.Topics,
		Urls:      c.App.Urls,
		I18n:      c.I18n,
	})

	accountFactory := account.NewFactory()
	accountRepo := adapters.Mongo.NewAccount(accountFactory, c.Mongo.GetCollection(c.App.DB.Account.Collection))

	ownerFactory := owner.NewFactory()
	ownerRepo := adapters.Mongo.NewOwner(ownerFactory, c.Mongo.GetCollection(c.App.DB.Owner.Collection))

	base := decorator.NewBase()

	return app.Application{
		Commands: app.Commands{
			Login: command.NewLoginHandler(command.LoginHandlerConfig{
				UserRepo:     userRepo,
				AccountRepo:  accountRepo,
				OwnerRepo:    ownerRepo,
				AuthTopics:   c.App.Topics.Auth,
				VerifyTopics: c.App.Topics.Verify,
				Publisher:    c.EventEngine,
				TokenSrv:     c.TokenSrv,
				SessionSrv:   c.SessionSrv,
				Errors:       userFactory.Errors,
				CqrsBase:     base,
			}),
			LoginVerified: command.NewLoginVerifiedHandler(command.LoginVerifiedHandlerConfig{
				AuthTopics:  c.App.Topics.Auth,
				AccountRepo: accountRepo,
				OwnerRepo:   ownerRepo,
				Publisher:   c.EventEngine,
				TokenSrv:    c.TokenSrv,
				SessionSrv:  c.SessionSrv,
				Errors:      userFactory.Errors,
				Repo:        userRepo,
				CqrsBase:    base,
			}),
			Register: command.NewRegisterHandler(command.RegisterHandlerConfig{
				Repo:     userRepo,
				Events:   userEvents,
				Factory:  userFactory,
				CqrsBase: base,
			}),
			RefreshToken: command.NewRefreshTokenHandler(command.RefreshTokenHandlerConfig{
				AuthTopics: c.App.Topics.Auth,
				Publisher:  c.EventEngine,
				TokenSrv:   c.TokenSrv,
				SessionSrv: c.SessionSrv,
				Errors:     userFactory.Errors,
				UserRepo:   userRepo,
				CqrsBase:   base,
			}),
			Logout: command.NewLogoutHandler(command.LogoutHandlerConfig{
				AuthTopics: c.App.Topics.Auth,
				Publisher:  c.EventEngine,
				TokenSrv:   c.TokenSrv,
				SessionSrv: c.SessionSrv,
				Errors:     userFactory.Errors,
				CqrsBase:   base,
			}),
			UserUpdated: command.NewUserUpdatedHandler(command.UserUpdatedHandlerConfig{
				Repo:     userRepo,
				CqrsBase: base,
			}),
			Verify: command.NewVerifyHandler(command.VerifyHandlerConfig{
				Repo:     userRepo,
				Factory:  userFactory,
				Events:   userEvents,
				CqrsBase: base,
			}),
			ReSendVerification: command.NewReSendVerificationHandler(command.ReSendVerificationHandlerConfig{
				Repo:     userRepo,
				Factory:  userFactory,
				Events:   userEvents,
				CqrsBase: base,
			}),
			AccountCreate: command.NewAccountCreateHandler(command.AccountCreateHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
			}),
			AccountUpdate: command.NewAccountUpdateHandler(command.AccountUpdateHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
			}),
			AccountDelete: command.NewAccountDeleteHandler(command.AccountDeleteHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
			}),
			AccountDisable: command.NewAccountDisableHandler(command.AccountDisableHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
			}),
			AccountEnable: command.NewAccountEnableHandler(command.AccountEnableHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
			}),
			OwnerCreate: command.NewOwnerCreateHandler(command.OwnerCreateHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				CqrsBase: base,
			}),
			OwnerAddUser: command.NewOwnerAddUserHandler(command.OwnerAddUserHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
			OwnerVerify: command.NewOwnerVerifyHandler(command.OwnerVerifyHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
			OwnerRemoveUser: command.NewOwnerRemoveUserHandler(command.OwnerRemoveUserHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
			OwnerAddUserPermission: command.NewOwnerAddUserPermissionHandler(command.OwnerAddUserPermissionHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
			OwnerRemoveUserPermission: command.NewOwnerRemoveUserPermissionHandler(command.OwnerRemoveUserPermissionHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
			OwnerEnable: command.NewOwnerEnableHandler(command.OwnerEnableHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
			OwnerDisable: command.NewOwnerDisableHandler(command.OwnerDisableHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
			OwnerDelete: command.NewOwnerDeleteHandler(command.OwnerDeleteHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
			OwnerRecover: command.NewOwnerRecoverHandler(command.OwnerRecoverHandlerConfig{
				Repo:     ownerRepo,
				CqrsBase: base,
			}),
		},
		Queries: app.Queries{
			CheckEmail: query.NewCheckEmailHandler(query.CheckEmailHandlerConfig{
				Repo:     userRepo,
				CqrsBase: base,
				CacheSrv: c.CacheSrv,
			}),
			UserList: query.NewUserListHandler(query.UserListHandlerConfig{
				Repo:     userRepo,
				CqrsBase: base,
				CacheSrv: c.CacheSrv,
			}),
		},
	}
}
