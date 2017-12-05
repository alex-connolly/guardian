package validator

func (a Array) size() int {
	return a.Length * a.Value.size()
}

func (m Map) size() int {
	return 0
}

func (c Class) size() int {
	s := 0
	for _, f := range c.Properties {
		s += f.size()
	}
	return s
}

func (i Interface) size() int {
	return 0
}

func (t Tuple) size() int {
	s := 0
	for _, typ := range t.types {
		s += typ.size()
	}
	return s
}

func (nt NumericType) size() int {
	return nt.BitSize
}

func (bt BooleanType) size() int {
	return 8
}

func (c Contract) size() int {
	return 0
}

func (e Enum) size() int {
	return 0
}

func (s StandardType) size() int {
	return 0
}

func (f Func) size() int {
	return 0
}

func (a Aliased) size() int {
	return resolveUnderlying(a).size()
}

func (e Event) size() int {
	return 0
}
