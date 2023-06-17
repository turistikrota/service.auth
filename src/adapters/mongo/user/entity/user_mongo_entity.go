package entity

import (
	"time"

	"api.turistikrota.com/auth/src/domain/user"
)

type MongoUser struct {
	UUID             string    `bson:"_id,omitempty"`
	Email            string    `bson:"email"`
	Phone            string    `bson:"phone"`
	Roles            []string  `bson:"roles"`
	Password         []byte    `bson:"password"`
	Token            string    `bson:"token"`
	IsVerified       bool      `bson:"is_verified"`
	TwoFactorEnabled bool      `bson:"two_factor_enabled"`
	IsActive         bool      `bson:"is_active"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
}

type MongoIsVerify struct {
	Email bool `bson:"email"`
}

func (m *MongoUser) ToUser() *user.User {
	return &user.User{
		UUID:             m.UUID,
		Email:            m.Email,
		Phone:            m.Phone,
		Roles:            m.Roles,
		Password:         m.Password,
		IsVerified:       m.IsVerified,
		VerifyToken:      m.Token,
		TwoFactorEnabled: m.TwoFactorEnabled,
		IsActive:         m.IsActive,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}

func (m *MongoUser) FromUser(user *user.User) *MongoUser {
	if len(user.UUID) > 0 {
		m.UUID = user.UUID
	}
	m.Roles = user.Roles
	m.Email = user.Email
	m.Phone = user.Phone
	m.TwoFactorEnabled = user.TwoFactorEnabled
	m.Password = user.Password
	m.Token = user.VerifyToken
	m.IsActive = user.IsActive
	m.CreatedAt = user.CreatedAt
	m.UpdatedAt = user.UpdatedAt
	return m
}
