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
	Modifiers() *Modifiers
	SetModifiers(*Modifiers)
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
	Mods *Modifiers
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
	invalid: &StandardType{name: "invalid"},
	unknown: &StandardType{name: "unknown"},
	boolean: &StandardType{name: "bool"},
	void:    &StandardType{name: "void"},
}

// Array ...
type Array struct {
	Mods     *Modifiers
	Length   int
	Value    Type
	Variable bool
}

// Map ...
type Map struct {
	Mods  *Modifiers
	Key   Type
	Value Type
}

// Func ...
type Func struct {
	Mods     *Modifiers
	Name     string
	Generics []*Generic
	Params   *Tuple
	Results  *Tuple
}

type Tuple struct {
	Mods  *Modifiers
	Types []Type
}

func NewTuple(types ...Type) *Tuple {
	return &Tuple{
		Types: types,
	}
}

type Aliased struct {
	Mods       *Modifiers
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

type CancellationMap map[string]bool

// A Class is a collection of properties
type Class struct {
	Cancelled  CancellationMap
	Mods       *Modifiers
	Name       string
	Generics   []*Generic
	Lifecycles LifecycleMap
	Supers     []*Class
	Properties TypeMap
	Types      TypeMap
	Interfaces []*Interface
}

type Enum struct {
	Cancelled CancellationMap
	Mods      *Modifiers
	Name      string
	Supers    []*Enum
	Items     []string
}

type Interface struct {
	Cancelled CancellationMap
	Mods      *Modifiers
	Name      string
	Generics  []*Generic
	Supers    []*Interface
	Funcs     map[string]*Func
}

// Contract ...
type Contract struct {
	Cancelled  CancellationMap
	Mods       *Modifiers
	Name       string
	Generics   []*Generic
	Supers     []*Contract
	Interfaces []*Interface
	Lifecycles map[token.Type][]Lifecycle
	Types      TypeMap
	Properties TypeMap
}

type TypeMap map[string]Type

type Annotation struct {
	Name       string
	Parameters []string
	Required   int
}

// currently leaving these as arrays
// in the future, order might be important
type Modifiers struct {
	Annotations []*Annotation
	Modifiers   []string
}

func (m *Modifiers) AddAnnotation(a *Annotation) {
	if m.Annotations == nil {
		m.Annotations = make([]*Annotation, 0)
	}
	m.Annotations = append(m.Annotations, a)
}

func (m *Modifiers) AddModifier(mod string) {
	if m.Modifiers == nil {
		m.Modifiers = make([]string, 0)
	}
	m.Modifiers = append(m.Modifiers, mod)
}

func (m *Modifiers) Annotation(anno string) *Annotation {
	for _, annotation := range m.Annotations {
		if annotation.Name == anno {
			return annotation
		}
	}
	return nil
}

func (m *Modifiers) HasAnnotation(anno string) bool {
	for _, annotation := range m.Annotations {
		if annotation.Name == anno {
			return true
		}
	}
	return false
}

func (m *Modifiers) HasModifier(mod string) bool {
	for _, modifier := range m.Modifiers {
		if modifier == mod {
			return true
		}
	}
	return false
}

// Event ...
type Event struct {
	Mods       *Modifiers
	Name       string
	Generics   []*Generic
	Parameters *Tuple
}

type Package struct {
	Mods      *Modifiers
	Name      string
	Types     TypeMap
	Variables TypeMap
}
