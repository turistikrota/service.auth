package event_stream

import (
	"context"

	"api.turistikrota.com/auth/src/app"
	"api.turistikrota.com/auth/src/config"
	"api.turistikrota.com/auth/src/delivery/event_stream/dto"
	"github.com/sirupsen/logrus"
	"github.com/turistikrota/service.shared/events"
)

type Server struct {
	app    app.Application
	Topics config.Topics
	engine events.Engine
	ctx    context.Context
	dto    dto.Dto
}

type Config struct {
	App    app.Application
	Topics config.Topics
	Engine events.Engine
	Ctx    context.Context
}

func New(config Config) Server {
	return Server{
		app:    config.App,
		engine: config.Engine,
		Topics: config.Topics,
		ctx:    config.Ctx,
		dto:    dto.New(),
	}
}

func (s Server) Load() {
	logrus.Info("Loading event stream server")
	s.engine.Subscribe(s.Topics.Auth.UserUpdated, s.ListenUserUpdated)
	_ = s.engine.Subscribe(s.Topics.Account.Created, s.ListenAccountCreated)
	_ = s.engine.Subscribe(s.Topics.Account.Updated, s.ListenAccountUpdated)
	_ = s.engine.Subscribe(s.Topics.Account.Deleted, s.ListenAccountDeleted)
	_ = s.engine.Subscribe(s.Topics.Account.Enabled, s.ListenAccountEnabled)
	_ = s.engine.Subscribe(s.Topics.Account.Disabled, s.ListenAccountDisabled)
	_ = s.engine.Subscribe(s.Topics.Owner.Created, s.ListenOwnerCreated)
	_ = s.engine.Subscribe(s.Topics.Owner.Enabled, s.ListenOwnerEnabled)
	_ = s.engine.Subscribe(s.Topics.Owner.Disabled, s.ListenOwnerDisabled)
	_ = s.engine.Subscribe(s.Topics.Owner.DeletedByAdmin, s.ListenOwnerDeleted)
	_ = s.engine.Subscribe(s.Topics.Owner.RecoverByAdmin, s.ListenOwnerRecovered)
	_ = s.engine.Subscribe(s.Topics.Owner.VerifiedByAdmin, s.ListenOwnerVerified)
	_ = s.engine.Subscribe(s.Topics.Owner.UserAdded, s.ListenOwnerUserAdded)
	_ = s.engine.Subscribe(s.Topics.Owner.UserRemoved, s.ListenOwnerUserRemoved)
	_ = s.engine.Subscribe(s.Topics.Owner.UserPermissionAdded, s.ListenOwnerUserPermissionAdded)
	_ = s.engine.Subscribe(s.Topics.Owner.UserPermissionRemoved, s.ListenOwnerUserPermissionRemoved)
}
