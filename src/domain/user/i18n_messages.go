package user

type messages struct {
	EmailEmpty                string
	NotFound                  string
	AlreadyExists             string
	Failed                    string
	InvalidPassword           string
	Deleted                   string
	RefreshTokenNotAvailable  string
	TwoFactorStarted          string
	TokenExpired              string
	AlreadyVerified           string
	NotVerified               string
	TokenNotExpired           string
	AuthRegisteredMailSubject string
}

var I18nMessages = messages{
	EmailEmpty:                "error_user_email_empty",
	NotFound:                  "error_user_not_found",
	Deleted:                   "error_user_deleted",
	Failed:                    "error_user_failed",
	AlreadyExists:             "error_user_already_exists",
	InvalidPassword:           "error_user_invalid_password",
	RefreshTokenNotAvailable:  "error_user_refresh_token_not_available",
	TwoFactorStarted:          "error_user_two_factor_started",
	TokenExpired:              "error_user_token_expired",
	AlreadyVerified:           "error_user_already_verified",
	NotVerified:               "error_user_not_verified",
	TokenNotExpired:           "error_user_token_not_expired",
	AuthRegisteredMailSubject: "auth_registered_mail_subject",
}
