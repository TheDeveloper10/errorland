package errorland

import "fmt"

func IntMinMax(in int, name string, min int, max int) error {
	if in < min {
		return fmt.Errorf(IntMinErrMsg, name, min)
	} else if in > max {
		return fmt.Errorf(IntMaxErrMsg, name, max)
	}

	return nil
}