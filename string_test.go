package inslice

import "testing"

func TestString(t *testing.T) {
	slice := []string{
		"text1",
		"text2",
		"text3",
		"text2",
	}
	if len(String(slice, "text", 5)) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(String(slice, "text", 0)) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(String(slice, "text2", 0)) != 0 {
		t.Log("Looking for text with count 0 must always return 0 results")
		t.FailNow()
	}

	if len(String(slice, "text2", 1)) != 1 {
		t.Log("Should have found 1 result")
		t.FailNow()
	}

	if len(String(slice, "text2", 5)) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestStringVar(t *testing.T) {
	slice := []*string{}
	item1 := "text1"
	slice = append(slice, &item1)
	item2 := "text2"
	slice = append(slice, &item2)
	item3 := "text3"
	slice = append(slice, &item3)
	item4 := "text2"
	slice = append(slice, &item4)
	if len(StringVar(slice, "text", 5)) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(StringVar(slice, "text", 0)) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(StringVar(slice, "text2", 0)) != 0 {
		t.Log("Looking for text with count 0 must always return 0 results")
		t.FailNow()
	}

	if len(StringVar(slice, "text2", 1)) != 1 {
		t.Log("Should have found 1 result")
		t.FailNow()
	}

	if len(StringVar(slice, "text2", 5)) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestStringAll(t *testing.T) {
	slice := []string{
		"text1",
		"text2",
		"text3",
		"text2",
	}
	if len(StringAll(slice, "text")) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(StringAll(slice, "text")) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(StringAll(slice, "text2")) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}

func TestStringVarAll(t *testing.T) {
	slice := []*string{}
	item1 := "text1"
	slice = append(slice, &item1)
	item2 := "text2"
	slice = append(slice, &item2)
	item3 := "text3"
	slice = append(slice, &item3)
	item4 := "text2"
	slice = append(slice, &item4)
	if len(StringVarAll(slice, "text")) != 0 {
		t.Log("Looking for similar string should still not match")
		t.FailNow()
	}

	if len(StringVarAll(slice, "text")) != 0 {
		t.Log("Looking for similar string should still not match, len of 0 should not matter")
		t.FailNow()
	}

	if len(StringVarAll(slice, "text2")) != 2 {
		t.Log("Should have found 2 results")
		t.FailNow()
	}
}
