package typing

import (
	"bytes"

	"github.com/end-r/guardian/lexer"
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
	compare(Type) bool
	inherits(Type) bool
	implements(Type) bool
	size() int
}
type LifecycleMap map[lexer.TokenType][]Lifecycle

type BaseType int

const (
	Invalid BaseType = iota
	Unknown
	Bool
)

type StandardType struct {
	name string
}

var standards = map[BaseType]StandardType{
	Invalid: StandardType{"invalid"},
	Unknown: StandardType{"unknown"},
	Bool:    StandardType{"bool"},
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
	types []Type
}

type Aliased struct {
	alias      string
	underlying Type
}

type Lifecycle struct {
	Type       lexer.TokenType
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
	Lifecycles map[lexer.TokenType][]Lifecycle
	Types      map[string]Type
	Properties map[string]Type
}

// Event ...
type Event struct {
	Name       string
	Parameters Tuple
}
