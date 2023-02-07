package errorland

import "fmt"

func StringMinMaxLength(in string, name string, minLength int, maxLength int) error {
	if len(in) < minLength {
		return fmt.Errorf(StrMinLenErrMsg, name, minLength)
	} else if len(in) > maxLength {
		return fmt.Errorf(StrMaxLenErrMsg, name, maxLength)
	}

	return nil
}