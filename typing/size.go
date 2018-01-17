package typing

// All Size values should be in bits
const (
	byteSize = 8
)

func (g *Generic) Size() uint {
	return 0
}

func (a *Array) Size() uint {
	return uint(a.Length) * a.Value.Size()
}

func (m *Map) Size() uint {
	return 0
}

func (p *Package) Size() uint {
	return 0
}

func (c *Class) Size() uint {
	s := uint(0)
	for _, v := range c.Properties {
		s += v.Size()
	}
	return s
}

func (i *Interface) Size() uint {
	return 0
}

func (t *Tuple) Size() uint {
	s := uint(0)
	for _, typ := range t.Types {
		s += typ.Size()
	}
	return s
}

func (nt *NumericType) Size() uint {
	return uint(nt.BitSize)
}

func (bt *BooleanType) Size() uint {
	return 8
}

func (c *Contract) Size() uint {
	return 0
}

func (e *Enum) Size() uint {
	return 0
}

func (s *StandardType) Size() uint {
	return 0
}

func (f *Func) Size() uint {
	return 0
}

func (a *Aliased) Size() uint {
	return ResolveUnderlying(a).Size()
}

func (e *Event) Size() uint {
	return 0
}

func (v *VoidType) Size() uint {
	return 0
}
