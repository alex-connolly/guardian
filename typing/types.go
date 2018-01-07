package typing

import (
	"bytes"

	"github.com/end-r/guardian/token"
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

// LifecycleMap ...
type LifecycleMap map[token.Type][]Lifecycle

type baseType int

const (
	invalid baseType = iota
	unknown
	boolean
	void
)

// StandardType ...
type StandardType struct {
	name string
}

// Invalid ...
func Invalid() Type {
	return standards[invalid]
}

// Unknown ...
func Unknown() Type {
	return standards[unknown]
}

// Boolean ...
func Boolean() Type {
	return standards[boolean]
}

// Void ...
func Void() Type {
	return standards[void]
}

var standards = map[baseType]*StandardType{
	invalid: &StandardType{"invalid"},
	unknown: &StandardType{"unknown"},
	boolean: &StandardType{"bool"},
	void:    &StandardType{"void"},
}

// Array ...
type Array struct {
	Length   int
	Value    Type
	Variable bool
}

// Map ...
type Map struct {
	Modifiers
	Key   Type
	Value Type
}

// Func ...
type Func struct {
	Modifiers Modifiers
	Name      string
	Generics  []*Generic
	Params    *Tuple
	Results   *Tuple
}

type Tuple struct {
	Types []Type
}

func NewTuple(types ...Type) *Tuple {
	return &Tuple{
		Types: types,
	}
}

type Aliased struct {
	Alias      string
	Underlying Type
}

func ResolveUnderlying(t Type) Type {
	for al, ok := t.(*Aliased); ok; al, ok = t.(*Aliased) {
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
	Modifiers  Modifiers
	Name       string
	Generics   []*Generic
	Lifecycles LifecycleMap
	Supers     []*Class
	Properties map[string]Type
	Types      map[string]Type
	Interfaces []*Interface
}

type Enum struct {
	Modifiers Modifiers
	Name      string
	Supers    []*Enum
	Items     []string
}

type Interface struct {
	Modifiers Modifiers
	Name      string
	Generics  []*Generic
	Supers    []*Interface
	Funcs     map[string]*Func
}

// Contract ...
type Contract struct {
	Modifiers  Modifiers
	Name       string
	Generics   []*Generic
	Supers     []*Contract
	Interfaces []*Interface
	Lifecycles map[token.Type][]Lifecycle
	Types      map[string]Type
	Properties map[string]Type
}

type Annotation struct {
	Name       string
	Parameters []string
	Required   int
}

type Modifiers struct {
	Annotations []*Annotation
	Modifiers   []string
}

// Event ...
type Event struct {
	Modifiers  Modifiers
	Name       string
	Generics   []*Generic
	Parameters *Tuple
}
