package service

import (
	"api.turistikrota.com/auth/src/adapters"
	"api.turistikrota.com/auth/src/app"
	"api.turistikrota.com/auth/src/app/command"
	"api.turistikrota.com/auth/src/app/query"
	"api.turistikrota.com/auth/src/config"
	"api.turistikrota.com/auth/src/domain/account"
	"api.turistikrota.com/auth/src/domain/owner"
	"api.turistikrota.com/auth/src/domain/user"
	"api.turistikrota.com/shared/auth/session"
	"api.turistikrota.com/shared/auth/token"
	"api.turistikrota.com/shared/db/mongo"
	"api.turistikrota.com/shared/db/redis"
	"api.turistikrota.com/shared/decorator"
	"api.turistikrota.com/shared/events"
	"github.com/mixarchitecture/i18np"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	TokenSrv    token.Service
	SessionSrv  session.Service
	Mongo       *mongo.DB
	I18n        *i18np.I18n
	CacheSrv    redis.Service
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
		},
	}
}
