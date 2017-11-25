package validator

import (
	"fmt"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/parser"
)

// Validate...
func Validate(scope *ast.ScopeNode, vm guardian.VMImplementation) {
	v := new(Validator)
	ts := &TypeScope{
		parent: nil,
		scope:  scope,
	}

	for name, typ := range vm.Primitives() {
		v.DeclareType(name, typ)
	}

	builtins := parser.ParseString(vm.GetBuiltins()).Scope

	v.isParsingBuiltins = true

	if builtins.Declarations != nil {
		for _, i := range builtins.Declarations.Map() {
			// add in placeholders for all declarations
			v.validateDeclaration(i.(ast.Node))
		}
	}

	if builtins.Sequence != nil {
		for _, node := range builtins.Sequence {
			v.validate(node)
		}
	}

	v.isParsingBuiltins = false

	v.scope = ts

}

// ValidateScope validates an ast...
func ValidateScope(scope *ast.ScopeNode) *Validator {
	v := new(Validator)
	v.validateScope(scope)
	return v
}

func (v *Validator) validateScope(scope *ast.ScopeNode) (map[string]Type, map[string]Type) {

	ts := &TypeScope{
		parent: v.scope,
		scope:  scope,
	}

	v.scope = ts

	v.isParsingBuiltins = false

	v.validateDeclarations(scope)

	v.validateSequence(scope)

	types := v.scope.types
	properties := v.scope.variables

	v.scope = v.scope.parent

	return types, properties
}

func (v *Validator) validateDeclarations(scope *ast.ScopeNode) {
	if scope.Declarations != nil {
		// order doesn't matter here
		for _, i := range scope.Declarations.Map() {
			// add in placeholders for all declarations
			v.validateDeclaration(i.(ast.Node))
		}
	}
}

func (v *Validator) validateSequence(scope *ast.ScopeNode) {
	for _, node := range scope.Sequence {
		v.validate(node)
	}
}

func (v *Validator) validate(node ast.Node) {
	if node.Type() == ast.CallExpression {
		v.validateCallExpression(node.(ast.CallExpressionNode))
	} else {
		v.validateStatement(node)
	}
}

// Validator ...
type Validator struct {
	scope             *TypeScope
	errors            []string
	isParsingBuiltins bool
}

// TypeScope ...
type TypeScope struct {
	parent           *TypeScope
	scope            *ast.ScopeNode
	variables        map[string]Type
	types            map[string]Type
	builtinVariables map[string]Type
	builtinTypes     map[string]Type
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{
		scope: new(TypeScope),
	}
}

func (v *Validator) requireVisibleType(names ...string) Type {
	typ := v.getNamedType(names...)
	if typ == standards[Unknown] {
		v.addError(errTypeNotVisible, makeName(names))
	}
	return typ
}

func (v *Validator) findVariable(name string) Type {
	for scope := v.scope; scope != nil; scope = scope.parent {
		if scope.builtinVariables != nil {
			if typ, ok := scope.builtinVariables[name]; ok {
				return typ
			}
		}
		if scope.variables != nil {
			if typ, ok := scope.variables[name]; ok {
				return typ
			}
		}
	}
	return standards[Unknown]
}

// DeclareVarOfType ...
func (v *Validator) DeclareVarOfType(name string, t Type) {
	if v.scope.variables == nil {
		v.scope.variables = make(map[string]Type)
	}
	v.scope.variables[name] = t
}

// DeclareBuiltinOfType ...
func (v *Validator) DeclareBuiltinOfType(name string, t Type) {
	if v.scope.builtinVariables == nil {
		v.scope.builtinVariables = make(map[string]Type)
	}
	v.scope.builtinVariables[name] = t
}

// DeclareType ...
func (v *Validator) DeclareType(name string, t Type) {
	if v.scope.types == nil {
		v.scope.types = make(map[string]Type)
	}
	v.scope.types[name] = t
}

// DeclareBuiltinType ...
func (v *Validator) DeclareBuiltinType(name string, t Type) {
	if v.scope.builtinTypes == nil {
		v.scope.builtinTypes = make(map[string]Type)
	}
	v.scope.builtinTypes[name] = t
}

func (v *Validator) getNamedType(names ...string) Type {
	search := names[0]
	// always check standards first
	// not declaring them in top scope means not having to go up each time
	// can simply go to the local scope
	for _, s := range standards {
		if search == s.name {
			return s
		}
	}
	for s := v.scope; s != nil; s = s.parent {
		if s.types != nil {
			for k, typ := range s.types {
				if k == search {
					// found top level type
					pType, ok := getPropertiesType(typ, names[1:])
					if !ok {

					}
					return pType
				}
			}
		}
	}
	return standards[Unknown]
}

func (v *Validator) requireType(expected, actual Type) bool {
	if resolveUnderlying(expected) != resolveUnderlying(actual) {
		v.addError("required type %s, got %s", WriteType(expected), WriteType(actual))
		return false
	}
	return true
}

func (v *Validator) addError(err string, data ...interface{}) {
	v.errors = append(v.errors, fmt.Sprintf(err, data...))
}

func (v *Validator) formatErrors() string {
	whole := ""
	whole += fmt.Sprintf("%d errors\n", len(v.errors))
	for _, e := range v.errors {
		whole += fmt.Sprintf("%s\n", e)
	}
	return whole
}
