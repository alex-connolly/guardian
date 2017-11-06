package validator

func (a Array) compare(t Type) bool {
	if at, ok := resolveUnderlying(t).(Array); !ok {
		return false
	} else {
		// arrays are equal if they share the same value type
		return a.Value.compare(at.Value)
	}
}

func (m Map) compare(t Type) bool {
	if other, ok := resolveUnderlying(t).(Map); !ok {
		return false
	} else {
		// map types are equal if they share the same key and value
		return m.Key.compare(other.Key) && m.Value.compare(other.Value)
	}
}

func (t Tuple) compare(o Type) bool {
	if o == nil {
		return false
	}
	if other, ok := resolveUnderlying(o).(Tuple); !ok {
		return false
	} else {
		// short circuit if not the same length
		if other.types == nil && t.types != nil {
			return false
		}
		if len(t.types) != len(other.types) {
			return false
		}
		for i, typ := range t.types {
			if typ != nil {
				if !typ.compare(other.types[i]) && typ != standards[Unknown] {
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

func (f Func) compare(t Type) bool {
	if other, ok := resolveUnderlying(t).(Func); !ok {
		return false
	} else {
		// func types are equal if they share the same params and results
		return f.Params.compare(other.Params) && f.Results.compare(other.Results)
	}
}

func (a Aliased) compare(t Type) bool {
	return resolveUnderlying(a).compare(t)
}

func (s StandardType) compare(t Type) bool {
	if other, ok := resolveUnderlying(t).(StandardType); !ok {
		return false
	} else {
		return s == other
	}
}

func (c Class) compare(t Type) bool {
	if other, ok := resolveUnderlying(t).(Class); !ok {
		return false
	} else {
		return c.Name == other.Name
	}
}

func (i Interface) compare(t Type) bool {
	if other, ok := resolveUnderlying(t).(Interface); !ok {
		return false
	} else {
		return i.Name == other.Name
	}
}

func (e Enum) compare(t Type) bool {
	if other, ok := resolveUnderlying(t).(Enum); !ok {
		return false
	} else {
		return e.Name == other.Name
	}
}

func (c Contract) compare(t Type) bool {
	if other, ok := resolveUnderlying(t).(Contract); !ok {
		return false
	} else {
		return c.Name == other.Name
	}
}

func (e Event) compare(t Type) bool {
	if other, ok := resolveUnderlying(t).(Event); !ok {
		return false
	} else {
		return e.Name == other.Name
	}
}
