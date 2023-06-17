package dto

type Dto interface {
	UserUpdated() *UserUpdated
	AccountCreated() *AccountCreated
	AccountDisabled() *AccountDisabled
	AccountEnabled() *AccountEnabled
	AccountDeleted() *AccountDeleted
	AccountUpdated() *AccountUpdated
	OwnerUpdated() *OwnerUpdated
	OwnerCreated() *OwnerCreated
	OwnerUserAdded() *OwnerUserAdded
	OwnerUserRemoved() *OwnerUserRemoved
	OwnerPermissionEvent() *OwnerPermissionEvent
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

func (d *dto) OwnerUpdated() *OwnerUpdated {
	return &OwnerUpdated{}
}

func (d *dto) OwnerCreated() *OwnerCreated {
	return &OwnerCreated{}
}

func (d *dto) OwnerUserAdded() *OwnerUserAdded {
	return &OwnerUserAdded{}
}

func (d *dto) OwnerUserRemoved() *OwnerUserRemoved {
	return &OwnerUserRemoved{}
}

func (d *dto) OwnerPermissionEvent() *OwnerPermissionEvent {
	return &OwnerPermissionEvent{}
}
