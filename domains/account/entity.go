package account

import "time"

type Entity struct {
	UserUUID   string     `json:"userUuid" bson:"user_uuid"`
	UserName   string     `json:"userName" bson:"user_name"`
	IsActive   bool       `json:"isActive" bson:"is_active"`
	IsDeleted  bool       `json:"isDeleted" bson:"is_deleted"`
	IsVerified bool       `json:"isVerified" bson:"is_verified"`
	BirthDate  *time.Time `json:"birthDate" bson:"birth_date"`
}
