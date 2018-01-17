package typing

func (a *Array) Compare(t Type) bool {
	if at, ok := ResolveUnderlying(t).(*Array); !ok {
		return false
	} else {
		// arrays are equal if they share the same value type
		if a.Variable && !at.Variable {
			return false
		}
		if !a.Variable && at.Variable {
			return false
		}
		if a.Length != at.Length {
			return false
		}
		return a.Value.Compare(at.Value)
	}
}

func (m *Map) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Map); !ok {
		return false
	} else {
		// map Types are equal if they share the same key and value
		return m.Key.Compare(other.Key) && m.Value.Compare(other.Value)
	}
}

func (t *Tuple) Compare(o Type) bool {
	if o == nil {
		return false
	}
	if other, ok := ResolveUnderlying(o).(*Tuple); !ok {
		return false
	} else {
		// short circuit if not the same length
		if other.Types == nil && t.Types != nil {
			return false
		}
		if len(t.Types) != len(other.Types) {
			return false
		}
		for i, typ := range t.Types {
			if typ != nil {
				if !AssignableTo(typ, other.Types[i], true) {
					return false
				}
			} else {
				// if type is nil, return false
				return false
			}
		}
		return true
	}

}

func (f *Func) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Func); !ok {
		return false
	} else {
		// func Types are equal if they share the same params and Results
		return f.Params.Compare(other.Params) && f.Results.Compare(other.Results)
	}
}

func (a *Aliased) Compare(t Type) bool {
	return ResolveUnderlying(a).Compare(t)
}

func (s *StandardType) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*StandardType); !ok {
		return false
	} else {
		return s == other
	}
}

func (c *Class) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Class); !ok {
		return false
	} else {
		return c.Name == other.Name
	}
}

func (i *Interface) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Interface); !ok {
		return false
	} else {
		return i.Name == other.Name
	}
}

func (e *Enum) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Enum); !ok {
		return false
	} else {
		return e.Name == other.Name
	}
}

func (c *Contract) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Contract); !ok {
		return false
	} else {
		return c.Name == other.Name
	}
}

func (e *Event) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Event); !ok {
		return false
	} else {
		return e.Name == other.Name
	}
}

func (p *Package) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*Package); !ok {
		return false
	} else {
		return p.Name == other.Name
	}
}

func (nt *NumericType) Compare(t Type) bool {
	if other, ok := ResolveUnderlying(t).(*NumericType); !ok {
		return false
	} else {
		if nt.Integer != other.Integer {
			return false
		}
		if nt.BitSize != other.BitSize {
			return false
		}
		if nt.Signed != other.Signed {
			return false
		}
		return true
	}
}

func (nt *BooleanType) Compare(t Type) bool {
	_, ok := ResolveUnderlying(t).(*BooleanType)
	return ok
}

func (g *Generic) Compare(t Type) bool {
	other, ok := ResolveUnderlying(t).(*Generic)
	if !ok {
		return false
	}
	for _, i := range g.Inherits {
		if !other.inherits(i) {
			return false
		}
	}
	for _, i := range g.Interfaces {
		if !other.implements(i) {
			return false
		}
	}
	return true
}
