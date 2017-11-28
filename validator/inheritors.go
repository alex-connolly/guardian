package validator

func (c Class) inherits(t Type) bool {
	if other, ok := resolveUnderlying(t).(Class); ok {
		for _, super := range c.Supers {
			if super.compare(other) || super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (i Interface) inherits(t Type) bool {
	if other, ok := resolveUnderlying(t).(Interface); ok {
		for _, super := range i.Supers {
			if super.compare(other) || super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (e Enum) inherits(t Type) bool {
	if other, ok := resolveUnderlying(t).(Enum); ok {
		for _, super := range e.Supers {
			if super.compare(other) || super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (c Contract) inherits(t Type) bool {
	if other, ok := resolveUnderlying(t).(Contract); ok {
		for _, super := range c.Supers {
			if super.compare(other) || super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (a Aliased) inherits(t Type) bool {
	return resolveUnderlying(a).inherits(t)
}

// types which don't inherit or implement
func (s StandardType) inherits(t Type) bool { return false }
func (p Tuple) inherits(t Type) bool        { return false }
func (f Func) inherits(t Type) bool         { return false }
func (a Array) inherits(t Type) bool        { return false }
func (m Map) inherits(t Type) bool          { return false }
func (e Event) inherits(t Type) bool        { return false }

func (n NumericType) inherits(t Type) bool { return false }
func (n BooleanType) inherits(t Type) bool { return false }
