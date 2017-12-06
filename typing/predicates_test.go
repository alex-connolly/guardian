package typing

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIsAssignableEqualTypes(t *testing.T) {
	a := standards[boolean]
	b := standards[boolean]
	goutil.AssertNow(t, assignableTo(b, a), "equal types should be assignable")
}

func TestIsAssignableSuperClass(t *testing.T) {
	a := Class{Name: "Dog"}
	b := Class{Name: "Cat", Supers: []*Class{&a}}
	goutil.AssertNow(t, assignableTo(b, a), "super class types should be assignable")
}

func TestIsAssignableMultipleSuperClass(t *testing.T) {
	a := Class{Name: "Dog"}
	b := Class{Name: "Cat", Supers: []*Class{&a}}
	c := Class{Name: "Rat", Supers: []*Class{&b}}
	goutil.AssertNow(t, assignableTo(c, a), "super class types should be assignable")
}

func TestIsAssignableParentInterface(t *testing.T) {
	a := Interface{Name: "Dog"}
	b := Interface{Name: "Cat", Supers: []*Interface{&a}}
	goutil.AssertNow(t, assignableTo(b, a), "interface types should be assignable")
}

func TestIsAssignableClassImplementingInterface(t *testing.T) {
	a := Interface{Name: "Dog"}
	b := Class{Name: "Cat", Interfaces: []*Interface{&a}}
	goutil.AssertNow(t, assignableTo(b, a), "interface types should be assignable")
}

func TestIsAssignableSuperClassImplementingInterface(t *testing.T) {
	a := Interface{Name: "Dog"}
	b := Class{Name: "Cat", Interfaces: []*Interface{&a}}
	c := Class{Name: "Cat", Supers: []*Class{&b}}
	goutil.AssertNow(t, assignableTo(c, a), "interface types should be assignable")
}

func TestIsAssignableSuperClassImplementingSuperInterface(t *testing.T) {
	a := Interface{Name: "Dog"}
	b := Interface{Name: "Lion", Supers: []*Interface{&a}}
	c := Class{Name: "Cat", Interfaces: []*Interface{&b}}
	d := Class{Name: "Tiger", Supers: []*Class{&c}}
	goutil.AssertNow(t, assignableTo(d, a), "type should be assignable")
}

func TestIsAssignableClassDoesNotInherit(t *testing.T) {
	c := Class{Name: "Cat"}
	d := Class{Name: "Tiger"}
	goutil.AssertNow(t, !assignableTo(d, c), "class should not be assignable")
}

func TestIsAssignableClassFlipped(t *testing.T) {
	d := Class{Name: "Tiger"}
	c := Class{Name: "Cat", Supers: []*Class{&d}}
	goutil.AssertNow(t, !assignableTo(d, c), "class should not be assignable")
}

func TestIsAssignableClassInterfaceNot(t *testing.T) {
	c := Class{Name: "Cat"}
	d := Interface{Name: "Tiger"}
	goutil.AssertNow(t, !assignableTo(d, c), "class should not be assignable")
	goutil.AssertNow(t, !assignableTo(c, d), "interface should not be assignable")
}
