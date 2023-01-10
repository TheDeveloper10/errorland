package errorland

type Land struct {
	Errors []string `json:"errors"`
}

func(el *Land) Merge(errs *Land) *Land {
	el.Errors = append(el.Errors, (errs.Errors)...)
	return el
}

func(el *Land) AppendString(errMsg string) *Land {
	el.Errors = append(el.Errors, errMsg)
	return el
}

func(el *Land) AppendError(err error) *Land {
	return el.AppendString(err.Error())
}

func(el *Land) RemoveIndex(index int) *Land {
	el.Errors = append(el.Errors[:index], el.Errors[index+1:]...)
	return el
}

func(el *Land) GetCount() int {
	return len(el.Errors)
}