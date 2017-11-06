package validator

import (
	"bytes"
	"fmt"
)

func WriteType(t Type) string {
	b := new(bytes.Buffer)
	t.write(b)
	return b.String()
}

func (s StandardType) write(b *bytes.Buffer) {

	b.WriteString(s.name)
}

func (t Tuple) write(b *bytes.Buffer) {
	b.WriteByte('(')
	if t.types != nil {
		for i, v := range t.types {
			if i > 0 {
				b.WriteString(", ")
			}
			if v == nil {
				fmt.Println("type should not be nil")
			} else {
				v.write(b)
			}
		}
	}
	b.WriteByte(')')
}

func (a Aliased) write(b *bytes.Buffer) {
	b.WriteString(a.alias)
}

func (a Array) write(b *bytes.Buffer) {
	b.WriteString("[]")
	a.Value.write(b)
}

func (m Map) write(b *bytes.Buffer) {
	b.WriteString("map[")
	m.Key.write(b)
	b.WriteByte(']')
	m.Value.write(b)
}

func (f Func) write(b *bytes.Buffer) {
	b.WriteString("func")
	f.Params.write(b)
	f.Results.write(b)
}

func (c Class) write(b *bytes.Buffer) {
	b.WriteString("class")
	b.WriteString(c.Name)
}

func (i Interface) write(b *bytes.Buffer) {
	b.WriteString("interface")
	b.WriteString(i.Name)
}

func (e Enum) write(b *bytes.Buffer) {
	b.WriteString("enum")
	b.WriteString(e.Name)
}

func (c Contract) write(b *bytes.Buffer) {
	b.WriteString("contract")
	b.WriteString(c.Name)
}

func (e Event) write(b *bytes.Buffer) {
	b.WriteString("event")
	e.Parameters.write(b)
}
