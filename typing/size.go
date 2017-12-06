package typing

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
	for _, typ := range t.Types {
		s += typ.size()
	}
	return s
}

func (nt NumericType) size() int {
	return nt.BitSize
}

func (bt booleaneanType) size() int {
	return 8
}

func (c Contract) size() int {
	return 0
}

func (e Enum) size() int {
	return 0
}

func (s standardType) size() int {
	return 0
}

func (f Func) size() int {
	return 0
}

func (a Aliased) size() int {
	return ResolveUnderlying(a).size()
}

func (e Event) size() int {
	return 0
}
