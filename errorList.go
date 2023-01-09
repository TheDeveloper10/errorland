package errorland

type ErrorList struct {
	errors []string `json:"errors"`
}

func(el *ErrorList) Merge(errs *ErrorList) {
	el.errors = append(el.errors, (errs.errors)...)
}