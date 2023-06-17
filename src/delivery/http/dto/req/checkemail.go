package req

import (
	"api.turistikrota.com/auth/src/app/query"
)

type CheckEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (r *CheckEmailRequest) ToQuery() query.CheckEmailQuery {
	return query.CheckEmailQuery{
		Email: r.Email,
	}
}

func (r *request) CheckEmail() *CheckEmailRequest {
	return &CheckEmailRequest{}
}
