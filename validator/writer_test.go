package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestWriteMapType(t *testing.T) {
	m := NewMap(standards[Bool], standards[Bool])
	expected := "map[bool]bool"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteArrayType(t *testing.T) {
	m := NewArray(standards[Unknown], 0, true)
	expected := "[]unknown"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeEmpty(t *testing.T) {
	m := NewTuple()
	expected := "()"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeSingle(t *testing.T) {
	m := NewTuple(standards[Bool])
	expected := "(bool)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeMultiple(t *testing.T) {
	m := NewTuple(standards[Bool], standards[Unknown])
	expected := "(bool, unknown)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsEmptyResults(t *testing.T) {
	m := NewFunc(NewTuple(), NewTuple())
	expected := "func()()"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsSingleResults(t *testing.T) {
	m := NewFunc(NewTuple(), NewTuple(standards[Bool]))
	expected := "func()(bool)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncMultipleParamsMultipleResults(t *testing.T) {
	m := NewFunc(NewTuple(standards[Bool], standards[Unknown]), NewTuple(standards[Bool], standards[Unknown]))
	expected := "func(bool, unknown)(bool, unknown)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteClass(t *testing.T) {

}
