package inslice

import "testing"

func TestIntFuncVar(t *testing.T) {
	slice := []*int{}
	item1 := int(1)
	slice = append(slice, &item1)
	item2 := int(2)
	slice = append(slice, &item2)
	item3 := int(3)
	slice = append(slice, &item3)
	item4 := int(2)
	slice = append(slice, &item4)
	ct := CompareInts(func(a int, b int) bool {
		if a == b {
			return true
		}
		return false
	})
	if len(IntVarFunc(slice, 2, ct, 3)) != 2 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(IntVarFunc(slice, 2, ct, 1)) != 1 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(IntVarFunc(slice, 5, ct, 3)) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(IntVarFuncAll(slice, 2, ct)) != 2 {
		t.Log("this should not have matched")
		t.FailNow()
	}
	if len(IntVarFuncAll(slice, 5, ct)) != 0 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestIntMatchVar(t *testing.T) {
	slice := []*int{}
	item1 := int(1)
	slice = append(slice, &item1)
	item2 := int(2)
	slice = append(slice, &item2)
	item3 := int(3)
	slice = append(slice, &item3)
	item4 := int(2)
	slice = append(slice, &item4)
	if IntVarMatch(slice, 1) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if IntVarMatch(slice, 5) != -1 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestHasIntVar(t *testing.T) {
	slice := []*int{}
	item1 := int(1)
	slice = append(slice, &item1)
	item2 := int(2)
	slice = append(slice, &item2)
	item3 := int(3)
	slice = append(slice, &item3)
	item4 := int(2)
	slice = append(slice, &item4)
	if HasIntVar(slice, 1) != true {
		t.Log("this should have matched")
		t.FailNow()
	}
	if HasIntVar(slice, 5) != false {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestIntFunc(t *testing.T) {
	slice := []int{
		1,
		2,
		3,
		2,
	}
	ct := CompareInts(func(a int, b int) bool {
		if a == b {
			return true
		}
		return false
	})
	if len(IntFunc(slice, 2, ct, 3)) != 2 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(IntFunc(slice, 2, ct, 1)) != 1 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(IntFunc(slice, 5, ct, 3)) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(IntFuncAll(slice, 2, ct)) != 2 {
		t.Log("this should not have matched")
		t.FailNow()
	}
	if len(IntFuncAll(slice, 5, ct)) != 0 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestIntMatch(t *testing.T) {
	slice := []int{
		1,
		2,
		3,
		2,
	}
	if IntMatch(slice, 1) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if IntMatch(slice, 5) != -1 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestHasInt(t *testing.T) {
	slice := []int{
		1,
		2,
		3,
		2,
	}
	if HasInt(slice, 1) != true {
		t.Log("this should have matched")
		t.FailNow()
	}
	if HasInt(slice, 5) != false {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

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
