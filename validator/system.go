package validator

import "github.com/end-r/guardian/lexer"

func (v *Validator) requireVisibleType(names ...string) Type {
	typ := v.getNamedType(names...)
	if typ == standards[Unknown] {
		v.addError(errTypeNotVisible, makeName(names))
	}
	return typ
}

func (v *Validator) findVariable(name string) Type {
	if v.builtinVariables != nil {
		if typ, ok := v.builtinVariables[name]; ok {
			return typ
		}
	}
	for scope := v.scope; scope != nil; scope = scope.parent {
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
	if v.builtinVariables == nil {
		v.builtinVariables = make(map[string]Type)
	}
	v.builtinVariables[name] = t
}

// DeclareType ...
func (v *Validator) DeclareType(name string, t Type) {
	if v.scope.types == nil {
		v.scope.types = make(map[string]Type)
	}
	v.scope.types[name] = t
}

func (v *Validator) declareLifecycle(tk lexer.TokenType, l Lifecycle) {
	if v.scope.lifecycles == nil {
		v.scope.lifecycles = lifecycleMap{}
	}
	v.scope.lifecycles[tk] = append(v.scope.lifecycles[tk], l)
}

// DeclareBuiltinType ...
func (v *Validator) DeclareBuiltinType(name string, t Type) {
	if v.primitives == nil {
		v.primitives = make(map[string]Type)
	}
	v.primitives[name] = t
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
	if v.primitives != nil {
		for k, typ := range v.primitives {
			if k == search {
				// found top level type
				return v.getPropertiesType(typ, names[1:])
			}
		}
	}
	for s := v.scope; s != nil; s = s.parent {
		if s.types != nil {
			for k, typ := range s.types {
				if k == search {
					// found top level type
					return v.getPropertiesType(typ, names[1:])
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
