package typing

import (
	"bytes"
	"fmt"
)

// WriteType creates a string representation of a Guardian type
func WriteType(t Type) string {
	b := new(bytes.Buffer)
	t.write(b)
	return b.String()
}

func (g *Generic) write(b *bytes.Buffer) {
	b.WriteString(g.Identifier)
}

func (p *Package) write(b *bytes.Buffer) {
	b.WriteString(p.Name)
}

func (s *StandardType) write(b *bytes.Buffer) {
	b.WriteString(s.name)
}

func (t *Tuple) write(b *bytes.Buffer) {
	b.WriteByte('(')
	if t.Types != nil {
		for i, v := range t.Types {
			if i > 0 {
				b.WriteString(", ")
			}
			if v == nil {
				b.WriteString("INVALID NIL TYPE")
			} else {
				v.write(b)
			}
		}
	}
	b.WriteByte(')')
}

func (a *Aliased) write(b *bytes.Buffer) {
	b.WriteString(a.Alias)
}

func (a *Array) write(b *bytes.Buffer) {
	b.WriteString("[")
	if !a.Variable {
		b.WriteString(fmt.Sprintf("%d", a.Length))
	}
	b.WriteString("]")
	a.Value.write(b)
}

func (m *Map) write(b *bytes.Buffer) {
	b.WriteString("map[")
	m.Key.write(b)
	b.WriteByte(']')
	m.Value.write(b)
}

func (f *Func) write(b *bytes.Buffer) {
	b.WriteString("func ")
	b.WriteString(f.Name)
	if f.Generics != nil && len(f.Generics) > 0 {
		b.WriteString("<")
		for _, g := range f.Generics {
			b.WriteString(g.Identifier)
		}
		b.WriteString(">")
	}
	f.Params.write(b)
	f.Results.write(b)
}

func (c *Class) write(b *bytes.Buffer) {
	b.WriteString(c.Name)
	if c.Generics != nil && len(c.Generics) > 0 {
		b.WriteString("<")
		for _, g := range c.Generics {
			b.WriteString(g.Identifier)
		}
		b.WriteString(">")
	}
}

func (i *Interface) write(b *bytes.Buffer) {
	b.WriteString(i.Name)
}

func (e *Enum) write(b *bytes.Buffer) {
	b.WriteString(e.Name)
}

func (c *Contract) write(b *bytes.Buffer) {
	b.WriteString(c.Name)
	if c.Generics != nil && len(c.Generics) > 0 {
		b.WriteString("<")
		for _, g := range c.Generics {
			b.WriteString(g.Identifier)
		}
		b.WriteString(">")
	}
}

func (e *Event) write(b *bytes.Buffer) {
	b.WriteString("event")
	e.Parameters.write(b)
}

func (nt *NumericType) write(b *bytes.Buffer) {
	b.WriteString(nt.Name)
}

func (nt *BooleanType) write(b *bytes.Buffer) {
	b.WriteString("bool")
}

func (v *VoidType) write(b *bytes.Buffer) {
	b.WriteString("void")
}
