package entity

import (
	"time"

	"github.com/turistikrota/service.auth/src/domain/business"
	"github.com/turistikrota/service.shared/jwt"
)

type MongoBusiness struct {
	UUID         string               `bson:"_id,omitempty"`
	NickName     string               `bson:"nick_name"`
	BusinessType string               `bson:"business_type"`
	Users        []*MongoBusinessUser `bson:"users"`
	IsEnabled    bool                 `bson:"is_enabled"`
	IsVerified   bool                 `bson:"is_verified"`
	IsDeleted    bool                 `bson:"is_deleted"`
}

type MongoBusinessUser struct {
	UUID   string    `bson:"uuid"`
	Name   string    `bson:"name"`
	Roles  []string  `bson:"roles"`
	JoinAt time.Time `bson:"join_at"`
}

func (m *MongoBusiness) FromBusiness(o *business.Entity) *MongoBusiness {
	m.NickName = o.NickName
	m.BusinessType = string(o.BusinessType)
	m.Users = m.fromBusinessUsers(o.Users)
	m.IsEnabled = o.IsEnabled
	m.IsVerified = o.IsVerified
	return m
}

func (m *MongoBusiness) ToBusiness() *business.Entity {
	e := &business.Entity{
		UUID:         m.UUID,
		NickName:     m.NickName,
		BusinessType: business.Type(m.BusinessType),
		IsEnabled:    m.IsEnabled,
		IsVerified:   m.IsVerified,
	}
	if m.Users != nil {
		e.Users = m.ToBusinessUsers()
	}

	return e
}

func (m *MongoBusiness) ToClaim(userUUID string) jwt.UserClaimBusiness {
	user := m.findUser(userUUID)
	return jwt.UserClaimBusiness{
		UUID:        m.UUID,
		NickName:    m.NickName,
		AccountName: user.Name,
		Roles:       user.Roles,
	}
}

func (m *MongoBusiness) findUser(userUUID string) *MongoBusinessUser {
	for _, user := range m.Users {
		if user.UUID == userUUID {
			return user
		}
	}
	return nil
}

func (m *MongoBusiness) ToBusinessUsers() []business.User {
	var users []business.User
	for _, user := range m.Users {
		users = append(users, user.ToBusinessUser())
	}
	return users
}

func (u *MongoBusinessUser) ToBusinessUser() business.User {
	return business.User{
		UUID:   u.UUID,
		Name:   u.Name,
		Roles:  u.Roles,
		JoinAt: u.JoinAt,
	}
}

func (u *MongoBusinessUser) FromBusinessUser(user *business.User) *MongoBusinessUser {
	u.UUID = user.UUID
	u.Name = user.Name
	u.Roles = user.Roles
	u.JoinAt = user.JoinAt
	return u
}

func (m *MongoBusiness) fromBusinessUsers(users []business.User) []*MongoBusinessUser {
	var mongoUsers []*MongoBusinessUser
	for _, user := range users {
		mongoUsers = append(mongoUsers, &MongoBusinessUser{
			UUID:   user.UUID,
			Name:   user.Name,
			Roles:  user.Roles,
			JoinAt: user.JoinAt,
		})
	}
	return mongoUsers
}
