package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.auth/app"
	"github.com/turistikrota/service.auth/app/command"
	"github.com/turistikrota/service.auth/app/query"
	"github.com/turistikrota/service.auth/config"
	"github.com/turistikrota/service.auth/domains/user"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validation.Validator
	SessionSrv  session.Service
	I18n        *i18np.I18n
}

func NewApplication(config Config) app.Application {
	userFactory := user.NewFactory()
	userRepo := user.NewRepo(config.Mongo.GetCollection(config.App.DB.Auth.Collection), userFactory)
	userEvents := user.NewEvents(user.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
		Urls:      config.App.Urls,
		I18n:      config.I18n,
	})
	return app.Application{
		Commands: app.Commands{
			ChangePassword:         command.NewChangePasswordHandler(userRepo, userFactory),
			SetFcmToken:            command.NewSetFcmTokenHandler(config.SessionSrv),
			Login:                  command.NewLoginHandler(userRepo, userFactory, config.SessionSrv, config.App.Rpc),
			Logout:                 command.NewLogoutHandler(config.SessionSrv),
			ReSendVerificationCode: command.NewReSendVerificationCodeHandler(userRepo, userFactory, userEvents),
			RefreshToken:           command.NewRefreshTokenHandler(config.SessionSrv, userRepo, userFactory, config.App.Rpc),
			Register:               command.NewRegisterHandler(userRepo, userFactory, userEvents),
			SessionDestroyAll:      command.NewSessionDestroyAllHandler(config.SessionSrv),
			SessionDestroyOthers:   command.NewSessionDestroyOthersHandler(config.SessionSrv),
			SessionDestroy:         command.NewSessionDestroyHandler(config.SessionSrv),
			UserDelete:             command.NewUserDeleteHandler(config.SessionSrv, userRepo),
			UserRolesAdd:           command.NewUserRolesAddHandler(userRepo),
			UserRolesRemove:        command.NewUserRolesRemoveHandler(userRepo),
			Verify:                 command.NewVerifyHandler(userRepo, userFactory),
			TwoFactorDisable:       command.NewTwoFactorDisableHandler(userRepo),
			TwoFactorEnable:        command.NewTwoFactorEnableHandler(userRepo),
		},
		Queries: app.Queries{
			CheckEmail:  query.NewCheckEmailHandler(userRepo),
			SessionList: query.NewSessionListHandler(config.SessionSrv),
			UserList:    query.NewUserListHandler(userRepo),
		},
	}
}
