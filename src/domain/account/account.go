package account

import "time"

type Entity struct {
	UserUUID   string     `json:"user_uuid"`
	UserName   string     `json:"user_name"`
	UserCode   string     `json:"user_code"`
	IsActive   bool       `json:"is_active"`
	IsDeleted  bool       `json:"is_deleted"`
	IsVerified bool       `json:"is_verified"`
	BirthDate  *time.Time `json:"birth_date"`
}
