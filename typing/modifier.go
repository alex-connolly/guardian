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
func (a *Aliased) Modifiers() *Modifiers      { return a.Mods }
func (t *Tuple) Modifiers() *Modifiers        { return nil }
func (e *Event) Modifiers() *Modifiers        { return e.Mods }

func (g *Generic) ResetModifiers()      { g.Mods = nil }
func (a *Array) ResetModifiers()        { a.Mods = nil }
func (m *Map) ResetModifiers()          { m.Mods = nil }
func (c *Class) ResetModifiers()        { c.Mods = nil }
func (e *Enum) ResetModifiers()         { e.Mods = nil }
func (i *Interface) ResetModifiers()    { i.Mods = nil }
func (c *Contract) ResetModifiers()     { c.Mods = nil }
func (c *NumericType) ResetModifiers()  { c.Mods = nil }
func (c *BooleanType) ResetModifiers()  { c.Mods = nil }
func (c *VoidType) ResetModifiers()     { c.Mods = nil }
func (c *Func) ResetModifiers()         { c.Mods = nil }
func (a *StandardType) ResetModifiers() {}
func (a *Aliased) ResetModifiers()      { a.Mods = nil }
func (t *Tuple) ResetModifiers()        {}
func (e *Event) ResetModifiers()        { e.Mods = nil }
