package validator

import (
	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func isVisible(ts *TypeScope, name string) (typing.Type, bool) {
	for scope := ts; scope != nil; scope = scope.parent {
		if scope.context != nil {
			// check parents
			switch c := scope.context.(type) {
			case *ast.ClassDeclarationNode:
				if t, ok := c.Resolved.(*typing.Class).Properties[name]; ok {
					return t, true
				}
				break
			case *ast.ContractDeclarationNode:
				if t, ok := c.Resolved.(*typing.Contract).Properties[name]; ok {
					return t, true
				}
				break
			}
		}
		if scope.variables != nil {
			if t, ok := scope.variables[name]; ok {
				return t, true
			}
		}
	}
	return typing.Unknown(), false
}

func (v *Validator) isVarVisible(name string) (typing.Type, bool) {
	if t, ok := isVisible(v.scope, name); ok {
		return t, ok
	}
	if t, ok := isVisible(v.builtinScope, name); ok {
		return t, ok
	}
	return typing.Unknown(), false
}

func (v *Validator) declareContextualVar(loc util.Location, name string, typ typing.Type) {
	if _, ok := v.isVarVisible(name); ok {
		v.addError(loc, errDuplicateVarDeclaration, name)
		return
	}
	if v.isParsingBuiltins {
		if v.builtinScope.variables == nil {
			v.builtinScope.variables = make(typing.TypeMap)
		}
		v.builtinScope.variables[name] = typ
	} else {
		if v.scope.variables == nil {
			v.scope.variables = make(typing.TypeMap)
		}
		v.scope.variables[name] = typ
	}
}

func (v *Validator) declareContextualType(loc util.Location, name string, typ typing.Type) {
	if v.getNamedType(name) != typing.Unknown() {
		v.addError(loc, errDuplicateTypeDeclaration, name)
		return
	}
	if v.isParsingBuiltins {
		if v.builtinScope.types == nil {
			v.builtinScope.types = make(typing.TypeMap)
		}
		v.builtinScope.types[name] = typ
	} else {
		if v.scope.types == nil {
			v.scope.types = make(typing.TypeMap)
		}
		v.scope.types[name] = typ
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
		if t, ok := v.primitives[search]; ok {
			return t
		}
	}
	for s := v.scope; s != nil; s = s.parent {
		if s.types != nil {
			if t, ok := s.types[search]; ok {
				return t
			}
		}
	}
	for s := v.builtinScope; s != nil; s = s.parent {
		if s.types != nil {
			if t, ok := s.types[search]; ok {
				return t
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
