package req

type Request interface {
	Login() *LoginRequest
	LoginVerified() *LoginVerifiedRequest
	Register() *RegisterRequest
	Logout() *LogoutRequest
	RefreshToken() *RefreshTokenRequest
	Verify() *VerifyRequest
	ReSendVerification() *ReSendVerificationRequest
	CheckEmail() *CheckEmailRequest
}

type request struct{}

func New() Request {
	return &request{}
}
