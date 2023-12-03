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
	"github.com/turistikrota/service.auth/src/domain/business"
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

	businessFactory := business.NewFactory()
	businessRepo := adapters.Mongo.NewBusiness(businessFactory, c.Mongo.GetCollection(c.App.DB.Business.Collection))

	base := decorator.NewBase()

	return app.Application{
		Commands: app.Commands{
			Login: command.NewLoginHandler(command.LoginHandlerConfig{
				UserRepo:     userRepo,
				AccountRepo:  accountRepo,
				BusinessRepo: businessRepo,
				AuthTopics:   c.App.Topics.Auth,
				VerifyTopics: c.App.Topics.Verify,
				Publisher:    c.EventEngine,
				TokenSrv:     c.TokenSrv,
				SessionSrv:   c.SessionSrv,
				Errors:       userFactory.Errors,
				CqrsBase:     base,
			}),
			LoginVerified: command.NewLoginVerifiedHandler(command.LoginVerifiedHandlerConfig{
				AuthTopics:   c.App.Topics.Auth,
				AccountRepo:  accountRepo,
				BusinessRepo: businessRepo,
				Publisher:    c.EventEngine,
				TokenSrv:     c.TokenSrv,
				SessionSrv:   c.SessionSrv,
				Errors:       userFactory.Errors,
				Repo:         userRepo,
				CqrsBase:     base,
			}),
			Register: command.NewRegisterHandler(command.RegisterHandlerConfig{
				Repo:     userRepo,
				Events:   userEvents,
				Factory:  userFactory,
				CqrsBase: base,
			}),
			RefreshToken: command.NewRefreshTokenHandler(command.RefreshTokenHandlerConfig{
				AuthTopics:   c.App.Topics.Auth,
				AccountRepo:  accountRepo,
				BusinessRepo: businessRepo,
				Publisher:    c.EventEngine,
				TokenSrv:     c.TokenSrv,
				SessionSrv:   c.SessionSrv,
				Errors:       userFactory.Errors,
				UserRepo:     userRepo,
				CqrsBase:     base,
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
			UserDelete: command.NewUserDeleteHandler(command.UserDeleteHandlerConfig{
				Repo:       userRepo,
				CqrsBase:   base,
				SessionSrv: c.SessionSrv,
			}),
			ChangePassword: command.NewChangePasswordHandler(command.ChangePasswordHandlerConfig{
				UserRepo: userRepo,
				Errors:   userFactory.Errors,
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
			BusinessCreate: command.NewBusinessCreateHandler(command.BusinessCreateHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				CqrsBase: base,
			}),
			BusinessAddUser: command.NewBusinessAddUserHandler(command.BusinessAddUserHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			BusinessVerify: command.NewBusinessVerifyHandler(command.BusinessVerifyHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			BusinessRemoveUser: command.NewBusinessRemoveUserHandler(command.BusinessRemoveUserHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			BusinessAddUserPermission: command.NewBusinessAddUserPermissionHandler(command.BusinessAddUserPermissionHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			BusinessRemoveUserPermission: command.NewBusinessRemoveUserPermissionHandler(command.BusinessRemoveUserPermissionHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			BusinessEnable: command.NewBusinessEnableHandler(command.BusinessEnableHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			BusinessDisable: command.NewBusinessDisableHandler(command.BusinessDisableHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			BusinessDelete: command.NewBusinessDeleteHandler(command.BusinessDeleteHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			BusinessRecover: command.NewBusinessRecoverHandler(command.BusinessRecoverHandlerConfig{
				Repo:     businessRepo,
				CqrsBase: base,
			}),
			UserRolesAdd: command.NewUserRolesAddHandler(command.UserRolesAddHandlerConfig{
				Repo:     userRepo,
				CqrsBase: base,
			}),
			UserRolesRemove: command.NewUserRolesRemoveHandler(command.UserRolesRemoveHandlerConfig{
				Repo:     userRepo,
				CqrsBase: base,
			}),
			SessionDestroy: command.NewSessionDestroyHandler(command.SessionDestroyHandlerConfig{
				SessionSrv: c.SessionSrv,
				CqrsBase:   base,
			}),
			SessionDestroyOthers: command.NewSessionDestroyOthersHandler(command.SessionDestroyOthersHandlerConfig{
				SessionSrv: c.SessionSrv,
				CqrsBase:   base,
			}),
			SessionDestroyAll: command.NewSessionDestroyAllHandler(command.SessionDestroyAllHandlerConfig{
				SessionSrv: c.SessionSrv,
				CqrsBase:   base,
			}),
			FcmSet: command.NewFcmSetHandler(command.FcmSetHandlerConfig{
				SessionSrv: c.SessionSrv,
				CqrsBase:   base,
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
			SessionList: query.NewSessionListHandler(query.SessionListHandlerConfig{
				SessionSrv: c.SessionSrv,
				CqrsBase:   base,
			}),
		},
	}
}
