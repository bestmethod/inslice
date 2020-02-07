package inslice

import (
	"testing"
)

func TestReflectAll(t *testing.T) {
	slice := []string{
		"text1",
		"text2",
		"text3",
		"text2",
	}
	var i interface{}
	i = slice
	st := "text2"
	var s interface{}
	s = st
	if a, b := ReflectAll(i, s); len(a) != 2 || b != nil {
		t.Log("Must return 2")
		t.FailNow()
	}
}

func TestReflect(t *testing.T) {
	slice := []string{
		"text1",
		"text2",
		"text3",
		"text2",
	}
	var i interface{}
	i = slice
	st := "text2"
	var s interface{}
	s = st
	if a, b := Reflect(i, s, 1); len(a) != 1 || b != nil {
		t.Log("Must return 1")
		t.FailNow()
	}
	if a, b := Reflect(i, s, 0); len(a) != 0 || b != nil {
		t.Log("Must return 0")
		t.FailNow()
	}
	if a, b := Reflect(i, s, 5); len(a) != 2 || b != nil {
		t.Log("Must return 2")
		t.FailNow()
	}
	no := 5
	s = no
	if a, b := Reflect(i, s, 5); len(a) != 0 || b != nil {
		t.Log("Must return 0 and not panic")
		t.FailNow()
	}
	if _, b := Reflect(s, s, 1); b == nil {
		t.Log("slice is not slice, this should return an error")
		t.FailNow()
	}
}

func TestReflectPtr(t *testing.T) {
	slice := []*string{}
	item1 := "text1"
	slice = append(slice, &item1)
	item2 := "text2"
	slice = append(slice, &item2)
	item3 := "text3"
	slice = append(slice, &item3)
	item4 := "text2"
	slice = append(slice, &item4)
	var i interface{}
	i = slice
	st := "text2"
	var s interface{}
	s = &st
	if a, b := Reflect(i, s, 1); len(a) != 1 || b != nil {
		t.Log("Must return 1")
		t.FailNow()
	}
	if a, b := Reflect(i, s, 0); len(a) != 0 || b != nil {
		t.Log("Must return 0")
		t.FailNow()
	}
	if a, b := Reflect(i, s, 5); len(a) != 2 || b != nil {
		t.Log("Must return 2")
		t.FailNow()
	}
	no := 5
	s = no
	if a, b := Reflect(i, s, 5); len(a) != 0 || b != nil {
		t.Log("Must return 0 and not panic")
		t.FailNow()
	}
}
