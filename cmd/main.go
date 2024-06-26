package main

import (
	"fmt"

	"github.com/cilloparch/cillop/env"
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/events/nats"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/ssibrahimbas/turnstile"
	"github.com/turistikrota/service.auth/config"
	"github.com/turistikrota/service.auth/domains/notify"
	"github.com/turistikrota/service.auth/domains/user"
	event_stream "github.com/turistikrota/service.auth/server/event-stream"
	"github.com/turistikrota/service.auth/server/http"
	"github.com/turistikrota/service.auth/service"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/auth/verify"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/db/redis"
)

func main() {
	cnf := config.App{}
	env.Load(&cnf)
	i18n := i18np.New(cnf.I18n.Fallback)
	i18n.Load(cnf.I18n.Dir, cnf.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     cnf.Nats.Url,
		Streams: cnf.Nats.Streams,
	})
	valid := validation.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	mongo := loadMongo(cnf)
	r := redis.New(&redis.Config{
		Host:     cnf.Redis.Host,
		Port:     cnf.Redis.Port,
		Password: cnf.Redis.Pw,
		DB:       cnf.Redis.Db,
	})
	verifyRedis := redis.New(&redis.Config{
		Host:     cnf.CacheRedis.Host,
		Port:     cnf.CacheRedis.Port,
		Password: cnf.CacheRedis.Pw,
		DB:       cnf.VerifyRedis.DB,
	})
	tknSrv := token.New(token.Config{
		Expiration:     cnf.TokenSrv.Expiration,
		PublicKeyFile:  cnf.RSA.PublicKeyFile,
		PrivateKeyFile: cnf.RSA.PrivateKeyFile,
	})
	session := session.NewSessionApp(session.Config{
		Redis:       r,
		EventEngine: eventEngine,
		TokenSrv:    tknSrv,
		Topic:       cnf.Session.Topic,
		Project:     cnf.TokenSrv.Project,
	})
	verifySrv := verify.New(verify.Config{
		Redis:        verifyRedis,
		SessionSrv:   session.Service,
		NotifySender: createNotificationSender(i18n, eventEngine, cnf),
	})
	app := service.NewApplication(service.Config{
		App:         cnf,
		EventEngine: eventEngine,
		Validator:   valid,
		SessionSrv:  session.Service,
		Mongo:       mongo,
		I18n:        i18n,
	})
	turnstileSrv := turnstile.New(turnstile.Config{
		Secret:       cnf.Turnstile.Secret,
		BackupSecret: cnf.Turnstile.MobileSecret,
	})
	http := http.New(http.Config{
		Env:          cnf,
		App:          app,
		I18n:         i18n,
		Validator:    *valid,
		HttpHeaders:  cnf.HttpHeaders,
		TokenSrv:     tknSrv,
		SessionSrv:   session.Service,
		TurnstileSrv: turnstileSrv,
		VerifySrv:    verifySrv,
	})
	eventStream := event_stream.New(event_stream.Config{
		App:    app,
		Engine: eventEngine,
		Topics: cnf.Topics,
	})
	go eventStream.Listen()
	http.Listen()
}

func loadMongo(cnf config.App) *mongo.DB {
	uri := mongo.CalcMongoUri(mongo.UriParams{
		Host:  cnf.DB.Auth.Host,
		Port:  cnf.DB.Auth.Port,
		User:  cnf.DB.Auth.Username,
		Pass:  cnf.DB.Auth.Password,
		Db:    cnf.DB.Auth.Database,
		Query: cnf.DB.Auth.Query,
	})
	d, err := mongo.New(uri, cnf.DB.Auth.Database)
	if err != nil {
		panic(err)
	}
	return d
}

func createNotificationSender(i18n *i18np.I18n, events events.Engine, cnf config.App) verify.NotifySender {
	return func(cmd verify.NotifyCommand) {
		fmt.Println("before mo or after me")
		if cmd.Phone != "" {
			_ = events.Publish(cnf.Topics.Notify.SendSpecialSms, notify.NotifySendSpecialSmsCmd{
				Phone:     cmd.Phone,
				Text:      fmt.Sprintf(i18n.Translate(user.I18nMessages.SMSVerificationCode, cmd.Locale), cmd.Code),
				Locale:    cmd.Locale,
				Translate: false,
			})
		}
		if cmd.Email != "" {
			subject := i18n.Translate(user.I18nMessages.SubjectVerificationCode, cmd.Locale)
			template := fmt.Sprintf("verify/code.%s", cmd.Locale)
			if cmd.IpAddress == "" {
				cmd.IpAddress = "N/A"
			}
			if cmd.BrowserName == "" {
				cmd.BrowserName = "N/A"
			}
			if cmd.OperatingSystem == "" {
				cmd.OperatingSystem = "N/A"
			}
			_ = events.Publish(cnf.Topics.Notify.SendSpecialEmail, notify.SendSpecialEmailCmd{
				Email:    cmd.Email,
				Template: template,
				Subject:  subject,
				TemplateParams: i18np.P{
					"Code":       cmd.Code,
					"IP":         cmd.IpAddress,
					"Browser":    cmd.BrowserName,
					"OS":         cmd.OperatingSystem,
					"DeviceUUID": cmd.DeviceId,
				},
				Locale:    cmd.Locale,
				Translate: false,
			})
		}
	}
}
