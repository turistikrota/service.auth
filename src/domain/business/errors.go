package business

import "github.com/mixarchitecture/i18np"

type Errors interface {
	NotFound() *i18np.Error
	Failed(action string) *i18np.Error
}

type errors struct{}

func newBusinessErrors() Errors {
	return &errors{}
}

func (e *errors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.BusinessNotFound)
}

func (e *errors) Failed(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.BusinessFailed, i18np.P{
		"Action": action,
	})
}
