package inslice

import (
	"strings"
	"testing"
)

func TestStringFuncVar(t *testing.T) {
	slice := []*string{}
	item1 := "text1"
	slice = append(slice, &item1)
	item2 := "text2"
	slice = append(slice, &item2)
	item3 := "text3"
	slice = append(slice, &item3)
	item4 := "text2"
	slice = append(slice, &item4)
	ct := CompareStrings(strings.Contains)
	if len(StringVarFunc(slice, "xt", ct, 3)) != 3 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(StringVarFunc(slice, "st", ct, 3)) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(StringVarFuncAll(slice, "2", ct)) != 2 {
		t.Log("this should not have matched")
		t.FailNow()
	}
	if len(StringVarFuncAll(slice, "test1", ct)) != 0 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestStringMatchVar(t *testing.T) {
	slice := []*string{}
	item1 := "text1"
	slice = append(slice, &item1)
	item2 := "text2"
	slice = append(slice, &item2)
	item3 := "text3"
	slice = append(slice, &item3)
	item4 := "text2"
	slice = append(slice, &item4)
	if StringVarMatch(slice, "text1") != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if StringVarMatch(slice, "test1") != -1 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestHasStringVar(t *testing.T) {
	slice := []*string{}
	item1 := "text1"
	slice = append(slice, &item1)
	item2 := "text2"
	slice = append(slice, &item2)
	item3 := "text3"
	slice = append(slice, &item3)
	item4 := "text2"
	slice = append(slice, &item4)
	if HasStringVar(slice, "text1") != true {
		t.Log("this should have matched")
		t.FailNow()
	}
	if HasStringVar(slice, "test1") != false {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestStringFunc(t *testing.T) {
	slice := []string{
		"text1",
		"text2",
		"text3",
		"text2",
	}
	ct := CompareStrings(strings.Contains)
	if len(StringFunc(slice, "xt", ct, 3)) != 3 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(StringFunc(slice, "st", ct, 3)) != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if len(StringFuncAll(slice, "2", ct)) != 2 {
		t.Log("this should not have matched")
		t.FailNow()
	}
	if len(StringFuncAll(slice, "test1", ct)) != 0 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestStringMatch(t *testing.T) {
	slice := []string{
		"text1",
		"text2",
		"text3",
		"text2",
	}
	if StringMatch(slice, "text1") != 0 {
		t.Log("this should have matched")
		t.FailNow()
	}
	if StringMatch(slice, "test1") != -1 {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

func TestHasString(t *testing.T) {
	slice := []string{
		"text1",
		"text2",
		"text3",
		"text2",
	}
	if HasString(slice, "text1") != true {
		t.Log("this should have matched")
		t.FailNow()
	}
	if HasString(slice, "test1") != false {
		t.Log("this should not have matched")
		t.FailNow()
	}
}

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
