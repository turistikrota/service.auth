package user

import (
	"fmt"

	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/config"
	"github.com/turistikrota/service.auth/domains/notify"
)

type Events interface {
	SendVerification(event SendVerificationEvent)
}

type (
	SendVerificationEvent struct {
		Email string `json:"email"`
		Token string `json:"token"`
		Lang  string
	}
)

type userEvents struct {
	publisher events.Publisher
	topics    config.Topics
	urls      config.Urls
	i18n      *i18np.I18n
}

type EventConfig struct {
	Publisher events.Publisher
	Topics    config.Topics
	Urls      config.Urls
	I18n      *i18np.I18n
}

func NewEvents(cnf EventConfig) Events {
	return &userEvents{
		publisher: cnf.Publisher,
		topics:    cnf.Topics,
		urls:      cnf.Urls,
		i18n:      cnf.I18n,
	}
}

func (e userEvents) SendVerification(event SendVerificationEvent) {
	subject := e.i18n.Translate(I18nMessages.AuthRegisteredMailSubject, event.Lang)
	template := fmt.Sprintf("auth/registered.%s", event.Lang)
	_ = e.publisher.Publish(e.topics.Notify.SendSpecialEmail, notify.SendSpecialEmailCmd{
		Email:    event.Email,
		Template: template,
		Subject:  subject,
		TemplateParams: i18np.P{
			"Email": event.Email,
			"Token": event.Token,
		},
		Locale:    event.Lang,
		Translate: false,
	})
}
