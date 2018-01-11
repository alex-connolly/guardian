package typing

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIsAssignableEqualTypes(t *testing.T) {
	a := standards[boolean]
	b := standards[boolean]
	goutil.AssertNow(t, AssignableTo(b, a, false), "equal types should be Assignable")
}

func TestIsAssignableSuperClass(t *testing.T) {
	a := &Class{Name: "Dog"}
	b := &Class{Name: "Cat", Supers: []*Class{a}}
	goutil.AssertNow(t, AssignableTo(b, a, false), "super class types should be Assignable")
}

func TestIsAssignableMultipleSuperClass(t *testing.T) {
	a := &Class{Name: "Dog"}
	b := &Class{Name: "Cat", Supers: []*Class{a}}
	c := &Class{Name: "Rat", Supers: []*Class{b}}
	goutil.AssertNow(t, AssignableTo(c, a, false), "super class types should be Assignable")
}

func TestIsAssignableParentInterface(t *testing.T) {
	a := &Interface{Name: "Dog"}
	b := &Interface{Name: "Cat", Supers: []*Interface{a}}
	goutil.AssertNow(t, AssignableTo(b, a, false), "interface types should be Assignable")
}

func TestIsAssignableClassImplementingInterface(t *testing.T) {
	a := &Interface{Name: "Dog"}
	b := &Class{Name: "Cat", Interfaces: []*Interface{a}}
	goutil.AssertNow(t, AssignableTo(b, a, false), "interface types should be Assignable")
}

func TestIsAssignableSuperClassImplementingInterface(t *testing.T) {
	a := &Interface{Name: "Dog"}
	b := &Class{Name: "Cat", Interfaces: []*Interface{a}}
	c := &Class{Name: "Cat", Supers: []*Class{b}}
	goutil.AssertNow(t, AssignableTo(c, a, false), "interface types should be Assignable")
}

func TestIsAssignableSuperClassImplementingSuperInterface(t *testing.T) {
	a := &Interface{Name: "Dog"}
	b := &Interface{Name: "Lion", Supers: []*Interface{a}}
	c := &Class{Name: "Cat", Interfaces: []*Interface{b}}
	d := &Class{Name: "Tiger", Supers: []*Class{c}}
	goutil.AssertNow(t, AssignableTo(d, a, false), "type should be Assignable")
}

func TestIsAssignableClassDoesNotInherit(t *testing.T) {
	c := &Class{Name: "Cat"}
	d := &Class{Name: "Tiger"}
	goutil.AssertNow(t, !AssignableTo(d, c, false), "class should not be Assignable")
}

func TestIsAssignableClassFlipped(t *testing.T) {
	d := &Class{Name: "Tiger"}
	c := &Class{Name: "Cat", Supers: []*Class{d}}
	goutil.AssertNow(t, !AssignableTo(d, c, true), "class should not be Assignable")
}

func TestIsAssignableClassInterfaceNot(t *testing.T) {
	c := &Class{Name: "Cat"}
	d := &Interface{Name: "Tiger"}
	goutil.AssertNow(t, !AssignableTo(d, c, true), "class should not be Assignable")
	goutil.AssertNow(t, !AssignableTo(c, d, true), "interface should not be Assignable")
}

func TestIsAssignableBooleans(t *testing.T) {
	a := Boolean()
	b := Boolean()
	goutil.Assert(t, AssignableTo(a, b, false), "a --> b")
	goutil.Assert(t, AssignableTo(b, a, false), "b --> a")
}

func TestIsAssignableNamedBooleans(t *testing.T) {
	a := &Aliased{Alias: "a", Underlying: Boolean()}
	b := &Aliased{Alias: "b", Underlying: Boolean()}
	goutil.Assert(t, AssignableTo(a, b, false), "a --> b")
	goutil.Assert(t, AssignableTo(b, a, false), "b --> a")
}

func TestIsAssignableFuncParams(t *testing.T) {
	a := NewTuple(Boolean())
	b := NewTuple(&Aliased{Alias: "b", Underlying: Boolean()})
	goutil.Assert(t, AssignableTo(a, b, false), "a --> b")
	goutil.Assert(t, AssignableTo(b, a, false), "b --> a")
}

func TestNumericAssignability(t *testing.T) {
	a := &NumericType{BitSize: 10, Signed: true}
	b := &NumericType{BitSize: 10, Signed: true}
	goutil.Assert(t, AssignableTo(a, b, false), "a --> b")
	// ints --> larger ints
	a = &NumericType{BitSize: 10, Signed: true}
	b = &NumericType{BitSize: 11, Signed: true}
	goutil.Assert(t, AssignableTo(b, a, false), "b --> a")
	goutil.Assert(t, !AssignableTo(a, b, false), "a --> b")
	// uints --> larger uints
	a = &NumericType{BitSize: 10, Signed: false}
	b = &NumericType{BitSize: 11, Signed: false}
	goutil.Assert(t, AssignableTo(b, a, false), "b --> a")
	goutil.Assert(t, !AssignableTo(a, b, false), "a --> b")
	// uints --> larger ints
	a = &NumericType{BitSize: 10, Signed: false}
	b = &NumericType{BitSize: 11, Signed: true}
	goutil.Assert(t, AssignableTo(b, a, false), "b --> a")
	goutil.Assert(t, !AssignableTo(a, b, false), "a --> b")
	// can never go int --> uint
	a = &NumericType{BitSize: 10, Signed: false}
	b = &NumericType{BitSize: 100, Signed: true}
	goutil.Assert(t, !AssignableTo(a, b, false), "a --> b")
}

func TestIsAssignableUnknown(t *testing.T) {
	a := Unknown()
	b := Boolean()
	goutil.Assert(t, !AssignableTo(a, b, false), "a --> b")
	goutil.Assert(t, AssignableTo(a, b, true), "a --> b")
}

func TestIsAssignableEmptyTuples(t *testing.T) {
	a := NewTuple(Unknown())
	b := NewTuple()
	goutil.Assert(t, !AssignableTo(a, b, true), "a --> b")
}
