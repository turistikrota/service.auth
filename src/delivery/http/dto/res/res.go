package res

import (
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.auth/src/app/query"
	"github.com/turistikrota/service.auth/src/domain/user"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/jwt"
)

type Response interface {
	LoggedIn(token string) *AuthResponse
	VerifyRequired() *VerifyRequiredResponse
	CurrentUser(u *jwt.UserClaim) *CurrentUserResponse
	CheckEmail(result *query.CheckEmailResult) *CheckEmailResponse
	UserList(res *query.UserListResult) *list.Result[*user.ListEntity]
	SessionList(res *query.SessionListResult) []session.Entity
}

type response struct{}

func New() Response {
	return &response{}
}
