package config

import "github.com/turistikrota/service.shared/base_roles"

type roles struct {
	base_roles.Roles
	UserList string
}

var Roles = roles{
	Roles:    base_roles.BaseRoles,
	UserList: "user_list",
}
