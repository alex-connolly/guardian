package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIsAssignableEqualTypes(t *testing.T) {
	a := standards[Int]
	b := standards[Int]
	goutil.AssertNow(t, assignableTo(b, a), "equal types should be assignable")
}

func TestIsAssignableSuperClass(t *testing.T) {
	a := NewClass("Dog", nil, nil, nil)
	b := NewClass("Cat", []*Class{&a}, nil, nil)
	goutil.AssertNow(t, assignableTo(b, a), "super class types should be assignable")
}

func TestIsAssignableMultipleSuperClass(t *testing.T) {
	a := NewClass("Dog", nil, nil, nil)
	b := NewClass("Cat", []*Class{&a}, nil, nil)
	c := NewClass("Rat", []*Class{&b}, nil, nil)
	goutil.AssertNow(t, assignableTo(c, a), "super class types should be assignable")
}

func TestIsAssignableParentInterface(t *testing.T) {
	a := NewInterface("Dog", nil, nil)
	b := NewInterface("Cat", nil, []*Interface{&a})
	goutil.AssertNow(t, assignableTo(b, a), "interface types should be assignable")
}

func TestIsAssignableClassImplementingInterface(t *testing.T) {
	a := NewInterface("Dog", nil, nil)
	b := NewClass("Cat", nil, []*Interface{&a}, nil)
	goutil.AssertNow(t, assignableTo(b, a), "interface types should be assignable")
}

func TestIsAssignableSuperClassImplementingInterface(t *testing.T) {
	a := NewInterface("Dog", nil, nil)
	b := NewClass("Cat", nil, []*Interface{&a}, nil)
	c := NewClass("Cat", []*Class{&b}, nil, nil)
	goutil.AssertNow(t, assignableTo(c, a), "interface types should be assignable")
}

func TestIsAssignableSuperClassImplementingSuperInterface(t *testing.T) {
	a := NewInterface("Dog", nil, nil)
	b := NewInterface("Lion", nil, []*Interface{&a})
	c := NewClass("Cat", nil, []*Interface{&b}, nil)
	d := NewClass("Tiger", []*Class{&c}, nil, nil)
	goutil.AssertNow(t, assignableTo(d, a), "type should be assignable")
}
