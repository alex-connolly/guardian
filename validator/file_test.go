package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestSampleClass(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/class.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleContract(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/contracts.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleLoops(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/loops.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleSwitching(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/switching.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleTypes(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/types.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleIterator(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/iterator.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleKV(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/kv.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleAccessRestriction(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/common_patterns/access_restriction.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleRichest(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/common_patterns/richest.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleStateMachine(t *testing.T) {
	_, errs := ValidateFile(NewTestVM(), nil, "../samples/common_patterns/state_machine.grd")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestEmptyContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/empty.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestConstructorContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/constructors.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestFuncsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/funcs.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestClassesContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/classes.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestInterfacesContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/interfaces.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, fmt.Sprintln(errs))
}

func TestEventsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/events.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestEnumsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/enums.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestTypesContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/types.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestNestedModifiersContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/nested_modifiers.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestCommentsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/comments.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestGroupsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/groups.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}
