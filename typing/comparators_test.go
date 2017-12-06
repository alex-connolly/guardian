package typing

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestCompareArraysExplicitlyEqual(t *testing.T) {
	one := Array{Value: standards[Bool], Length: 0, Variable: true}
	two := Array{Value: standards[Bool], Length: 0, Variable: true}
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareArraysImplicitlyEqual(t *testing.T) {
	one := Array{Value: standards[Bool], Length: 0, Variable: true}
	two := Aliased{
		Alias:      "a",
		Underlying: Array{Value: standards[Bool], Length: 0, Variable: true},
	}
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareArraysExplicitlyWrongKey(t *testing.T) {
	one := Array{Value: standards[Bool], Length: 0, Variable: true}
	two := Array{Value: standards[Invalid], Length: 0, Variable: true}
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareArraysImplicitlyWrongKey(t *testing.T) {
	one := Array{Value: standards[Bool], Length: 0, Variable: true}
	two := Aliased{
		Alias:      "a",
		Underlying: Array{Value: standards[Invalid], Length: 0, Variable: true},
	}
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareArraysExplicitlyWrongType(t *testing.T) {
	one := Array{Value: standards[Bool], Length: 0, Variable: true}
	two := Func{Params: NewTuple(), Results: NewTuple()}
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareArraysImplicitlyWrongType(t *testing.T) {
	one := Array{Value: standards[Bool], Length: 0, Variable: true}
	two := Aliased{
		Alias:      "a",
		Underlying: Func{Params: NewTuple(), Results: NewTuple()},
	}
	goutil.Assert(t, !one.compare(two), "should not be equal")
}

func TestCompareMapsExplicitlyEqual(t *testing.T) {
	one := Map{Key: standards[Bool], Value: standards[Bool]}
	two := Map{Key: standards[Bool], Value: standards[Bool]}
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareMapsImplicitlyEqual(t *testing.T) {
	one := Map{Key: standards[Bool], Value: standards[Bool]}
	two := Aliased{
		Alias:      "a",
		Underlying: Map{Key: standards[Bool], Value: standards[Bool]},
	}
	goutil.Assert(t, one.compare(two), "should be equal")
}

func TestCompareEmptyFuncs(t *testing.T) {
	one := Func{Params: NewTuple(), Results: NewTuple()}
	two := Aliased{
		Alias:      "a",
		Underlying: Func{Params: NewTuple(), Results: NewTuple()},
	}
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
