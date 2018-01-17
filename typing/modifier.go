package typing

// shoutout to maz - let me have her seat on the bus to do this <3

// why do types have modifiers/annotations?
// so that they aren't lost when the declared variable is referenced
/*
   @Builtin("dog") var dog string

   x = dog // we need to know to replace this reference to dog

*/

func HasModifier(t Type, mod string) bool {
	if t.Modifiers() == nil {
		return false
	}
	return t.Modifiers().HasModifier(mod)
}

func AddModifier(t Type, mod string) {
	if t.Modifiers() == nil {
		t.SetModifiers(new(Modifiers))
	}
	mods := t.Modifiers()
	mods.AddModifier(mod)
}

func (g *Generic) Modifiers() *Modifiers      { return g.Mods }
func (a *Array) Modifiers() *Modifiers        { return a.Mods }
func (m *Map) Modifiers() *Modifiers          { return m.Mods }
func (c *Class) Modifiers() *Modifiers        { return c.Mods }
func (e *Enum) Modifiers() *Modifiers         { return e.Mods }
func (i *Interface) Modifiers() *Modifiers    { return i.Mods }
func (c *Contract) Modifiers() *Modifiers     { return c.Mods }
func (c *NumericType) Modifiers() *Modifiers  { return c.Mods }
func (c *BooleanType) Modifiers() *Modifiers  { return c.Mods }
func (c *VoidType) Modifiers() *Modifiers     { return c.Mods }
func (c *Func) Modifiers() *Modifiers         { return c.Mods }
func (a *StandardType) Modifiers() *Modifiers { return a.Mods }
func (a *Aliased) Modifiers() *Modifiers      { return a.Mods }
func (t *Tuple) Modifiers() *Modifiers        { return t.Mods }
func (e *Event) Modifiers() *Modifiers        { return e.Mods }
func (p *Package) Modifiers() *Modifiers      { return p.Mods }

func (g *Generic) SetModifiers(m *Modifiers)      { g.Mods = m }
func (a *Array) SetModifiers(m *Modifiers)        { a.Mods = m }
func (m *Map) SetModifiers(a *Modifiers)          { m.Mods = a }
func (c *Class) SetModifiers(m *Modifiers)        { c.Mods = m }
func (e *Enum) SetModifiers(m *Modifiers)         { e.Mods = m }
func (i *Interface) SetModifiers(m *Modifiers)    { i.Mods = m }
func (c *Contract) SetModifiers(m *Modifiers)     { c.Mods = m }
func (c *NumericType) SetModifiers(m *Modifiers)  { c.Mods = m }
func (c *BooleanType) SetModifiers(m *Modifiers)  { c.Mods = m }
func (c *VoidType) SetModifiers(m *Modifiers)     { c.Mods = m }
func (c *Func) SetModifiers(m *Modifiers)         { c.Mods = m }
func (a *StandardType) SetModifiers(m *Modifiers) { a.Mods = m }
func (a *Aliased) SetModifiers(m *Modifiers)      { a.Mods = m }
func (t *Tuple) SetModifiers(m *Modifiers)        { t.Mods = m }
func (e *Event) SetModifiers(m *Modifiers)        { e.Mods = m }
func (p *Package) SetModifiers(m *Modifiers)      { p.Mods = m }
