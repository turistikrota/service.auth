package business

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newBusinessErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}
