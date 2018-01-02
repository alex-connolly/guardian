package typing

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestCompareArraysExplicitlyEqual(t *testing.T) {
	one := &Array{Value: standards[boolean], Length: 0, Variable: true}
	two := &Array{Value: standards[boolean], Length: 0, Variable: true}
	goutil.Assert(t, one.Compare(two), "should be equal")
}

func TestCompareArraysImplicitlyEqual(t *testing.T) {
	one := &Array{Value: standards[boolean], Length: 0, Variable: true}
	two := &Aliased{
		Alias:      "a",
		Underlying: &Array{Value: standards[boolean], Length: 0, Variable: true},
	}
	goutil.Assert(t, one.Compare(two), "should be equal")
}

func TestCompareArraysExplicitlyWrongKey(t *testing.T) {
	one := &Array{Value: standards[boolean], Length: 0, Variable: true}
	two := &Array{Value: standards[invalid], Length: 0, Variable: true}
	goutil.Assert(t, !one.Compare(two), "should not be equal")
}

func TestCompareArraysImplicitlyWrongKey(t *testing.T) {
	one := &Array{Value: standards[boolean], Length: 0, Variable: true}
	two := &Aliased{
		Alias:      "a",
		Underlying: &Array{Value: standards[invalid], Length: 0, Variable: true},
	}
	goutil.Assert(t, !one.Compare(two), "should not be equal")
}

func TestCompareArraysExplicitlyWrongType(t *testing.T) {
	one := &Array{Value: standards[boolean], Length: 0, Variable: true}
	two := &Func{Params: NewTuple(), Results: NewTuple()}
	goutil.Assert(t, !one.Compare(two), "should not be equal")
}

func TestCompareArraysImplicitlyWrongType(t *testing.T) {
	one := &Array{Value: standards[boolean], Length: 0, Variable: true}
	two := &Aliased{
		Alias:      "a",
		Underlying: &Func{Params: NewTuple(), Results: NewTuple()},
	}
	goutil.Assert(t, !one.Compare(two), "should not be equal")
}

func TestCompareMapsExplicitlyEqual(t *testing.T) {
	one := &Map{Key: standards[boolean], Value: standards[boolean]}
	two := &Map{Key: standards[boolean], Value: standards[boolean]}
	goutil.Assert(t, one.Compare(two), "should be equal")
}

func TestCompareMapsImplicitlyEqual(t *testing.T) {
	one := &Map{Key: standards[boolean], Value: standards[boolean]}
	two := &Aliased{
		Alias:      "a",
		Underlying: &Map{Key: standards[boolean], Value: standards[boolean]},
	}
	goutil.Assert(t, one.Compare(two), "should be equal")
}

func TestCompareEmptyFuncs(t *testing.T) {
	one := &Func{Params: NewTuple(), Results: NewTuple()}
	two := &Aliased{
		Alias:      "a",
		Underlying: &Func{Params: NewTuple(), Results: NewTuple()},
	}
	goutil.Assert(t, one.Compare(two), "should be equal")
}

func TestCompareTuples(t *testing.T) {
	one := NewTuple()
	two := NewTuple()
	goutil.Assert(t, one.Compare(two), "should be equal")
}

func TestCompareTuplesWrongLength(t *testing.T) {
	one := NewTuple(standards[boolean], standards[unknown])
	two := NewTuple()
	goutil.Assert(t, !one.Compare(two), "should not be equal")
}

/*
func TestCompareTuplesWrongType(t *testing.T) {
	one := NewTuple(standards[boolean], standards[unknown])
	two := NewTuple(standards[unknown], standards[boolean])
	goutil.Assert(t, !one.Compare(two), "should not be equal")
}*/

func TestCompareStandards(t *testing.T) {
	one := standards[boolean]
	two := standards[boolean]
	goutil.Assert(t, one.Compare(two), "should be equal")
}

func TestCompareStandardsWrongType(t *testing.T) {
	one := standards[boolean]
	two := standards[unknown]
	goutil.Assert(t, !one.Compare(two), "should not be equal")
}

func TestCompareBooleans(t *testing.T) {
	one := Boolean()
	two := Boolean()
	goutil.Assert(t, one.Compare(two), "should be equal")
}

func TestCompareNamedBooleans(t *testing.T) {
	one := &Aliased{Alias: "a", Underlying: Boolean()}
	two := &Aliased{Alias: "b", Underlying: Boolean()}
	goutil.Assert(t, one.Compare(two), "should be equal")
}
