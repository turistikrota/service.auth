package delivery

import (
	"context"

	"api.turistikrota.com/auth/src/app"
	"api.turistikrota.com/auth/src/config"
	"api.turistikrota.com/auth/src/delivery/event_stream"
	"api.turistikrota.com/auth/src/delivery/http"
	"api.turistikrota.com/shared/auth/session"
	"api.turistikrota.com/shared/auth/token"
	"api.turistikrota.com/shared/events"
	sharedHttp "api.turistikrota.com/shared/server/http"
	"api.turistikrota.com/shared/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
	"github.com/ssibrahimbas/turnstile"
)

type Delivery interface {
	Load()
}

type delivery struct {
	app          app.Application
	config       config.App
	i18n         *i18np.I18n
	ctx          context.Context
	eventEngine  events.Engine
	tknSrv       token.Service
	sessionSrv   session.Service
	turnstileSrv turnstile.Service
}

type Config struct {
	App          app.Application
	Config       config.App
	I18n         *i18np.I18n
	Ctx          context.Context
	EventEngine  events.Engine
	TokenSrv     token.Service
	SessionSrv   session.Service
	TurnstileSrv turnstile.Service
}

func New(config Config) Delivery {
	return &delivery{
		app:          config.App,
		config:       config.Config,
		i18n:         config.I18n,
		ctx:          config.Ctx,
		eventEngine:  config.EventEngine,
		tknSrv:       config.TokenSrv,
		sessionSrv:   config.SessionSrv,
		turnstileSrv: config.TurnstileSrv,
	}
}

func (d *delivery) Load() {
	d.loadEventStream().loadHTTP()
}

func (d *delivery) loadHTTP() *delivery {
	sharedHttp.RunServer(sharedHttp.Config{
		Host:  d.config.Server.Host,
		Port:  d.config.Server.Port,
		Group: d.config.Server.Group,
		I18n:  d.i18n,
		CreateHandler: func(router fiber.Router) fiber.Router {
			val := validator.New(d.i18n)
			val.ConnectCustom()
			val.RegisterTagName()
			return http.New(http.Config{
				App:          d.app,
				I18n:         *d.i18n,
				Validator:    *val,
				Context:      d.ctx,
				Config:       d.config,
				TokenSrv:     d.tknSrv,
				SessionSrv:   d.sessionSrv,
				TurnstileSrv: d.turnstileSrv,
			}).Load(router)
		},
	})
	return d
}

func (d *delivery) loadEventStream() *delivery {
	eventStream := event_stream.New(event_stream.Config{
		App:    d.app,
		Topics: d.config.Topics,
		Engine: d.eventEngine,
		Ctx:    d.ctx,
	})
	err := d.eventEngine.Open()
	if err != nil {
		panic(err)
	}
	eventStream.Load()
	return d
}
