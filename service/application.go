package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.auth/app"
	"github.com/turistikrota/service.auth/app/command"
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
}

func NewApplication(config Config) app.Application {
	userFactory := user.NewFactory()
	userRepo := user.NewRepo(config.Mongo.GetCollection(config.App.DB.Auth.Collection), userFactory)
	userEvents := user.NewEvents(user.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
	})
	return app.Application{
		Commands: app.Commands{
			ChangePassword: command.NewChangePasswordHandler(userRepo, userFactory),
			SetFcmToken:    command.NewSetFcmTokenHandler(config.SessionSrv),
			Login:          command.NewLoginHandler(userRepo, userFactory, config.SessionSrv, config.App.Rpc),
		},
		Queries: app.Queries{},
	}
}
