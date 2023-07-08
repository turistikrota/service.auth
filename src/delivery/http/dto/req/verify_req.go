package req

import "github.com/turistikrota/service.auth/src/app/command"

type VerifyRequest struct {
	Token string `params:"token" validate:"required,uuid"`
}

func (r *VerifyRequest) ToCommand() command.VerifyCommand {
	return command.VerifyCommand{
		Token: r.Token,
	}
}

func (r *request) Verify() *VerifyRequest {
	return &VerifyRequest{}
}
