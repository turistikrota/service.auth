package event_stream

import (
	"context"

	"github.com/mixarchitecture/microp/events"
	"github.com/sirupsen/logrus"
	"github.com/turistikrota/service.auth/src/app"
	"github.com/turistikrota/service.auth/src/config"
	"github.com/turistikrota/service.auth/src/delivery/event_stream/dto"
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
	_ = s.engine.Subscribe(s.Topics.Account.Restored, s.ListenAccountRestored)
	_ = s.engine.Subscribe(s.Topics.Account.Enabled, s.ListenAccountEnabled)
	_ = s.engine.Subscribe(s.Topics.Account.Disabled, s.ListenAccountDisabled)
	_ = s.engine.Subscribe(s.Topics.Business.Created, s.ListenBusinessCreated)
	_ = s.engine.Subscribe(s.Topics.Business.Enabled, s.ListenBusinessEnabled)
	_ = s.engine.Subscribe(s.Topics.Business.Disabled, s.ListenBusinessDisabled)
	_ = s.engine.Subscribe(s.Topics.Business.DeletedByAdmin, s.ListenBusinessDeleted)
	_ = s.engine.Subscribe(s.Topics.Business.RecoverByAdmin, s.ListenBusinessRecovered)
	_ = s.engine.Subscribe(s.Topics.Business.VerifiedByAdmin, s.ListenBusinessVerified)
	_ = s.engine.Subscribe(s.Topics.Business.UserAdded, s.ListenBusinessUserAdded)
	_ = s.engine.Subscribe(s.Topics.Business.UserRemoved, s.ListenBusinessUserRemoved)
	_ = s.engine.Subscribe(s.Topics.Business.UserPermissionAdded, s.ListenBusinessUserPermissionAdded)
	_ = s.engine.Subscribe(s.Topics.Business.UserPermissionRemoved, s.ListenBusinessUserPermissionRemoved)
	_ = s.engine.Subscribe(s.Topics.Admin.PermissionsAdded, s.ListenUserRolesAdded)
	_ = s.engine.Subscribe(s.Topics.Admin.PermissionsRemoved, s.ListenUserRolesRemoved)
}
