package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestWriteMapType(t *testing.T) {
	m := NewMap(standards[boolean], standards[boolean])
	expected := "map[bool]bool"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteArrayType(t *testing.T) {
	m := NewArray(standards[unknown], 0, true)
	expected := "[]unknown"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeEmpty(t *testing.T) {
	m := NewTuple()
	expected := "()"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeSingle(t *testing.T) {
	m := NewTuple(standards[boolean])
	expected := "(bool)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteTupleTypeMultiple(t *testing.T) {
	m := NewTuple(standards[boolean], standards[unknown])
	expected := "(bool, unknown)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsEmptyResults(t *testing.T) {
	m := NewFunc(NewTuple(), NewTuple())
	expected := "func()()"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsSingleResults(t *testing.T) {
	m := NewFunc(NewTuple(), NewTuple(standards[boolean]))
	expected := "func()(bool)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncMultipleParamsMultipleResults(t *testing.T) {
	m := NewFunc(NewTuple(standards[boolean], standards[unknown]), NewTuple(standards[boolean], standards[unknown]))
	expected := "func(bool, unknown)(bool, unknown)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteClass(t *testing.T) {

}
