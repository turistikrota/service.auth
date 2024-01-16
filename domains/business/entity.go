package business

import (
	"time"

	"github.com/turistikrota/service.shared/jwt"
)

type Entity struct {
	UUID         string `json:"uuid" bson:"uuid"`
	NickName     string `json:"nickName" bson:"nick_name"`
	Users        []User `json:"users" bson:"users"`
	BusinessType Type   `json:"businessType" bson:"business_type"`
	IsVerified   bool   `json:"isVerified" bson:"is_verified"`
	IsDeleted    bool   `json:"isDeleted" bson:"is_deleted"`
	IsEnabled    bool   `json:"isEnabled" bson:"is_enabled"`
}

type User struct {
	UUID   string    `json:"uuid" bson:"uuid"`
	Name   string    `json:"name" bson:"name"`
	Roles  []string  `json:"roles" bson:"roles"`
	JoinAt time.Time `json:"joinAt" bson:"join_at"`
}

type Type string

type businessTypes struct {
	Individual  Type
	Corporation Type
}

var Types = businessTypes{
	Individual:  "individual",
	Corporation: "corporation",
}

func (e *Entity) ToClaim(userUUID string) jwt.UserClaimBusiness {
	user := e.findUser(userUUID)
	return jwt.UserClaimBusiness{
		UUID:        e.UUID,
		NickName:    e.NickName,
		AccountName: user.Name,
		Roles:       user.Roles,
	}
}

func (e *Entity) findUser(userUUID string) *User {
	for _, user := range e.Users {
		if user.UUID == userUUID {
			return &user
		}
	}
	return nil
}
