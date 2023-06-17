package res

import (
	"api.turistikrota.com/auth/src/app/query"
	"github.com/turistikrota/service.shared/jwt"
)

type Response interface {
	LoggedIn(token string) *AuthResponse
	VerifyRequired() *VerifyRequiredResponse
	CurrentUser(u *jwt.UserClaim) *CurrentUserResponse
	CheckEmail(result *query.CheckEmailResult) *CheckEmailResponse
}

type response struct{}

func New() Response {
	return &response{}
}
