package user

import (
	"time"

	"api.turistikrota.com/shared/auth/session"
)

type User struct {
	UUID             string    `json:"uuid"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	Roles            []string  `json:"roles"`
	Password         []byte    `json:"password,omitempty"`
	TwoFactorEnabled bool      `json:"two_factor_enabled"`
	IsActive         bool      `json:"is_active"`
	IsVerified       bool      `json:"is_verified"`
	VerifyToken      string    `json:"email_verify_token"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (u *User) SetPassword(password []byte) {
	u.Password = password
}

func (u *User) CleanPassword() {
	u.Password = nil
}

func (u *User) ToSession() *session.SessionUser {
	return &session.SessionUser{
		UUID:  u.UUID,
		Email: u.Email,
		Roles: u.Roles,
	}
}
