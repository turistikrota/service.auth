package res

import (
	"github.com/turistikrota/service.shared/helper"
	"github.com/turistikrota/service.shared/jwt"
)

type (
	CurrentUserAccount struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}
)

type CurrentUserResponse struct {
	UUID     string               `json:"uuid"`
	Email    string               `json:"email"`
	Roles    []string             `json:"roles"`
	Accounts []CurrentUserAccount `json:"accounts"`
	Owners   []jwt.UserClaimOwner `json:"owners"`
}

func (r *response) CurrentUser(u *jwt.UserClaim) *CurrentUserResponse {
	return &CurrentUserResponse{
		UUID:     u.UUID,
		Email:    u.Email,
		Roles:    u.Roles,
		Accounts: r.CurrentUserAccount(u),
		Owners:   u.Owners,
	}
}

func (r *response) CurrentUserAccount(u *jwt.UserClaim) []CurrentUserAccount {
	accounts := make([]CurrentUserAccount, 0)
	for _, a := range u.Accounts {
		accounts = append(accounts, CurrentUserAccount{
			Name:   a.Name,
			Avatar: helper.CDN.DressUser(a.Name),
		})
	}
	return accounts
}
