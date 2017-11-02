package validator

func (c Class) implements(i Interface) bool {
	for _, ifc := range c.Interfaces {
		if ifc.compare(i) {
			return true
		}
	}
	return false
}

func (i Interface) inherits(ext Interface) bool {
	return false
}

func (c Class) inherits(ext Class) bool {
	return false
}

func (c Class) hasProperty(f Func) bool {
	return false
}
