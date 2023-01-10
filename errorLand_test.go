package errorland

import (
	"errors"
	"testing"
)

func Test_1_ErrorLand(t *testing.T) {
	errs := NewLand()

	if errs.GetCount() != 0 {
		t.Error("Empty land not initialized correctly")
	}
}

func Test_2_ErrorLand(t *testing.T) {
	errs := NewLandFromString("aaa bbb ccc")

	if errs.GetCount() != 1 {
		t.Error("Land not initialized correctly")
	} else if errs.Errors[0] != "aaa bbb ccc" {
		t.Error("Land's content is wrong")
	}
}

func Test_3_ErrorLand(t *testing.T) {
	errs := NewLandFromError(errors.New("aaa ccc bbb"))

	if errs.GetCount() != 1 {
		t.Error("Land not initialized correctly")
	} else if errs.Errors[0] != "aaa ccc bbb" {
		t.Error("Land's content is wrong")
	}
}

func Test_4_ErrorLand(t *testing.T) {
	errs1 := NewLandFromError(errors.New("a"))
	errs2 := NewLandFromString("b")

	errs1.Merge(errs2)

	if errs1.GetCount() != 2 {
		t.Error("Merge failed")
	} else if errs2.GetCount() != 1 {
		t.Error("Land 2 has contents it shouldn't have")
	} else if errs1.Errors[0] != "a" || errs1.Errors[1] != "b" {
		t.Error("Content after merge has been messed up")
	}
}

func Test_5_ErrorLand(t *testing.T) {
	errs := NewLand()

	errs.AppendError(errors.New("a"))
	errs.AppendError(errors.New("b"))
	errs.AppendError(errors.New("c"))

	if errs.GetCount() != 3 {
		t.Error("Append failed")
	} else if errs.Errors[0] != "a" || errs.Errors[1] != "b" || errs.Errors[2] != "c" {
		t.Error("Content after append is wrong")
	}
}

func Test_6_ErrorLand(t *testing.T) {
	errs := NewLand()

	errs.AppendString("a")
	errs.AppendString("b")
	errs.AppendString("c")

	if errs.GetCount() != 3 {
		t.Error("Append failed")
	} else if errs.Errors[0] != "a" || errs.Errors[1] != "b" || errs.Errors[2] != "c" {
		t.Error("Content after append is wrong")
	}
}