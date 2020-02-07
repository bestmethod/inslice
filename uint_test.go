package inslice

import "testing"

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
