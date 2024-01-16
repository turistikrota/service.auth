package business

type fieldsType struct {
	UUID         string
	NickName     string
	Users        string
	BusinessType string
	IsVerified   string
	IsDeleted    string
	IsEnabled    string
}

type userFieldsType struct {
	UUID   string
	Name   string
	Roles  string
	JoinAt string
}

var fields = fieldsType{
	UUID:         "uuid",
	NickName:     "nick_name",
	Users:        "users",
	BusinessType: "business_type",
	IsVerified:   "is_verified",
	IsDeleted:    "is_deleted",
	IsEnabled:    "is_enabled",
}

var userFields = userFieldsType{
	UUID:   "uuid",
	Name:   "name",
	Roles:  "roles",
	JoinAt: "join_at",
}

func userFieldInArray(name string) string {
	return fields.Users + ".$." + name
}

func userField(name string) string {
	return fields.Users + "." + name
}

func userArrayFieldInArray(name string) string {
	return fields.Users + ".$[]." + name
}
