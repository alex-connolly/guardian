package lexer

import (
	"fmt"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/util"
)

func (l *Lexer) Bytes() []byte {
	return l.buffer
}

func (l *Lexer) Offset() uint {
	return l.byteOffset
}

func (l *Lexer) SetOffset(o uint) {
	l.byteOffset = o
}

func (l *Lexer) Location() util.Location {
	return l.getCurrentLocation()
}

func (l *Lexer) getCurrentLocation() util.Location {
	return util.Location{
		Filename: "default.grd",
		Offset:   l.byteOffset,
		Line:     l.line,
	}
}

func (l *Lexer) finalise(t token.Token) {
	fmt.Printf("start: %d, end: %d\n", t.Start.Offset, t.End.Offset)
	t.Data = make([]byte, t.End.Offset-t.Start.Offset)
	copy(t.Data, l.buffer[t.Start.Offset:t.End.Offset])
}

func (l *Lexer) next() {
	if l.byteOffset == uint(len(l.buffer)) {
		return
	}
	pt := token.NextProtoToken(l)
	if pt != nil {
		t := pt.Process(l)
		t.Proto = pt
		if pt.Type == token.None {
			l.byteOffset++
		} else {
			l.finalise(t)
			l.tokens = append(l.tokens, t)
		}
	} else {
		l.addError(l.getCurrentLocation(), "Unrecognised token")
		l.byteOffset++
	}
	l.next()
}

func (l *Lexer) addError(loc util.Location, err string, data ...interface{}) {
	if l.errors == nil {
		l.errors = make([]util.Error, 0)
	}
	l.errors = append(l.errors, util.Error{
		Location: loc,
		Message:  fmt.Sprintf(err, data...),
	})
}
