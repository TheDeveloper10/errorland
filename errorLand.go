package errorland

type errorLand struct {
	Errors []string `json:"errors"`
}

func(el *errorLand) Merge(errs *errorLand) *errorLand {
	el.Errors = append(el.Errors, (errs.Errors)...)
	return el
}

func(el *errorLand) AppendString(errMsg string) *errorLand {
	el.Errors = append(el.Errors, errMsg)
	return el
}

func(el *errorLand) AppendError(err error) *errorLand {
	return el.AppendString(err.Error())
}

func(el *errorLand) RemoveIndex(index int) *errorLand {
	el.Errors = append(el.Errors[:index], el.Errors[index+1:]...)
	return el
}

func(el *errorLand) GetCount() int {
	return len(el.Errors)
}