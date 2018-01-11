package typing

// Generic ...
type Generic struct {
	Mods       *Modifiers
	Identifier string
	Interfaces []*Interface
	Inherits   []Type
}

// Accepts ...
func (g *Generic) Accepts(t Type) bool {
	for _, super := range g.Inherits {
		if !t.inherits(super) {
			return false
		}
	}
	for _, ifc := range g.Interfaces {
		if !t.implements(ifc) {
			return false
		}
	}
	return true
}
