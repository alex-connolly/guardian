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
		// scope where declarations can be made
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
	//v.scope.declaredTypes[name]
}

func (v *Validator) validate(node ast.Node) {
	if node.Type() == ast.CallExpression {

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

func (v *Validator) resolveType(node ast.Node) Type {
	switch node.Type() {
	case ast.Reference:
		ref := node.(ast.ReferenceNode)
		return v.findReference(ref.Names...)
	case ast.MapType:
		m := node.(ast.MapTypeNode)
		return v.resolveMapType(m)
	case ast.ArrayType:
		a := node.(ast.ArrayTypeNode)
		return v.resolveArrayType(a)
	case ast.FuncType:
		f := node.(ast.FuncTypeNode)
		return v.resolveFuncType(f)
	}
	return standards[Invalid]
}

func (v *Validator) resolveArrayType(node ast.ArrayTypeNode) Array {
	a := Array{}
	a.Value = v.resolveType(node.Value)
	return a
}

func (v *Validator) resolveMapType(node ast.MapTypeNode) Map {
	m := Map{}
	m.Key = v.resolveType(node.Key)
	m.Value = v.resolveType(node.Value)
	return m
}

func (v *Validator) resolveFuncType(node ast.FuncTypeNode) Func {
	f := Func{}
	f.Params = v.resolveTuple(node.Parameters)
	f.Results = v.resolveTuple(node.Results)
	return f
}

func (v *Validator) resolveTuple(nodes []ast.Node) Tuple {
	t := Tuple{}
	t.types = make([]Type, len(nodes))
	for i, n := range nodes {
		t.types[i] = v.resolveType(n)
	}
	return t
}

func (v *Validator) validateType(node ast.Node) {
	switch node.Type() {
	case ast.Reference:
		ref := node.(ast.ReferenceNode)
		v.requireVisibleType(ref.Names...)
		break
	case ast.MapType:
		ref := node.(ast.MapTypeNode)
		v.validateType(ref.Key)
		v.validateType(ref.Value)
		break
	case ast.ArrayType:
		ref := node.(ast.ArrayTypeNode)
		v.validateType(ref.Value)
		break
	case ast.FuncType:
		ref := node.(ast.FuncTypeNode)
		for _, p := range ref.Parameters {
			v.validateType(p)
		}
		for _, r := range ref.Results {
			v.validateType(r)
		}
		break
	}
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
