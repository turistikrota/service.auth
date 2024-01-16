package user

import "github.com/cilloparch/cillop/i18np"

type TwoFactor struct {
	TempToken string `json:"tempToken"`
	Verify    bool   `json:"verify"`
}

type Errors interface {
	NotFound(email string) *i18np.Error
	AlreadyExists(email string) *i18np.Error
	Failed(operation string) *i18np.Error
	InvalidPassword() *i18np.Error
	InvalidUUID() *i18np.Error
	RefreshTokenNotAvailable() *i18np.Error
	TwoFactorStarted(token string) *i18np.Error
	TokenExpired(p interface{}) *i18np.Error
	AlreadyVerified() *i18np.Error
	NotVerified(p interface{}) *i18np.Error
	Deleted() *i18np.Error
	TokenNotExpired() *i18np.Error
}

type userErrors struct{}

func newUserErrors() Errors {
	return &userErrors{}
}

func (e *userErrors) NotFound(email string) *i18np.Error {
	return i18np.NewError(I18nMessages.NotFound, i18np.P{"Email": email})
}

func (e *userErrors) AlreadyExists(email string) *i18np.Error {
	return i18np.NewError(I18nMessages.AlreadyExists, i18np.P{"Email": email})
}

func (e *userErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *userErrors) InvalidPassword() *i18np.Error {
	return i18np.NewError(I18nMessages.InvalidPassword)
}

func (e *userErrors) InvalidUUID() *i18np.Error {
	return i18np.NewError(I18nMessages.InvalidUUID)
}

func (e *userErrors) RefreshTokenNotAvailable() *i18np.Error {
	return i18np.NewError(I18nMessages.RefreshTokenNotAvailable)
}

func (e *userErrors) TwoFactorStarted(token string) *i18np.Error {
	return i18np.NewErrorDetails(I18nMessages.TwoFactorStarted, &TwoFactor{
		TempToken: token,
		Verify:    true,
	})
}

func (e *userErrors) TokenExpired(p interface{}) *i18np.Error {
	return i18np.NewErrorDetails(I18nMessages.TokenExpired, p)
}

func (e *userErrors) AlreadyVerified() *i18np.Error {
	return i18np.NewError(I18nMessages.AlreadyVerified)
}

func (e *userErrors) NotVerified(p interface{}) *i18np.Error {
	return i18np.NewErrorDetails(I18nMessages.NotVerified, p)
}

func (e *userErrors) Deleted() *i18np.Error {
	return i18np.NewError(I18nMessages.Deleted)
}

func (e *userErrors) TokenNotExpired() *i18np.Error {
	return i18np.NewError(I18nMessages.TokenNotExpired)
}
