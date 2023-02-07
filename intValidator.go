package errorland

import (
	"fmt"
	"strconv"
)

func IntMin(in int, name string, min int) error {
	if in < min {
		return fmt.Errorf(IntMinErrMsg, name, min)
	}

	return nil
}

func IntMax(in int, name string, min int) error {
	if in > min {
		return fmt.Errorf(IntMaxErrMsg, name, min)
	}

	return nil
}

func IntMinMax(in int, name string, min int, max int) error {
	if in < min {
		return fmt.Errorf(IntMinErrMsg, name, min)
	} else if in > max {
		return fmt.Errorf(IntMaxErrMsg, name, max)
	}

	return nil
}

func StrToInt64(in string, name string) (int64, error) {
	res, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(StrToInt64ErrMsg, name)
	}

	return res, nil
}