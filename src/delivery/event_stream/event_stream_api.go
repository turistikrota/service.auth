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

func (s Server) ListenOwnerCreated(data []byte) {
	d := s.dto.OwnerCreated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerCreate.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenOwnerEnabled(data []byte) {
	d := s.dto.OwnerUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerEnable.Handle(s.ctx, d.ToEnableCommand())
}

func (s Server) ListenOwnerDisabled(data []byte) {
	d := s.dto.OwnerUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerDisable.Handle(s.ctx, d.ToDisableCommand())
}

func (s Server) ListenOwnerDeleted(data []byte) {
	d := s.dto.OwnerUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerDelete.Handle(s.ctx, d.ToDeleteCommand())
}

func (s Server) ListenOwnerRecovered(data []byte) {
	d := s.dto.OwnerUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerRecover.Handle(s.ctx, d.ToRecoverCommand())
}

func (s Server) ListenOwnerVerified(data []byte) {
	d := s.dto.OwnerUpdated()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerVerify.Handle(s.ctx, d.ToVerifyCommand())
}

func (s Server) ListenOwnerUserAdded(data []byte) {
	d := s.dto.OwnerUserAdded()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerAddUser.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenOwnerUserRemoved(data []byte) {
	d := s.dto.OwnerUserRemoved()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerRemoveUser.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenOwnerUserPermissionAdded(data []byte) {
	d := s.dto.OwnerPermissionEvent()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerAddUserPermission.Handle(s.ctx, d.ToAddCommand())
}

func (s Server) ListenOwnerUserPermissionRemoved(data []byte) {
	d := s.dto.OwnerPermissionEvent()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.OwnerRemoveUserPermission.Handle(s.ctx, d.ToRemoveCommand())
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
