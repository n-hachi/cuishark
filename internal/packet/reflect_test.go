package packet

import (
	"reflect"
	"testing"
)

type SimpleStruct struct {
	I int
	U uint
}

func TestOneline_SimpleStruct(t *testing.T) {
	ss := SimpleStruct{I: 1, U: 2}
	cmp := "{I:1, U:2}"

	if ol := Oneline(ss); ol != cmp {
		t.Errorf("Oneline(ss) should be %s, but actually %s", cmp, ol)
	}
}

type NestedStruct struct {
	SS SimpleStruct
	I  int
}

func TestOneline_NestesStruct(t *testing.T) {
	ss := SimpleStruct{I: 1, U: 2}
	ns := NestedStruct{SS: ss, I: 3}
	cmp := "{SS:{I:1, U:2}, I:3}"

	if ol := Oneline(ns); ol != cmp {
		t.Errorf("Oneline(ns) should be %s, but actually %s", cmp, ol)
	}
}

type ArrayStruct struct {
	Ary []int
}

func TestOneline_ArrayStruct(t *testing.T) {
	as := ArrayStruct{Ary: []int{1, 2, 3}}
	cmp := "{Ary:[1 2 3]}"
	if ol := Oneline(as); ol != cmp {
		t.Errorf("Oneline(as) should be %s, but actually %s", cmp, ol)
	}
}

func TestDetail_SimpleStruct(t *testing.T) {
	ss := SimpleStruct{I: 1, U: 2}
	cmp := []string{"I:1", "U:2"}
	if ol := Detail(ss); reflect.DeepEqual(ol, cmp) {
		t.Errorf("Oneline(as) should be %v, but actually %v", cmp, ol)
	}
}
