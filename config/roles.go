package config

import "github.com/turistikrota/service.shared/base_roles"

type roles struct {
	base_roles.Roles
	User userRoles
}

type userRoles struct {
	List string
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	User: userRoles{
		List: "user.list",
	},
}
