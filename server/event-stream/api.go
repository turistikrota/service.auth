package event_stream

import (
	"context"
	"encoding/json"

	"github.com/turistikrota/service.auth/app/command"
)

func (s srv) OnUserRoleAdd(data []byte) {
	cmd := command.UserRolesAddCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.UserRolesAdd(context.TODO(), cmd)
}

func (s srv) OnUserRoleRemove(data []byte) {
	cmd := command.UserRolesRemoveCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.UserRolesRemove(context.TODO(), cmd)
}
