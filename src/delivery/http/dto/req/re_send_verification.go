package req

import "github.com/turistikrota/service.auth/src/app/command"

type ReSendVerificationRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (r *ReSendVerificationRequest) ToCommand(lang string) command.ReSendVerificationCommand {
	return command.ReSendVerificationCommand{
		Email: r.Email,
		Lang:  lang,
	}
}

func (r *request) ReSendVerification() *ReSendVerificationRequest {
	return &ReSendVerificationRequest{}
}
