package validator

import (
	"fmt"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) requireVisibleType(names ...string) typing.Type {
	typ := v.getNamedType(names...)
	if typ == typing.Unknown() {
		v.addError(errTypeNotVisible, makeName(names))
	}
	return typ
}

func (v *Validator) findVariable(name string) typing.Type {
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
				if c.Resolved == nil {
					fmt.Println("nil resolved")
				}
				t, _ := v.getTypeProperty(c.Resolved, name)
				return t
			case *ast.EnumDeclarationNode:
				t, _ := v.getTypeProperty(c.Resolved, name)
				return t
				break
			case *ast.InterfaceDeclarationNode:
				t, _ := v.getTypeProperty(c.Resolved, name)
				return t
				break
			case *ast.ContractDeclarationNode:
				t, _ := v.getTypeProperty(c.Resolved, name)
				return t
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
func (v *Validator) DeclareVarOfType(name string, t typing.Type) {
	if v.scope.variables == nil {
		v.scope.variables = make(map[string]typing.Type)
	}
	v.scope.variables[name] = t
}

// DeclareBuiltinOfType ...
func (v *Validator) DeclareBuiltinOfType(name string, t typing.Type) {
	if v.builtinVariables == nil {
		v.builtinVariables = make(map[string]typing.Type)
	}
	v.builtinVariables[name] = t
}

// DeclareType ...
func (v *Validator) DeclareType(name string, t typing.Type) {
	if v.scope.types == nil {
		v.scope.types = make(map[string]typing.Type)
	}
	v.scope.types[name] = t
}

func (v *Validator) declareLifecycle(tk token.Type, l typing.Lifecycle) {
	if v.scope.lifecycles == nil {
		v.scope.lifecycles = typing.LifecycleMap{}
	}
	v.scope.lifecycles[tk] = append(v.scope.lifecycles[tk], l)
}

// DeclareBuiltinType ...
func (v *Validator) DeclareBuiltinType(name string, t typing.Type) {
	if v.primitives == nil {
		v.primitives = make(map[string]typing.Type)
	}
	v.primitives[name] = t
}

func (v *Validator) getNamedType(names ...string) typing.Type {
	search := names[0]
	if v.primitives != nil {
		for k, typ := range v.primitives {
			if k == search {
				// found top level type
				//return v.getPropertiesType(typ, names[1:])
				return typ
			}
		}
	}
	for s := v.scope; s != nil; s = s.parent {
		if s.types != nil {
			for k, typ := range s.types {
				if k == search {
					// found top level type
					//return v.getPropertiesType(typ, names[1:])
					return typ
				}
			}
		}

	}
	return typing.Unknown()
}

func (v *Validator) requireType(expected, actual typing.Type) bool {
	if typing.ResolveUnderlying(expected) != typing.ResolveUnderlying(actual) {
		v.addError(errRequiredType, typing.WriteType(expected), typing.WriteType(actual))
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
