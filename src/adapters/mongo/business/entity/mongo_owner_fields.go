package entity

type fields struct {
	UUID         string
	NickName     string
	BusinessType string
	Users        string
	IsEnabled    string
	IsVerified   string
	IsDeleted    string
}

type userFields struct {
	UUID   string
	Name   string
	Roles  string
	JoinAt string
}

var Fields = fields{
	UUID:         "uuid",
	NickName:     "nick_name",
	BusinessType: "business_type",
	Users:        "users",
	IsEnabled:    "is_enabled",
	IsVerified:   "is_verified",
	IsDeleted:    "is_deleted",
}

var UserFields = userFields{
	UUID:   "uuid",
	Name:   "name",
	Roles:  "roles",
	JoinAt: "join_at",
}

func UserField(field string) string {
	return Fields.Users + "." + field
}

func UserArrayFieldInArray(field string) string {
	return Fields.Users + ".$[]." + field
}
