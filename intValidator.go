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

func UintMin(in uint, name string, min uint) error {
	if in < min {
		return fmt.Errorf(IntMinErrMsg, name, min)
	}

	return nil
}

func UintMax(in uint, name string, min uint) error {
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

func UintMinMax(in uint, name string, min uint, max uint) error {
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

func StrToUint64(in string, name string) (uint64, error) {
	res, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(StrToUint64ErrMsg, name)
	}

	return res, nil
}
