package validator

import (
	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) isVarDeclaredInScope(ts *TypeScope, name string) (typing.Type, bool) {
	if ts == nil {
		return typing.Invalid(), false
	}
	if ts.variables != nil {
		if t, ok := ts.variables[name]; ok {
			return t, true
		}
	}
	if ts.context != nil {
		// check parents
		switch c := ts.context.(type) {
		case *ast.ClassDeclarationNode:
			if t, ok := v.getClassProperty(util.Location{}, c.Resolved.(*typing.Class), name); ok {
				return t, true
			}
			break
		case *ast.ContractDeclarationNode:
			if t, ok := v.getContractProperty(util.Location{}, c.Resolved.(*typing.Contract), name); ok {
				return t, true
			}
			break
		}
	}
	if ts.parent != nil {
		return v.isVarDeclaredInScope(ts.parent, name)
	}
	return typing.Unknown(), false
}

func (v *Validator) isVarVisibleInScope(ts *TypeScope, name string) (typing.Type, bool) {
	if ts == nil {
		return typing.Invalid(), false
	}
	if ts.variables != nil {
		if t, ok := ts.variables[name]; ok {
			return t, true
		}
	}
	if ts.context != nil {
		// check parents
		switch c := ts.context.(type) {
		case *ast.ClassDeclarationNode:
			if t, ok := v.getClassProperty(util.Location{}, c.Resolved.(*typing.Class), name); ok {
				return t, true
			}
			break
		case *ast.ContractDeclarationNode:
			if t, ok := v.getContractProperty(util.Location{}, c.Resolved.(*typing.Contract), name); ok {
				return t, true
			}
			break
		}
	}
	if ts.scopes != nil {
		for _, s := range ts.scopes {
			if s != nil {
				decl := s.GetDeclaration(name)
				if decl != nil {
					saved := v.scope
					v.scope = ts
					v.validateDeclaration(decl)
					v.scope = saved
					if t, ok := ts.variables[name]; ok {
						return t, true
					}
				}
			}
		}
	}
	if ts.parent != nil {
		return v.isVarVisibleInScope(ts.parent, name)
	}
	return typing.Unknown(), false
}

func (v *Validator) isVarDirectlyVisibleInScope(ts *TypeScope, name string) (typing.Type, bool) {
	if ts == nil {
		return typing.Invalid(), false
	}
	if ts.variables != nil {
		if t, ok := ts.variables[name]; ok {
			return t, true
		}
	}
	if ts.scopes != nil {
		for _, s := range ts.scopes {
			if s != nil {
				decl := s.GetDeclaration(name)
				if decl != nil {
					saved := v.scope
					v.scope = ts
					v.validateDeclaration(decl)
					v.scope = saved
					if t, ok := ts.variables[name]; ok {
						return t, true
					}
				}
			}
		}
	}
	if ts.parent != nil {
		return v.isVarDirectlyVisibleInScope(ts.parent, name)
	}
	return typing.Unknown(), false
}

func (v *Validator) isVarDirectlyVisible(name string) (typing.Type, bool) {
	if v.builtinScope != nil {
		if t, ok := v.builtinScope.variables[name]; ok {
			return t, ok
		}
	}
	if t, ok := v.isVarDirectlyVisibleInScope(v.scope, name); ok {
		return t, ok
	}
	return typing.Unknown(), false
}

func (v *Validator) isVarDeclared(name string) (typing.Type, bool) {
	if v.builtinScope != nil {
		if t, ok := v.builtinScope.variables[name]; ok {
			return t, ok
		}
	}
	if t, ok := v.isVarDeclaredInScope(v.scope, name); ok {
		return t, ok
	}
	return typing.Unknown(), false
}

func (v *Validator) isVarVisible(name string) (typing.Type, bool) {
	if v.builtinScope != nil {
		if t, ok := v.builtinScope.variables[name]; ok {
			return t, ok
		}
	}
	if t, ok := v.isVarVisibleInScope(v.scope, name); ok {
		return t, ok
	}
	return typing.Unknown(), false
}

func (v *Validator) isTypeDeclaredInScope(ts *TypeScope, name string) (typing.Type, bool) {
	if ts == nil {
		return typing.Invalid(), false
	}
	if ts.types != nil {
		if t, ok := ts.types[name]; ok {
			return t, true
		}
	}
	if ts.context != nil {
		// check parents
		switch c := ts.context.(type) {
		case *ast.ClassDeclarationNode:
			if t, ok := c.Resolved.(*typing.Class).Types[name]; ok {
				return t, true
			}
			break
		case *ast.ContractDeclarationNode:
			if t, ok := c.Resolved.(*typing.Contract).Types[name]; ok {
				return t, true
			}
			break
		}
	}
	if ts.parent != nil {
		return v.isTypeDeclaredInScope(ts.parent, name)
	}
	return typing.Unknown(), false
}

func (v *Validator) isTypeVisibleInScope(ts *TypeScope, name string) (typing.Type, bool) {
	if ts == nil {
		return typing.Invalid(), false
	}
	if ts.types != nil {
		if t, ok := ts.types[name]; ok {
			return t, true
		}
	}
	if ts.context != nil {
		// check parents
		switch c := ts.context.(type) {
		case *ast.ClassDeclarationNode:
			if t, ok := c.Resolved.(*typing.Class).Types[name]; ok {
				return t, true
			}
			break
		case *ast.ContractDeclarationNode:
			if t, ok := c.Resolved.(*typing.Contract).Types[name]; ok {
				return t, true
			}
			break
		}
	}
	if ts.scopes != nil {
		for _, s := range ts.scopes {
			if s != nil {
				decl := s.GetDeclaration(name)
				if decl != nil {
					saved := v.scope
					v.scope = ts
					v.validateDeclaration(decl)
					v.scope = saved
					if t, ok := ts.types[name]; ok {
						return t, true
					}
				}
			}
		}
	}

	if ts.parent != nil {
		return v.isTypeVisibleInScope(ts.parent, name)
	}
	return typing.Unknown(), false
}

func (v *Validator) isTypeDeclared(name string) (typing.Type, bool) {
	if v.primitives != nil {
		if t, ok := v.primitives[name]; ok {
			return t, ok
		}
	}
	if v.builtinScope != nil {
		if t, ok := v.builtinScope.types[name]; ok {
			return t, ok
		}
	}
	if t, ok := v.isTypeDeclaredInScope(v.scope, name); ok {
		return t, ok
	}
	return typing.Unknown(), false
}

func (v *Validator) isTypeVisible(name string) (typing.Type, bool) {
	if v.primitives != nil {
		if t, ok := v.primitives[name]; ok {
			return t, ok
		}
	}
	if v.builtinScope != nil {
		if t, ok := v.builtinScope.types[name]; ok {
			return t, ok
		}
	}
	if t, ok := v.isTypeVisibleInScope(v.scope, name); ok {
		return t, ok
	}
	return typing.Unknown(), false
}

func (v *Validator) declareVar(loc util.Location, name string, typ typing.Type) {
	if _, ok := v.isVarDeclared(name); ok {
		v.addError(loc, errDuplicateVarDeclaration, name)
		return
	}
	if v.scope.variables == nil {
		v.scope.variables = make(typing.TypeMap)
	}
	v.scope.variables[name] = typ
}

func (v *Validator) declareType(loc util.Location, name string, typ typing.Type) {
	if _, ok := v.isTypeDeclared(name); ok {
		v.addError(loc, errDuplicateTypeDeclaration, name)
	}
	if v.scope.types == nil {
		v.scope.types = make(typing.TypeMap)
	}
	v.scope.types[name] = typ
}

func (v *Validator) declareLifecycle(tk token.Type, l typing.Lifecycle) {
	if v.scope.lifecycles == nil {
		v.scope.lifecycles = make(typing.LifecycleMap)
	}
	v.scope.lifecycles[tk] = append(v.scope.lifecycles[tk], l)
}

func (v *Validator) requireType(loc util.Location, expected, actual typing.Type) bool {
	e := typing.ResolveUnderlying(expected)
	a := typing.ResolveUnderlying(actual)
	if !e.Compare(a) {
		// bool and (bool) are equivalent, for example
		if t, ok := a.(*typing.Tuple); ok && len(t.Types) == 1 {
			if e.Compare(t.Types[0]) {
				return true
			}
		}
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
