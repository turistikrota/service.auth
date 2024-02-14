package event_stream

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/server"
	"github.com/turistikrota/service.auth/app"
	"github.com/turistikrota/service.auth/config"
)

type srv struct {
	app    app.Application
	topics config.Topics
	engine events.Engine
}

type Config struct {
	App    app.Application
	Engine events.Engine
	Topics config.Topics
}

func New(config Config) server.Server {
	return srv{
		app:    config.App,
		engine: config.Engine,
		topics: config.Topics,
	}
}

func (s srv) Listen() error {
	_ = s.engine.Subscribe(s.topics.Admin.PermissionsAdded, s.OnUserRoleAdd)
	_ = s.engine.Subscribe(s.topics.Admin.PermissionsRemoved, s.OnUserRoleRemove)
	return nil
}
