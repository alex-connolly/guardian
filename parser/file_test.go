package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestEmptyContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/empty.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestConstructorContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/constructors.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestFuncsContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/funcs.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestClassesContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/classes.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestInterfacesContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/interfaces.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, fmt.Sprintln(errs))
}

func TestEventsContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/events.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestEnumsContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/enums.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestTypesContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/types.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestNestedModifiersContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/nested_modifiers.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestCommentsContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/comments.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestGroupsContract(t *testing.T) {
	p, errs := ParseFile("../samples/tests/groups.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}
