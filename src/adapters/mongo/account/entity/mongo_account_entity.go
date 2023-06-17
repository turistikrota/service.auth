package entity

import (
	"time"

	"api.turistikrota.com/auth/src/domain/account"
	"api.turistikrota.com/shared/jwt"
)

type MongoAccount struct {
	UserUUID   string     `bson:"user_uuid"`
	UserName   string     `bson:"user_name"`
	UserCode   string     `bson:"user_code"`
	IsActive   bool       `bson:"is_active"`
	IsDeleted  bool       `bson:"is_deleted"`
	IsVerified bool       `bson:"is_verified"`
	BirthDate  *time.Time `bson:"birth_date"`
}

func (e *MongoAccount) ToEntity() *account.Entity {
	return &account.Entity{
		UserUUID:   e.UserUUID,
		UserName:   e.UserName,
		UserCode:   e.UserCode,
		IsActive:   e.IsActive,
		IsDeleted:  e.IsDeleted,
		IsVerified: e.IsVerified,
		BirthDate:  e.BirthDate,
	}
}

func (e *MongoAccount) ToClaim() jwt.UserClaimAccount {
	return jwt.UserClaimAccount{
		Name: e.UserName,
		Code: e.UserCode,
	}
}

func (e *MongoAccount) FromEntity(entity *account.Entity) *MongoAccount {
	e.UserUUID = entity.UserUUID
	e.UserName = entity.UserName
	e.UserCode = entity.UserCode
	e.IsActive = entity.IsActive
	e.IsDeleted = entity.IsDeleted
	e.IsVerified = entity.IsVerified
	e.BirthDate = entity.BirthDate
	return e
}
