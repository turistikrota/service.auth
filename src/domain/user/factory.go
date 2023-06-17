package user

import (
	"time"

	"github.com/mixarchitecture/i18np"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newUserErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) NewUser(email string, password []byte, token string) *User {
	t := time.Now()
	return &User{
		UUID:             "",
		Email:            email,
		Password:         password,
		Roles:            []string{"user"},
		VerifyToken:      token,
		IsVerified:       false,
		IsActive:         true,
		TwoFactorEnabled: false,
		Phone:            "",
		CreatedAt:        t,
		UpdatedAt:        t,
	}
}

func (f Factory) Unmarshal(uuid string, email string, isActive bool) *User {
	return &User{
		UUID:     uuid,
		Email:    email,
		IsActive: isActive,
	}
}

func (f Factory) Validate(u *User) *i18np.Error {
	if err := f.validateEmail(u.Email); err != nil {
		return err
	}
	return nil
}

func (f Factory) validateEmail(email string) *i18np.Error {
	if email == "" {
		return i18np.NewError(I18nMessages.EmailEmpty)
	}
	return nil
}
