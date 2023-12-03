package dto

type Dto interface {
	UserUpdated() *UserUpdated
	AccountCreated() *AccountCreated
	AccountDisabled() *AccountDisabled
	AccountEnabled() *AccountEnabled
	AccountDeleted() *AccountDeleted
	AccountUpdated() *AccountUpdated
	BusinessUpdated() *BusinessUpdated
	BusinessCreated() *BusinessCreated
	BusinessUserAdded() *BusinessUserAdded
	BusinessUserRemoved() *BusinessUserRemoved
	BusinessPermissionEvent() *BusinessPermissionEvent
	UserRolesEvent() *UserRolesEvent
}

type dto struct{}

func New() Dto {
	return &dto{}
}

func (d *dto) UserUpdated() *UserUpdated {
	return &UserUpdated{}
}

func (d *dto) AccountCreated() *AccountCreated {
	return &AccountCreated{}
}

func (d *dto) AccountDisabled() *AccountDisabled {
	return &AccountDisabled{}
}

func (d *dto) AccountEnabled() *AccountEnabled {
	return &AccountEnabled{}
}

func (d *dto) AccountDeleted() *AccountDeleted {
	return &AccountDeleted{}
}

func (d *dto) AccountUpdated() *AccountUpdated {
	return &AccountUpdated{}
}

func (d *dto) BusinessUpdated() *BusinessUpdated {
	return &BusinessUpdated{}
}

func (d *dto) BusinessCreated() *BusinessCreated {
	return &BusinessCreated{}
}

func (d *dto) BusinessUserAdded() *BusinessUserAdded {
	return &BusinessUserAdded{}
}

func (d *dto) BusinessUserRemoved() *BusinessUserRemoved {
	return &BusinessUserRemoved{}
}

func (d *dto) BusinessPermissionEvent() *BusinessPermissionEvent {
	return &BusinessPermissionEvent{}
}

func (d *dto) UserRolesEvent() *UserRolesEvent {
	return &UserRolesEvent{}
}
