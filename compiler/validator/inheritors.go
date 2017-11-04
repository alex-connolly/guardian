package validator

func (c Class) inherits(t Type) bool {
	other, ok := resolveUnderlying(t).(Class)
	if !ok {
		return false
	}
	for _, super := range c.Supers {
		if super.compare(other) {
			return true
		} else {
			if super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (i Interface) inherits(t Type) bool {
	other, ok := resolveUnderlying(t).(Interface)
	if !ok {
		return false
	}
	for _, super := range i.Supers {
		if super.compare(other) {
			return true
		} else {
			if super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (e Enum) inherits(t Type) bool {
	other, ok := resolveUnderlying(t).(Enum)
	if !ok {
		return false
	}
	for _, super := range e.Supers {
		if super.compare(other) {
			return true
		} else {
			if super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (c Contract) inherits(t Type) bool {
	other, ok := resolveUnderlying(t).(Contract)
	if !ok {
		return false
	}
	for _, super := range c.Supers {
		if super.compare(other) {
			return true
		} else {
			if super.inherits(other) {
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
