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
		Filename: l.fileName,
		Offset:   l.byteOffset,
		Line:     l.line,
	}
}

func (l *Lexer) current() byte {
	return l.Bytes()[l.Offset()]
}

func (l *Lexer) isWhitespace() bool {
	return (l.current() == ' ') || (l.current() == '\t') || (l.current() == '\v') || (l.current() == '\f')
}

func (l *Lexer) next() {
	if l.byteOffset >= uint(len(l.buffer)) {
		return
	}
	for l.isWhitespace() {
		l.byteOffset++
		if l.byteOffset >= uint(len(l.buffer)) {
			return
		}
	}
	pt := token.NextProtoToken(l)
	if pt != nil {
		t := pt.Process(l)
		t.Proto = pt
		if pt.Type == token.None {
			l.byteOffset++
		} else {
			l.Tokens = append(l.Tokens, t)
		}
		l.line += uint(t.LineIncrement)
	} else {
		l.addError(l.getCurrentLocation(), "Unrecognised token")
		l.byteOffset++
	}
	l.next()
}

func (l *Lexer) addError(loc util.Location, err string, data ...interface{}) {
	if l.Errors == nil {
		l.Errors = make([]util.Error, 0)
	}
	l.Errors = append(l.Errors, util.Error{
		Location: loc,
		Message:  fmt.Sprintf(err, data...),
	})
}
