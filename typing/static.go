package typing

func (g *Generic) MakeStatic()      { g.static = true }
func (a *Array) MakeStatic()        { a.static = true }
func (m *Map) MakeStatic()          { m.static = true }
func (c *Class) MakeStatic()        { c.static = true }
func (e *Enum) MakeStatic()         { e.static = true }
func (i *Interface) MakeStatic()    { i.static = true }
func (c *Contract) MakeStatic()     { c.static = true }
func (c *NumericType) MakeStatic()  { c.static = true }
func (c *BooleanType) MakeStatic()  { c.static = true }
func (c *VoidType) MakeStatic()     { c.static = true }
func (c *Func) MakeStatic()         { c.static = true }
func (a *StandardType) MakeStatic() { a.static = true }
func (a *Aliased) MakeStatic()      { a.static = true }
func (t *Tuple) MakeStatic()        { t.static = true }
func (e *Event) MakeStatic()        { e.static = true }

func (g *Generic) Static() bool      { return false }
func (a *Array) Static() bool        { return a.static }
func (m *Map) Static() bool          { return m.static }
func (c *Class) Static() bool        { return c.static }
func (e *Enum) Static() bool         { return e.static }
func (i *Interface) Static() bool    { return i.static }
func (c *Contract) Static() bool     { return c.static }
func (c *NumericType) Static() bool  { return c.static }
func (c *BooleanType) Static() bool  { return c.static }
func (c *VoidType) Static() bool     { return c.static }
func (c *Func) Static() bool         { return c.static }
func (a *StandardType) Static() bool { return a.static }
func (a *Aliased) Static() bool      { return a.static }
func (t *Tuple) Static() bool        { return t.static }
func (e *Event) Static() bool        { return e.static }
