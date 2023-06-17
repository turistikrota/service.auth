package owner

import "github.com/mixarchitecture/i18np"

type Errors interface {
	NotFound() *i18np.Error
	Failed(action string) *i18np.Error
}

type errors struct{}

func newOwnerErrors() Errors {
	return &errors{}
}

func (e *errors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.OwnerNotFound)
}

func (e *errors) Failed(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.OwnerFailed, i18np.P{
		"Action": action,
	})
}
