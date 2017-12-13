package typing

import (
	"axia/guardian/token"
	"bytes"
)

// There are 5 first-class guardian types:
// Literal: int, string etc.
// Array: arrays[Type]
// NOTE: array = golang's slice, there is no golang array equivalent
// Map: map[Type]Type
// Func: func(Tuple)Tuple

// There are 2 second-class guardian types:
// Tuple: (Type...)
// Aliased: string -> Type

type Type interface {
	write(*bytes.Buffer)
	Compare(Type) bool
	inherits(Type) bool
	implements(Type) bool
	Size() uint
}
type LifecycleMap map[token.Type][]Lifecycle

type baseType int

const (
	invalid baseType = iota
	unknown
	boolean
)

type StandardType struct {
	name string
}

func Invalid() Type {
	return standards[invalid]
}

func Unknown() Type {
	return standards[unknown]
}

func Boolean() Type {
	return standards[boolean]
}

var standards = map[baseType]StandardType{
	invalid: StandardType{"invalid"},
	unknown: StandardType{"unknown"},
	boolean: StandardType{"bool"},
}

type Array struct {
	Length   int
	Value    Type
	Variable bool
}

type Map struct {
	Key   Type
	Value Type
}

type Func struct {
	Name    string
	Params  Tuple
	Results Tuple
}

type Tuple struct {
	Types []Type
}

func NewTuple(types ...Type) Tuple {
	return Tuple{
		Types: types,
	}
}

type Aliased struct {
	Alias      string
	Underlying Type
}

func ResolveUnderlying(t Type) Type {
	for al, ok := t.(Aliased); ok; al, ok = t.(Aliased) {
		t = al.Underlying
	}
	return t
}

type Lifecycle struct {
	Type       token.Type
	Parameters []Type
}

// A Class is a collection of properties
type Class struct {
	Name       string
	Lifecycles LifecycleMap
	Supers     []*Class
	Properties map[string]Type
	Types      map[string]Type
	Interfaces []*Interface
}

type Enum struct {
	Name   string
	Supers []*Enum
	Items  []string
}

type Interface struct {
	Name   string
	Supers []*Interface
	Funcs  map[string]Func
}

// Contract ...
type Contract struct {
	Name       string
	Supers     []*Contract
	Interfaces []*Interface
	Lifecycles map[token.Type][]Lifecycle
	Types      map[string]Type
	Properties map[string]Type
}

// Event ...
type Event struct {
	Name       string
	Parameters Tuple
}
