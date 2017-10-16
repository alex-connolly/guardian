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

	switch scope.Type() {
	case ast.ContractDeclaration, ast.ClassDeclaration:
		// scope where order-less declarations can be made
		for k, node := range scope.Declarations {
			v.addDeclaration(k, node)
		}
		for _, node := range scope.Declarations {
			v.validateDeclaration(node)
		}
		break
	default:
		// scope where order is preserved
		for _, node := range scope.Sequence {
			v.validate(node)
		}
		break
	}
}

func (v *Validator) addDeclaration(name string, node ast.Node) {
	if v.scope.declaredTypes == nil {
		v.scope.declaredTypes = make(map[string]Type)
	}
	v.scope.declaredTypes[name] = v.resolveType(node)
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

func (v *Validator) requireVisibleType(names ...string) {
	if v.findReference(names...) == standards[Invalid] {
		v.addError("Type %s is not visible", makeName(names))
	}
}

// BUG: type lookup, should check that "a" is a valid type
// BUG: shouldn't return the underlying type --> abstraction
func (v *Validator) DeclareType(name string, t Type) {
	if v.scope.declaredTypes == nil {
		v.scope.declaredTypes = make(map[string]Type)
	}
	v.scope.declaredTypes[name] = NewAliased(name, t)
}

func (v *Validator) findReference(names ...string) Type {
	search := makeName(names)
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
	v.errors = append(v.errors, fmt.Sprintf(err, data))
}
