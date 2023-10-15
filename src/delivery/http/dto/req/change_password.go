package req

import "github.com/turistikrota/service.auth/src/app/command"

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,password"`
	NewPassword string `json:"new_password" validate:"required,password"`
}

func (r *ChangePasswordRequest) ToCommand(userUUID string) command.ChangePasswordCommand {
	return command.ChangePasswordCommand{
		OldPassword: r.OldPassword,
		NewPassword: r.NewPassword,
		UserUUID:    userUUID,
	}
}

func (r *request) ChangePassword() *ChangePasswordRequest {
	return &ChangePasswordRequest{}
}
