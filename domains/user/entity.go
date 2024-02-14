package user

import "time"

type Entity struct {
	UUID             string     `json:"uuid" bson:"_id,omitempty"`
	Email            string     `json:"email" bson:"email"`
	Phone            string     `json:"phone" bson:"phone"`
	Roles            []string   `json:"roles" bson:"roles"`
	Password         []byte     `json:"password,omitempty" bson:"password,omitempty"`
	TwoFactorEnabled bool       `json:"twoFactorEnabled" bson:"two_factor_enabled"`
	IsActive         bool       `json:"isActive" bson:"is_active"`
	IsDeleted        bool       `json:"isDeleted" bson:"is_deleted"`
	IsVerified       bool       `json:"isVerified" bson:"is_verified"`
	VerifyToken      string     `json:"emailVerifyToken" bson:"email_verify_token"`
	CreatedAt        time.Time  `json:"createdAt" bson:"created_at"`
	UpdatedAt        time.Time  `json:"updatedAt" bson:"updated_at"`
	DeletedAt        *time.Time `json:"deletedAt" bson:"deleted_at"`
}

func (e *Entity) SetPassword(password []byte) {
	e.Password = password
}

func (e *Entity) CleanPassword() {
	e.Password = nil
}
