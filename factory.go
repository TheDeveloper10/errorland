package errorland

func NewLand() *Land {
	return &Land{
		Errors: []string{},
	}
}

func NewLandFromError(err error) *Land {
	return &Land{
		Errors: []string{ err.Error() },
	}
}

func NewLandFromString(errMsg string) *Land {
	return &Land{
		Errors: []string{ errMsg },
	}
}