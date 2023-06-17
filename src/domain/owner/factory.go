package owner

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newOwnerErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}
