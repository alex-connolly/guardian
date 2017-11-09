package validator

func (c Class) implements(t Type) bool {
	if other, ok := resolveUnderlying(t).(Interface); !ok {
		return false
	} else {
		for _, ifc := range c.Interfaces {
			if ifc.compare(other) || ifc.inherits(other) {
				return true
			}
		}
		for _, super := range c.Supers {
			if super.implements(other) {
				return true
			}
		}
		return false
	}

}

func (c Contract) implements(t Type) bool {
	if other, ok := resolveUnderlying(t).(Interface); !ok {
		return false
	} else {
		for _, ifc := range c.Interfaces {
			if ifc.compare(other) || ifc.inherits(other) {
				return true
			}
		}
		for _, super := range c.Supers {
			if super.implements(other) {
				return true
			}
		}
		return false
	}
}

func (a Aliased) implements(t Type) bool {
	return resolveUnderlying(a).implements(t)
}

func (s StandardType) implements(t Type) bool { return false }
func (p Tuple) implements(t Type) bool        { return false }
func (f Func) implements(t Type) bool         { return false }
func (a Array) implements(t Type) bool        { return false }
func (m Map) implements(t Type) bool          { return false }
func (i Interface) implements(t Type) bool    { return false }
func (e Enum) implements(t Type) bool         { return false }
func (e Event) implements(t Type) bool        { return false }
