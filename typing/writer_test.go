package typing

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestWriteMapType(t *testing.T) {
	m := &Map{Key: standards[boolean], Value: standards[boolean]}
	expected := "map[bool]bool"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteArrayType(t *testing.T) {
	m := &Array{Value: standards[unknown], Length: 0, Variable: true}
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
	m := &Func{Params: NewTuple(), Results: NewTuple()}
	expected := "func ()()"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsSingleResults(t *testing.T) {
	m := &Func{Params: NewTuple(), Results: NewTuple(standards[boolean])}
	expected := "func ()(bool)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteFuncMultipleParamsMultipleResults(t *testing.T) {
	m := &Func{Params: NewTuple(standards[boolean], standards[unknown]), Results: NewTuple(standards[boolean], standards[unknown])}
	expected := "func (bool, unknown)(bool, unknown)"
	goutil.Assert(t, WriteType(m) == expected, fmt.Sprintf("wrong type written: %s\n", WriteType(m)))
}

func TestWriteClass(t *testing.T) {

}
