package validator

import (
	"fmt"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) requireVisibleType(loc util.Location, names ...string) typing.Type {
	typ := v.getNamedType(names[0])
	if typ == typing.Unknown() {
		v.addError(loc, errTypeNotVisible, makeName(names))
	}
	for _, n := range names[0:] {
		typ, ok := v.getTypeType(loc, typ, n)
		if !ok {
			break
		}
	}
	return typ
}

func (v *Validator) findVariable(loc util.Location, name string) typing.Type {
	if v.builtinVariables != nil {
		if typ, ok := v.builtinVariables[name]; ok {
			return typ
		}
	}
	for scope := v.scope; scope != nil; scope = scope.parent {
		if scope.context != nil {
			// check parents
			switch c := scope.context.(type) {
			case *ast.ClassDeclarationNode:
				if t, ok := c.Resolved.(*typing.Class).Properties[name]; ok {
					return t
				}
				break
			case *ast.EnumDeclarationNode:
				t, ok := v.getEnumProperty(loc, c, name)
				if ok {
					return t
				}
				break
			case *ast.InterfaceDeclarationNode:
				if t, ok := c.Resolved.(*typing.Interface).Funcs[name]; ok {
					return t
				}
				break
			case *ast.ContractDeclarationNode:
				if t, ok := c.Resolved.(*typing.Contract).Properties[name]; ok {
					return t
				}
				break
			}
		}
		if scope.variables != nil {
			if typ, ok := scope.variables[name]; ok {
				return typ
			}
		}
		if scope.scope != nil {
			if a := scope.scope.GetDeclaration(name); a != nil {
				saved := v.scope
				v.scope = scope
				v.validateDeclaration(a)
				v.scope = saved
				if t, ok := scope.variables[name]; ok {
					return t
				}
			}
		}

	}
	return typing.Unknown()
}

// DeclareVarOfType ...
func (v *Validator) DeclareVarOfType(loc util.Location, name string, t typing.Type) {
	if v.scope.variables == nil {
		v.scope.variables = make(map[string]typing.Type)
	}
	if v.findVariable(loc, name) != typing.Unknown() {
		v.addError(loc, errDuplicateVarDeclaration, name)
	} else {
		fmt.Println("declaring", name)
		v.scope.variables[name] = t
	}
}

// DeclareBuiltinOfType ...
func (v *Validator) DeclareBuiltinOfType(loc util.Location, name string, t typing.Type) {
	if v.builtinVariables == nil {
		v.builtinVariables = make(map[string]typing.Type)
	}
	if v.findVariable(loc, name) != typing.Unknown() {
		v.addError(loc, errDuplicateVarDeclaration, name)
	} else {
		v.builtinVariables[name] = t
	}
}

// DeclareBuiltinType ...
func (v *Validator) DeclareBuiltinType(loc util.Location, name string, t typing.Type) {
	if v.primitives == nil {
		v.primitives = make(map[string]typing.Type)
	}
	if v.getNamedType(name) != typing.Unknown() {
		v.addError(loc, errDuplicateTypeDeclaration, name)
	} else {
		v.primitives[name] = t
	}
}

// DeclareType ...
func (v *Validator) DeclareType(loc util.Location, name string, t typing.Type) {
	if v.scope.types == nil {
		v.scope.types = make(map[string]typing.Type)
	}
	if v.getNamedType(name) != typing.Unknown() {
		v.addError(loc, errDuplicateTypeDeclaration, name)
	} else {
		v.scope.types[name] = t
	}
}

func (v *Validator) declareLifecycle(tk token.Type, l typing.Lifecycle) {
	if v.scope.lifecycles == nil {
		v.scope.lifecycles = make(typing.LifecycleMap)
	}
	v.scope.lifecycles[tk] = append(v.scope.lifecycles[tk], l)
}

func (v *Validator) getNamedType(search string) typing.Type {
	if v.primitives != nil {
		for k, typ := range v.primitives {
			if k == search {
				return typ
			}
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
	return typing.Unknown()
}

func (v *Validator) requireType(loc util.Location, expected, actual typing.Type) bool {
	if typing.ResolveUnderlying(expected) != typing.ResolveUnderlying(actual) {
		v.addError(loc, errRequiredType, typing.WriteType(expected), typing.WriteType(actual))
		return false
	}
	return true
}

func (v *Validator) ExpressionTuple(exprs []ast.ExpressionNode) *typing.Tuple {
	var types []typing.Type
	for _, expression := range exprs {
		typ := v.resolveExpression(expression)
		// expression tuples force inner tuples to just be lists of types
		// ((int, string)) --> (int, string)
		// ((int), string) --> (int, string)
		// this is to facilitate assignment comparisons
		if tuple, ok := typ.(*typing.Tuple); ok {
			types = append(types, tuple.Types...)
		} else {
			types = append(types, typ)
		}
	}
	return typing.NewTuple(types...)
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
