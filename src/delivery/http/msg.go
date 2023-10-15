package http

type successMessages struct {
	Login                string
	Logout               string
	Register             string
	Extend               string
	Verify               string
	ReSendVerification   string
	CurrentUser          string
	EmailAvailable       string
	UserList             string
	SessionDestroy       string
	SessionDestroyOthers string
	SessionDestroyAll    string
	SessionList          string
	FcmSet               string
	ChangePassword       string
}

type errorMessages struct {
	Login                 string
	Logout                string
	Register              string
	Extend                string
	Unexpected            string
	LoginVerify           string
	TurnstileBadRequest   string
	TurnstileUnauthorized string
	AdminRoute            string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		Login:                "success_login",
		Logout:               "success_logout",
		Register:             "success_register",
		Extend:               "success_extend",
		Verify:               "success_verify",
		ReSendVerification:   "success_resend_verification",
		CurrentUser:          "success_current_user",
		EmailAvailable:       "success_email_available",
		UserList:             "success_user_list",
		SessionDestroy:       "success_session_destroy",
		SessionDestroyOthers: "success_session_destroy_others",
		SessionDestroyAll:    "success_session_destroy_all",
		SessionList:          "success_session_list",
		FcmSet:               "success_fcm_set",
		ChangePassword:       "success_change_password",
	},
	Error: errorMessages{
		Login:                 "error_login",
		Logout:                "error_logout",
		Register:              "error_register",
		Extend:                "error_extend",
		Unexpected:            "error_unexpected",
		LoginVerify:           "error_login_verify",
		TurnstileBadRequest:   "error_turnstile_bad_request",
		TurnstileUnauthorized: "error_turnstile_unauthorized",
		AdminRoute:            "error_admin_route",
	},
}
