package typing

func (c *Class) inherits(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Class); ok {
		for _, super := range c.Supers {
			if super.Compare(other) || super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (i *Interface) inherits(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Interface); ok {
		for _, super := range i.Supers {
			if super.Compare(other) || super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (e *Enum) inherits(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Enum); ok {
		for _, super := range e.Supers {
			if super.Compare(other) || super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (c *Contract) inherits(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Contract); ok {
		for _, super := range c.Supers {
			if super.Compare(other) || super.inherits(other) {
				return true
			}
		}
	}
	return false
}

func (a *Aliased) inherits(t Type) bool {
	return ResolveUnderlying(a).inherits(t)
}

// types which don't inherit or implement
func (s *StandardType) inherits(t Type) bool { return false }
func (p *Tuple) inherits(t Type) bool        { return false }
func (f *Func) inherits(t Type) bool         { return false }
func (a *Array) inherits(t Type) bool        { return false }
func (m *Map) inherits(t Type) bool          { return false }
func (e *Event) inherits(t Type) bool        { return false }
func (g *Generic) inherits(t Type) bool      { return false }
func (g *Package) inherits(t Type) bool      { return false }

func (n *NumericType) inherits(t Type) bool { return false }
func (n *BooleanType) inherits(t Type) bool { return false }
func (v *VoidType) inherits(t Type) bool    { return false }
