package event_stream

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func (s Server) ListenUserUpdated(data []byte) {
	logrus.Info("User updated event received")
	d := s.dto.UserUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	s.app.Commands.UserUpdated.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountCreated(data []byte) {
	logrus.Info("Account created event received")
	d := s.dto.AccountCreated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountCreate.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountUpdated(data []byte) {
	logrus.Info("Account updated event received")
	d := s.dto.AccountUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountUpdate.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountDeleted(data []byte) {
	logrus.Info("Account deleted event received")
	d := s.dto.AccountDeleted()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountDelete.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountEnabled(data []byte) {
	logrus.Info("Account enabled event received")
	d := s.dto.AccountEnabled()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountEnable.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenAccountDisabled(data []byte) {
	logrus.Info("Account disabled event received")
	d := s.dto.AccountDisabled()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.AccountDisable.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenBusinessCreated(data []byte) {
	d := s.dto.BusinessCreated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessCreate.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenBusinessEnabled(data []byte) {
	d := s.dto.BusinessUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessEnable.Handle(s.ctx, d.ToEnableCommand())
}

func (s Server) ListenBusinessDisabled(data []byte) {
	d := s.dto.BusinessUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessDisable.Handle(s.ctx, d.ToDisableCommand())
}

func (s Server) ListenBusinessDeleted(data []byte) {
	d := s.dto.BusinessUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessDelete.Handle(s.ctx, d.ToDeleteCommand())
}

func (s Server) ListenBusinessRecovered(data []byte) {
	d := s.dto.BusinessUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessRecover.Handle(s.ctx, d.ToRecoverCommand())
}

func (s Server) ListenBusinessVerified(data []byte) {
	d := s.dto.BusinessUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessVerify.Handle(s.ctx, d.ToVerifyCommand())
}

func (s Server) ListenBusinessUserAdded(data []byte) {
	d := s.dto.BusinessUserAdded()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessAddUser.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenBusinessUserRemoved(data []byte) {
	d := s.dto.BusinessUserRemoved()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessRemoveUser.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenBusinessUserPermissionAdded(data []byte) {
	d := s.dto.BusinessPermissionEvent()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessAddUserPermission.Handle(s.ctx, d.ToAddCommand())
}

func (s Server) ListenBusinessUserPermissionRemoved(data []byte) {
	d := s.dto.BusinessPermissionEvent()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.BusinessRemoveUserPermission.Handle(s.ctx, d.ToRemoveCommand())
}

func (s Server) ListenUserRolesAdded(data []byte) {
	d := s.dto.UserRolesEvent()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.UserRolesAdd.Handle(s.ctx, d.ToAddCommand())
}

func (s Server) ListenUserRolesRemoved(data []byte) {
	d := s.dto.UserRolesEvent()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.UserRolesRemove.Handle(s.ctx, d.ToRemoveCommand())
}
