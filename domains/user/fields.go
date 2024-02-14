package user

type fieldsType struct {
	UUID             string
	Email            string
	Phone            string
	Roles            string
	Password         string
	TwoFactorEnabled string
	IsActive         string
	IsDeleted        string
	IsVerified       string
	VerifyToken      string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}

var fields = fieldsType{
	UUID:             "_id",
	Email:            "email",
	Phone:            "phone",
	Roles:            "roles",
	Password:         "password",
	TwoFactorEnabled: "two_factor_enabled",
	IsActive:         "is_active",
	IsDeleted:        "is_deleted",
	IsVerified:       "is_verified",
	VerifyToken:      "email_verify_token",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}
