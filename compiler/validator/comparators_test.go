package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestCompareArraysExplicitlyEqual(t *testing.T) {
	one := NewArray(standards[Int])
	two := NewArray(standards[Int])
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareArraysImplicitlyEqual(t *testing.T) {
	one := NewArray(standards[Int])
	two := NewAliased("a", NewArray(standards[Int]))
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareArraysExplicitlyWrongKey(t *testing.T) {
	one := NewArray(standards[Int])
	two := NewArray(standards[String])
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareArraysImplicitlyWrongKey(t *testing.T) {
	one := NewArray(standards[Int])
	two := NewAliased("a", NewArray(standards[String]))
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareMapsExplicitlyEqual(t *testing.T) {
	one := NewMap(standards[Int], standards[Int])
	two := NewMap(standards[Int], standards[Int])
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareMapsImplicitlyEqual(t *testing.T) {
	one := NewMap(standards[Int], standards[Int])
	two := NewAliased("a", NewMap(standards[Int], standards[Int]))
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareFuncs(t *testing.T) {

}

func TestCompareTuples(t *testing.T) {

}

func TestCompareAliased(t *testing.T) {

}

func TestCompareStandard(t *testing.T) {

}
