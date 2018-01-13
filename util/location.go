package util

// modified version of token.Pos in golang --> don't reinvent the wheel
// not going to add the full thing for now
type Location struct {
	Filename string // filename, if any
	Offset   uint   // offset, starting at 0
	Line     uint   // line number, starting at 1
	Column   uint   // column number, starting at 1 (byte count)
}
