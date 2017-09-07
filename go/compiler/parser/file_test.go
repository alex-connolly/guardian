package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestEmptyContract(t *testing.T) {
	p := ParseFile("tests/empty.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}

func TestConstructorContract(t *testing.T) {
	p := ParseFile("tests/constructors.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}

func TestVariablesContract(t *testing.T) {
	p := ParseFile("tests/variables.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}

func TestFuncsContract(t *testing.T) {
	p := ParseFile("tests/funcs.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}

func TestClassesContract(t *testing.T) {
	p := ParseFile("tests/classes.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}

func TestInterfacesContract(t *testing.T) {
	p := ParseFile("tests/interfaces.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}

func TestEventsContract(t *testing.T) {
	p := ParseFile("tests/events.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}

func TestEnumsContract(t *testing.T) {
	p := ParseFile("tests/enums.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}

func TestTypesContract(t *testing.T) {
	p := ParseFile("tests/types.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
}
