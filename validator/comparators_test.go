package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestCompareArraysExplicitlyEqual(t *testing.T) {
	one := NewArray(standards[Bool], 0, true)
	two := NewArray(standards[Bool], 0, true)
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareArraysImplicitlyEqual(t *testing.T) {
	one := NewArray(standards[Bool], 0, true)
	two := NewAliased("a", NewArray(standards[Bool], 0, true))
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareArraysExplicitlyWrongKey(t *testing.T) {
	one := NewArray(standards[Bool], 0, true)
	two := NewArray(standards[Invalid], 0, true)
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareArraysImplicitlyWrongKey(t *testing.T) {
	one := NewArray(standards[Bool], 0, true)
	two := NewAliased("a", NewArray(standards[Invalid], 0, true))
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareArraysExplicitlyWrongType(t *testing.T) {
	one := NewArray(standards[Bool], 0, true)
	two := NewFunc(NewTuple(), NewTuple())
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareArraysImplicitlyWrongType(t *testing.T) {
	one := NewArray(standards[Bool], 0, true)
	two := NewAliased("a", NewFunc(NewTuple(), NewTuple()))
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareMapsExplicitlyEqual(t *testing.T) {
	one := NewMap(standards[Bool], standards[Bool])
	two := NewMap(standards[Bool], standards[Bool])
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareMapsImplicitlyEqual(t *testing.T) {
	one := NewMap(standards[Bool], standards[Bool])
	two := NewAliased("a", NewMap(standards[Bool], standards[Bool]))
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareEmptyFuncs(t *testing.T) {
	one := NewFunc(NewTuple(), NewTuple())
	two := NewAliased("a", NewFunc(NewTuple(), NewTuple()))
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareTuples(t *testing.T) {
	one := NewTuple()
	two := NewTuple()
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareTuplesWrongLength(t *testing.T) {
	one := NewTuple(standards[Bool], standards[Unknown])
	two := NewTuple()
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareTuplesWrongType(t *testing.T) {
	one := NewTuple(standards[Bool], standards[Unknown])
	two := NewTuple(standards[Unknown], standards[Bool])
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareStandards(t *testing.T) {
	one := standards[Bool]
	two := standards[Bool]
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareStandardsWrongType(t *testing.T) {
	one := standards[Bool]
	two := standards[Unknown]
	goutil.Assert(t, !one.compare(two), "should not be equal")
}
