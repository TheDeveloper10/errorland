package errorland

import (
	"errors"
	"testing"
)

func Test_1_Land(t *testing.T) {
	errs := NewLand()

	if errs.GetCount() != 0 {
		t.Error("Empty land not initialized correctly")
	}
}

func Test_2_Land(t *testing.T) {
	errs := NewLandFromString("aaa bbb ccc")

	if errs.GetCount() != 1 {
		t.Error("Land not initialized correctly")
	} else if errs.Errors[0] != "aaa bbb ccc" {
		t.Error("Land's content is wrong")
	}
}

func Test_3_Land(t *testing.T) {
	errs := NewLandFromError(errors.New("aaa ccc bbb"))

	if errs.GetCount() != 1 {
		t.Error("Land not initialized correctly")
	} else if errs.Errors[0] != "aaa ccc bbb" {
		t.Error("Land's content is wrong")
	}
}

func Test_4_Land(t *testing.T) {
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

func Test_5_Land(t *testing.T) {
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

func Test_6_Land(t *testing.T) {
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

func Test_7_Land(t *testing.T) {
	errs := NewLand()

	errs.AppendString("1")
	errs.AppendError(errors.New("2"))
	errs.AppendString("3")
	errs.AppendError(errors.New("4"))
	errs.AppendString("5")
	errs.AppendError(errors.New("6"))

	errs.RemoveIndex(1)
	errs.RemoveIndex(2)
	errs.RemoveIndex(2)

	if errs.GetCount() != 3 {
		t.Error("Remove failed")
	} else if errs.Errors[0] != "1" || errs.Errors[1] != "3" || errs.Errors[2] != "6" {
		t.Error("Remove failed: wrong items")
	}
}