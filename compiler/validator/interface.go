package validator

import "github.com/end-r/guardian/compiler/ast"

// ValidateScope validates an ast...
func ValidateScope(scope *ast.ScopeNode) {
	v := new(Validator)
	v.validateScope(scope)
}

func (v *Validator) validateScope(scope *ast.ScopeNode) {
	v.scope = &TypeScope{
		parent:        nil,
		scope:         scope,
		declaredTypes: nil,
	}
	validateTypeDeclarations(v)
}

type Validator struct {
	scope *TypeScope
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

func (v *Validator) DeclareType(name string, t Type) {
	if v.scope.declaredTypes == nil {
		v.scope.declaredTypes = make(map[string]Type)
	}
	v.scope.declaredTypes[name] = t
}

func (v *Validator) findReference(names ...string) Type {
	for s := v.scope; s != nil; s = s.parent {
		if s.declaredTypes != nil {
			for k, typ := range s.declaredTypes {
				if k == makeName(names) {
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
