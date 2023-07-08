package user

import (
	"fmt"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.auth/src/config"
	"github.com/turistikrota/service.shared/events"
	"github.com/turistikrota/service.shared/helper"
)

type Events interface {
	UserVerified(event UserVerifiedEvent)
	SendVerification(event SendVerificationEvent)
}

type (
	UserVerifiedEvent struct {
		UserUUID string `json:"userUUID"`
		User     User   `json:"user"`
	}
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

func (e userEvents) UserVerified(event UserVerifiedEvent) {
	_ = e.publisher.Publish(e.topics.Auth.Registered, event.User)
}

func (e userEvents) SendVerification(event SendVerificationEvent) {
	subject := e.i18n.Translate(I18nMessages.AuthRegisteredMailSubject, event.Lang)
	template := fmt.Sprintf("auth/registered.%s", event.Lang)
	_ = e.publisher.Publish(e.topics.Notify.SendMail, helper.Notify.BuildEmail(event.Email, subject, i18np.P{
		"Email": event.Email,
		"Token": event.Token,
	}, event.Email, template))
}
