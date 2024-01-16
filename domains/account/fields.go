package account

type fieldsType struct {
	UserUUID   string
	UserName   string
	IsActive   string
	IsDeleted  string
	IsVerified string
	BirthDate  string
}

var fields = fieldsType{
	UserUUID:   "user_uuid",
	UserName:   "user_name",
	IsActive:   "is_active",
	IsDeleted:  "is_deleted",
	IsVerified: "is_verified",
	BirthDate:  "birth_date",
}
