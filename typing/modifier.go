package typing

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
func (a *StandardType) Modifiers() *Modifiers { return nil }
