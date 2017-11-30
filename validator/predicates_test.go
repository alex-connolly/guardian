package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIsAssignableEqualTypes(t *testing.T) {
	a := standards[Bool]
	b := standards[Bool]
	goutil.AssertNow(t, assignableTo(b, a), "equal types should be assignable")
}

func TestIsAssignableSuperClass(t *testing.T) {
	a := NewClass("Dog", nil, nil, nil, nil, nil)
	b := NewClass("Cat", []*Class{&a}, nil, nil, nil, nil)
	goutil.AssertNow(t, assignableTo(b, a), "super class types should be assignable")
}

func TestIsAssignableMultipleSuperClass(t *testing.T) {
	a := NewClass("Dog", nil, nil, nil, nil, nil)
	b := NewClass("Cat", []*Class{&a}, nil, nil, nil, nil)
	c := NewClass("Rat", []*Class{&b}, nil, nil, nil, nil)
	goutil.AssertNow(t, assignableTo(c, a), "super class types should be assignable")
}

func TestIsAssignableParentInterface(t *testing.T) {
	a := NewInterface("Dog", nil, nil)
	b := NewInterface("Cat", []*Interface{&a}, nil)
	goutil.AssertNow(t, assignableTo(b, a), "interface types should be assignable")
}

func TestIsAssignableClassImplementingInterface(t *testing.T) {
	a := NewInterface("Dog", nil, nil)
	b := NewClass("Cat", nil, []*Interface{&a}, nil, nil, nil)
	goutil.AssertNow(t, assignableTo(b, a), "interface types should be assignable")
}

func TestIsAssignableSuperClassImplementingInterface(t *testing.T) {
	a := NewInterface("Dog", nil, nil)
	b := NewClass("Cat", nil, []*Interface{&a}, nil, nil, nil)
	c := NewClass("Cat", []*Class{&b}, nil, nil, nil, nil)
	goutil.AssertNow(t, assignableTo(c, a), "interface types should be assignable")
}

func TestIsAssignableSuperClassImplementingSuperInterface(t *testing.T) {
	a := NewInterface("Dog", nil, nil)
	b := NewInterface("Lion", []*Interface{&a}, nil)
	c := NewClass("Cat", nil, []*Interface{&b}, nil, nil, nil)
	d := NewClass("Tiger", []*Class{&c}, nil, nil, nil, nil)
	goutil.AssertNow(t, assignableTo(d, a), "type should be assignable")
}

func TestIsAssignableClassDoesNotInherit(t *testing.T) {
	c := NewClass("Cat", nil, nil, nil, nil, nil)
	d := NewClass("Tiger", nil, nil, nil, nil, nil)
	goutil.AssertNow(t, !assignableTo(d, c), "class should be assignable")
}

func TestIsAssignableClassFlipped(t *testing.T) {
	d := NewClass("Tiger", nil, nil, nil, nil, nil)
	c := NewClass("Cat", []*Class{&d}, nil, nil, nil, nil)
	goutil.AssertNow(t, !assignableTo(d, c), "class should not be assignable")
}

func TestIsAssignableClassInterfaceNot(t *testing.T) {
	c := NewClass("Cat", nil, nil, nil, nil, nil)
	d := NewInterface("Tiger", nil, nil)
	goutil.AssertNow(t, !assignableTo(d, c), "class should not be assignable")
	goutil.AssertNow(t, !assignableTo(c, d), "interface should not be assignable")
}
