package errorland

func NewLand() *errorLand {
	return &errorLand{
		Errors: []string{},
	}
}

func NewLandFromError(err error) *errorLand {
	return &errorLand{
		Errors: []string{ err.Error() },
	}
}

func NewLandFromString(errMsg string) *errorLand {
	return &errorLand{
		Errors: []string{ errMsg },
	}
}