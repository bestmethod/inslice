package inslice

import (
	"testing"
)

func TestUintFuncVar(t *testing.T) {
	slice := []*uint{}
	item1 := uint(1)
	slice = append(slice, &item1)
	item2 := uint(2)
	slice = append(slice, &item2)
	item3 := uint(3)
	slice = append(slice, &item3)
	item4 := uint(2)
	slice = append(slice, &item4)
	ct := CompareUints(func(a uint, b uint) bool {
		if a == b {
			return true
		}
		return false
	})
	if len(UintVarFunc(slice, 2, ct, 3)) != 2 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(UintVarFunc(slice, 2, ct, 1)) != 1 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(UintVarFunc(slice, 5, ct, 3)) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(UintVarFuncAll(slice, 2, ct)) != 2 {
		t.Log("this should not have matched")
		t.FailNow()
	}
	if len(UintVarFuncAll(slice, 5, ct)) != 0 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestUintMatchVar(t *testing.T) {
	slice := []*uint{}
	item1 := uint(1)
	slice = append(slice, &item1)
	item2 := uint(2)
	slice = append(slice, &item2)
	item3 := uint(3)
	slice = append(slice, &item3)
	item4 := uint(2)
	slice = append(slice, &item4)
	if UintVarMatch(slice, 1) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if UintVarMatch(slice, 5) != -1 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestHasUintVar(t *testing.T) {
	slice := []*uint{}
	item1 := uint(1)
	slice = append(slice, &item1)
	item2 := uint(2)
	slice = append(slice, &item2)
	item3 := uint(3)
	slice = append(slice, &item3)
	item4 := uint(2)
	slice = append(slice, &item4)
	if HasUintVar(slice, 1) != true {
		t.Log("this should have matched")
		t.FailNow()
	}
	if HasUintVar(slice, 5) != false {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestUintFunc(t *testing.T) {
	slice := []uint{
		1,
		2,
		3,
		2,
	}
	ct := CompareUints(func(a uint, b uint) bool {
		if a == b {
			return true
		}
		return false
	})
	if len(UintFunc(slice, 2, ct, 3)) != 2 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(UintFunc(slice, 2, ct, 1)) != 1 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(UintFunc(slice, 5, ct, 3)) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(UintFuncAll(slice, 2, ct)) != 2 {
		t.Log("this should not have matched")
		t.FailNow()
	}
	if len(UintFuncAll(slice, 5, ct)) != 0 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestUintMatch(t *testing.T) {
	slice := []uint{
		1,
		2,
		3,
		2,
	}
	if UintMatch(slice, 1) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if UintMatch(slice, 5) != -1 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestHasUint(t *testing.T) {
	slice := []uint{
		1,
		2,
		3,
		2,
	}
	if HasUint(slice, 1) != true {
		t.Log("this should have matched")
		t.FailNow()
	}
	if HasUint(slice, 5) != false {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestUint(t *testing.T) {
	slice := []uint{
		1,
		2,
		3,
		2,
	}
	if len(Uint(slice, uint(4), 5)) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(Uint(slice, uint(4), 0)) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(Uint(slice, uint(2), 0)) != 0 {
		t.Log("Looking for text with count 0 must always return 0 results")
		t.FailNow()
	}

	if len(Uint(slice, uint(2), 1)) != 1 {
		t.Log("Should have found 1 result")
		t.FailNow()
	}

	if len(Uint(slice, uint(2), 5)) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestUintVar(t *testing.T) {
	slice := []*uint{}
	item1 := uint(1)
	slice = append(slice, &item1)
	item2 := uint(2)
	slice = append(slice, &item2)
	item3 := uint(3)
	slice = append(slice, &item3)
	item4 := uint(2)
	slice = append(slice, &item4)
	if len(UintVar(slice, uint(4), 5)) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(UintVar(slice, uint(4), 0)) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(UintVar(slice, uint(2), 0)) != 0 {
		t.Log("Looking for text with count 0 must always return 0 results")
		t.FailNow()
	}

	if len(UintVar(slice, uint(2), 1)) != 1 {
		t.Log("Should have found 1 result")
		t.FailNow()
	}

	if len(UintVar(slice, uint(2), 5)) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestUintAll(t *testing.T) {
	slice := []uint{
		1, 2, 3, 2,
	}
	if len(UintAll(slice, uint(4))) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(UintAll(slice, uint(4))) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(UintAll(slice, uint(2))) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestUintVarAll(t *testing.T) {
	slice := []*uint{}
	item1 := uint(1)
	slice = append(slice, &item1)
	item2 := uint(2)
	slice = append(slice, &item2)
	item3 := uint(3)
	slice = append(slice, &item3)
	item4 := uint(2)
	slice = append(slice, &item4)
	if len(UintVarAll(slice, uint(4))) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(UintVarAll(slice, uint(4))) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(UintVarAll(slice, uint(2))) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}
