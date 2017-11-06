package validator

import (
	"fmt"

	"github.com/end-r/guardian/compiler/ast"
)

// ValidateScope validates an ast...
func ValidateScope(scope *ast.ScopeNode) *Validator {
	v := new(Validator)
	v.validateScope(scope)
	return v
}

func (v *Validator) validateScope(scope *ast.ScopeNode) map[string]Type {

	ts := &TypeScope{
		parent:    v.scope,
		scope:     scope,
		types:     nil,
		variables: nil,
	}

	v.scope = ts

	v.scanDeclarations(scope)

	v.validateDeclarations(scope)

	v.validateSequence(scope)

	properties := v.scope.types

	v.scope = v.scope.parent

	return properties
}

func (v *Validator) scanDeclarations(scope *ast.ScopeNode) {
	if scope.Declarations != nil {

		// there should be no declarations outside certain contexts
		for k, i := range scope.Declarations.Map() {
			// add in placeholders for all declarations
			v.DeclareType(k, v.resolveType(i.(ast.Node)))
		}
	}
}

func (v *Validator) validateDeclarations(scope *ast.ScopeNode) {
	if scope.Declarations != nil {
		for _, i := range scope.Declarations.Array() {
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

type Validator struct {
	scope  *TypeScope
	errors []string
}

type TypeScope struct {
	parent    *TypeScope
	scope     *ast.ScopeNode
	variables map[string]Type
	types     map[string]Type
}

func NewValidator() *Validator {
	return &Validator{
		scope: new(TypeScope),
	}
}

func (v *Validator) requireVisibleType(names ...string) Type {
	typ := v.findReference(names...)
	if typ == standards[Invalid] {
		v.addError(errTypeNotVisible, makeName(names))
	}
	return typ
}

func (v *Validator) DeclareVarOfType(name string, t Type) {
	if v.scope.variables == nil {
		v.scope.variables = make(map[string]Type)
	}
	v.scope.variables[name] = t
}

func (v *Validator) DeclareType(name string, t Type) {
	if v.scope.types == nil {
		v.scope.types = make(map[string]Type)
	}
	v.scope.types[name] = t
}

func (v *Validator) findReference(names ...string) Type {
	search := makeName(names)
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
					return typ
				}
			}
		}
	}
	return standards[Invalid]
}

func makeName(names []string) string {
	name := ""
	for i, n := range names {
		if i > 0 {
			name += "."
		}
		name += n
	}
	return name
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
