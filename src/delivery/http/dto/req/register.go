package req

import (
	"api.turistikrota.com/auth/src/app/command"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Privacy  *bool   `json:"privacy" validate:"required,eq=true"`
}

func (r *RegisterRequest) ToCommand(lang string) command.RegisterCommand {
	return command.RegisterCommand{
		Email:    r.Email,
		Password: r.Password,
		Lang:     lang,
	}
}

func (r *request) Register() *RegisterRequest {
	return &RegisterRequest{}
}
