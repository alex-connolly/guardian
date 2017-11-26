package validator

func (v *Validator) requireVisibleType(names ...string) Type {
	typ := v.getNamedType(names...)
	if typ == standards[Unknown] {
		v.addError(errTypeNotVisible, makeName(names))
	}
	return typ
}

func (v *Validator) findVariable(name string) Type {
	for scope := v.scope; scope != nil; scope = scope.parent {
		if scope.builtinVariables != nil {
			if typ, ok := scope.builtinVariables[name]; ok {
				return typ
			}
		}
		if scope.variables != nil {
			if typ, ok := scope.variables[name]; ok {
				return typ
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
	if v.scope.builtinVariables == nil {
		v.scope.builtinVariables = make(map[string]Type)
	}
	v.scope.builtinVariables[name] = t
}

// DeclareType ...
func (v *Validator) DeclareType(name string, t Type) {
	if v.scope.types == nil {
		v.scope.types = make(map[string]Type)
	}
	v.scope.types[name] = t
}

// DeclareBuiltinType ...
func (v *Validator) DeclareBuiltinType(name string, t Type) {
	if v.scope.builtinTypes == nil {
		v.scope.builtinTypes = make(map[string]Type)
	}
	v.scope.builtinTypes[name] = t
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
	for s := v.scope; s != nil; s = s.parent {
		if s.types != nil {
			for k, typ := range s.types {
				if k == search {
					// found top level type
					pType, ok := getPropertiesType(typ, names[1:])
					if !ok {

					}
					return pType
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
