package errorland

import (
	"fmt"
)

func Float32Min(in float32, name string, min float32) error {
	if in < min {
		return fmt.Errorf(IntMinErrMsg, name, min)
	}

	return nil
}

func Float32Max(in float32, name string, min float32) error {
	if in > min {
		return fmt.Errorf(IntMaxErrMsg, name, min)
	}

	return nil
}

func Float32MinMax(in float32, name string, min float32, max float32) error {
	if in < min {
		return fmt.Errorf(IntMinErrMsg, name, min)
	} else if in > max {
		return fmt.Errorf(IntMaxErrMsg, name, max)
	}

	return nil
}