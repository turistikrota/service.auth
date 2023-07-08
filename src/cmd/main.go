package main

import (
	"context"
	"fmt"

	"github.com/mixarchitecture/i18np"
	"github.com/ssibrahimbas/turnstile"
	"github.com/turistikrota/service.auth/src/config"
	"github.com/turistikrota/service.auth/src/delivery"
	"github.com/turistikrota/service.auth/src/service"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/db/redis"
	"github.com/turistikrota/service.shared/env"
	"github.com/turistikrota/service.shared/events/nats"
	"github.com/turistikrota/service.shared/logs"
)

func main() {
	logs.Init()
	ctx := context.Background()
	config := config.App{}
	env.Load(&config)
	i18n := i18np.New(config.I18n.Fallback)
	i18n.Load(config.I18n.Dir, config.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     config.Nats.Url,
		Streams: config.Nats.Streams,
	})
	fmt.Printf("redis ghost: %s\n", config.Redis.Host)
	r := redis.New(&redis.Config{
		Host:     config.Redis.Host,
		Port:     config.Redis.Port,
		Password: config.Redis.Pw,
		DB:       config.Redis.Db,
	})
	mongo := loadMongo(config)
	fmt.Printf("Redis ping:: %s\n", r.Ping())
	cache := redis.New(&redis.Config{
		Host:     config.CacheRedis.Host,
		Port:     config.CacheRedis.Port,
		Password: config.CacheRedis.Pw,
		DB:       config.CacheRedis.Db,
	})
	tknSrv := token.New(token.Config{
		Expiration:     config.TokenSrv.Expiration,
		PublicKeyFile:  config.RSA.PublicKeyFile,
		PrivateKeyFile: config.RSA.PrivateKeyFile,
	})
	session := session.NewSessionApp(session.Config{
		Redis:       r,
		EventEngine: eventEngine,
		Topic:       config.Session.Topic,
		TokenSrv:    tknSrv,
		Project:     config.TokenSrv.Project,
	})
	turnstileSrv := turnstile.New(turnstile.Config{
		Secret: config.Turnstile.Secret,
	})
	app := service.NewApplication(service.Config{
		App:         config,
		EventEngine: eventEngine,
		TokenSrv:    tknSrv,
		CacheSrv:    cache,
		SessionSrv:  session.Service,
		Mongo:       mongo,
		I18n:        i18n,
	})
	delivery := delivery.New(delivery.Config{
		App:          app,
		Config:       config,
		I18n:         i18n,
		Ctx:          ctx,
		EventEngine:  eventEngine,
		TokenSrv:     tknSrv,
		SessionSrv:   session.Service,
		TurnstileSrv: turnstileSrv,
	})
	delivery.Load()
}

func loadMongo(config config.App) *mongo.DB {
	uri := mongo.CalcMongoUri(mongo.UriParams{
		Host:  config.DB.Auth.Host,
		Port:  config.DB.Auth.Port,
		User:  config.DB.Auth.Username,
		Pass:  config.DB.Auth.Password,
		Db:    config.DB.Auth.Database,
		Query: config.DB.Auth.Query,
	})
	fmt.Printf("Mongo URI: %s", uri)
	d, err := mongo.New(uri, config.DB.Auth.Database)
	if err != nil {
		panic(err)
	}
	return d
}
