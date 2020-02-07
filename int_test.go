package inslice

import "testing"

func TestInt(t *testing.T) {
	slice := []int{
		1,
		2,
		3,
		2,
	}
	if len(Int(slice, 4, 5)) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(Int(slice, 4, 0)) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(Int(slice, 2, 0)) != 0 {
		t.Log("Looking for text with count 0 must always return 0 results")
		t.FailNow()
	}

	if len(Int(slice, 2, 1)) != 1 {
		t.Log("Should have found 1 result")
		t.FailNow()
	}

	if len(Int(slice, 2, 5)) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestIntVar(t *testing.T) {
	slice := []*int{}
	item1 := 1
	slice = append(slice, &item1)
	item2 := 2
	slice = append(slice, &item2)
	item3 := 3
	slice = append(slice, &item3)
	item4 := 2
	slice = append(slice, &item4)
	if len(IntVar(slice, 4, 5)) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(IntVar(slice, 4, 0)) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(IntVar(slice, 2, 0)) != 0 {
		t.Log("Looking for text with count 0 must always return 0 results")
		t.FailNow()
	}

	if len(IntVar(slice, 2, 1)) != 1 {
		t.Log("Should have found 1 result")
		t.FailNow()
	}

	if len(IntVar(slice, 2, 5)) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestIntAll(t *testing.T) {
	slice := []int{
		1, 2, 3, 2,
	}
	if len(IntAll(slice, 4)) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(IntAll(slice, 4)) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(IntAll(slice, 2)) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestIntVarAll(t *testing.T) {
	slice := []*int{}
	item1 := 1
	slice = append(slice, &item1)
	item2 := 2
	slice = append(slice, &item2)
	item3 := 3
	slice = append(slice, &item3)
	item4 := 2
	slice = append(slice, &item4)
	if len(IntVarAll(slice, 4)) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(IntVarAll(slice, 4)) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(IntVarAll(slice, 2)) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}
