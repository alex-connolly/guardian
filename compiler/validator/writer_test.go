package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestWriteMapType(t *testing.T) {
	m := NewMap(standards[Int], standards[Int])
	expected := "map[int]int"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteArrayType(t *testing.T) {
	m := NewArray(standards[String])
	expected := "[string]"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeEmpty(t *testing.T) {
	m := NewTuple()
	expected := "()"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeSingle(t *testing.T) {
	m := NewTuple(standards[Int])
	expected := "(int)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeMultiple(t *testing.T) {
	m := NewTuple(standards[Int], standards[String])
	expected := "(int, string)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsEmptyResults(t *testing.T) {
	m := NewFunc(NewTuple(), NewTuple())
	expected := "func()()"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsSingleResults(t *testing.T) {
	m := NewFunc(NewTuple(), NewTuple(standards[Int]))
	expected := "func()(int)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncMultipleParamsMultipleResults(t *testing.T) {
	m := NewFunc(NewTuple(standards[Int], standards[String]), NewTuple(standards[Int], standards[String]))
	expected := "func(int, string)(int, string)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}
