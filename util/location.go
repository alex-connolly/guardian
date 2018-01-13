package token

// modified version of token.Pos in golang --> don't reinvent the wheel
// not going to add the full thing for now
type Location struct {
	Filename string // filename, if any
	Offset   int    // offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (byte count)
}
