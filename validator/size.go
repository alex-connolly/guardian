package validator

func (a Array) Size() int {
	return a.Length * a.Value.Size()
}

func (m Map) Size() int {
	return 0
}

func (c Class) Size() int {
	s := 0
	for _, f := range c.Properties {
		s += f.Size()
	}
	return s
}

func (i Interface) Size() int {
	return 0
}

func (t Tuple) Size() int {
	s := 0
	for _, typ := range t.types {
		s += typ.Size()
	}
	return s
}

func (nt NumericType) Size() int {
	return nt.BitSize
}

func (bt BooleanType) Size() int {
	return 8
}

func (c Contract) Size() int {
	return 0
}

func (e Enum) Size() int {
	return 0
}

func (s StandardType) Size() int {
	return 0
}

func (f Func) Size() int {
	return 0
}

func (a Aliased) Size() int {
	return resolveUnderlying(a).Size()
}

func (e Event) Size() int {
	return 0
}
