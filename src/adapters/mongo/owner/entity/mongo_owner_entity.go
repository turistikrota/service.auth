package entity

import (
	"time"

	"api.turistikrota.com/auth/src/domain/owner"
	"api.turistikrota.com/shared/jwt"
)

type MongoOwner struct {
	UUID       string            `bson:"_id,omitempty"`
	NickName   string            `bson:"nick_name"`
	OwnerType  string            `bson:"owner_type"`
	Users      []*MongoOwnerUser `bson:"users"`
	IsEnabled  bool              `bson:"is_enabled"`
	IsVerified bool              `bson:"is_verified"`
	IsDeleted  bool              `bson:"is_deleted"`
}

type MongoOwnerUser struct {
	UUID   string    `bson:"uuid"`
	Name   string    `bson:"name"`
	Code   string    `bson:"code"`
	Roles  []string  `bson:"roles"`
	JoinAt time.Time `bson:"join_at"`
}

func (m *MongoOwner) FromOwner(o *owner.Entity) *MongoOwner {
	m.NickName = o.NickName
	m.OwnerType = string(o.OwnerType)
	m.Users = m.fromOwnerUsers(o.Users)
	m.IsEnabled = o.IsEnabled
	m.IsVerified = o.IsVerified
	return m
}

func (m *MongoOwner) ToOwner() *owner.Entity {
	e := &owner.Entity{
		UUID:       m.UUID,
		NickName:   m.NickName,
		OwnerType:  owner.Type(m.OwnerType),
		IsEnabled:  m.IsEnabled,
		IsVerified: m.IsVerified,
	}
	if m.Users != nil {
		e.Users = m.ToOwnerUsers()
	}

	return e
}

func (m *MongoOwner) ToClaim(userUUID string) jwt.UserClaimOwner {
	user := m.findUser(userUUID)
	return jwt.UserClaimOwner{
		UUID:        m.UUID,
		NickName:    m.NickName,
		AccountName: user.Name,
		AccountCode: user.Code,
		Roles:       user.Roles,
	}
}

func (m *MongoOwner) findUser(userUUID string) *MongoOwnerUser {
	for _, user := range m.Users {
		if user.UUID == userUUID {
			return user
		}
	}
	return nil
}

func (m *MongoOwner) ToOwnerUsers() []owner.User {
	var users []owner.User
	for _, user := range m.Users {
		users = append(users, user.ToOwnerUser())
	}
	return users
}

func (u *MongoOwnerUser) ToOwnerUser() owner.User {
	return owner.User{
		UUID:   u.UUID,
		Name:   u.Name,
		Code:   u.Code,
		Roles:  u.Roles,
		JoinAt: u.JoinAt,
	}
}

func (u *MongoOwnerUser) FromOwnerUser(user *owner.User) *MongoOwnerUser {
	u.UUID = user.UUID
	u.Name = user.Name
	u.Code = user.Code
	u.Roles = user.Roles
	u.JoinAt = user.JoinAt
	return u
}

func (m *MongoOwner) fromOwnerUsers(users []owner.User) []*MongoOwnerUser {
	var mongoUsers []*MongoOwnerUser
	for _, user := range users {
		mongoUsers = append(mongoUsers, &MongoOwnerUser{
			UUID:   user.UUID,
			Name:   user.Name,
			Code:   user.Code,
			Roles:  user.Roles,
			JoinAt: user.JoinAt,
		})
	}
	return mongoUsers
}
