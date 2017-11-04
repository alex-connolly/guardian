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

func (v *Validator) validateScope(scope *ast.ScopeNode) {
	v.scope = &TypeScope{
		parent:        nil,
		scope:         scope,
		declaredTypes: nil,
	}
	if scope.Declarations != nil {
		// there should be no declarations outside certain contexts
		for k, i := range scope.Declarations.Map() {
			v.addDeclaration(k, i.(ast.Node))
		}
		for _, i := range scope.Declarations.Array() {
			v.validateDeclaration(i.(ast.Node))
		}
	}
	for _, node := range scope.Sequence {
		v.validate(node)
	}
}

func (v *Validator) addDeclaration(name string, node ast.Node) {
	if v.scope == nil {
		// TODO: error
	} else {
		if v.scope.declaredTypes == nil {
			v.scope.declaredTypes = make(map[string]Type)
		}
		v.scope.declaredTypes[name] = v.resolveType(node)
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
	parent        *TypeScope
	scope         *ast.ScopeNode
	declaredTypes map[string]Type
}

func NewValidator() *Validator {
	return &Validator{
		scope: new(TypeScope),
	}
}

func (v *Validator) requireVisibleType(names ...string) Type {
	typ := v.findReference(names...)
	if typ == standards[Invalid] {
		v.addError("Type %s is not visible", makeName(names))
	}
	return typ
}

// BUG: type lookup, should check that "a" is a valid type
// BUG: shouldn't return the underlying type --> abstraction
func (v *Validator) DeclareType(name string, t Type) {
	if v.scope.declaredTypes == nil {
		v.scope.declaredTypes = make(map[string]Type)
	}
	v.scope.declaredTypes[name] = t
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
		if s.declaredTypes != nil {
			for k, typ := range s.declaredTypes {
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
