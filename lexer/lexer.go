package lexer

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/util"
)

func (l *Lexer) Bytes() []byte {
	return l.buffer
}

func (l *Lexer) Offset() int {
	return l.byteOffset
}

func (l *Lexer) SetOffset(o int) {
	l.byteOffset = o
}

func (l *Lexer) next() {
	if l.byteOffset == len(l.buffer) {
		return
	}
	pt := token.NextProtoToken(l)
	if pt != nil {
		t := pt.Process(l)
		if pt.Type == token.None {
			l.byteOffset++
		} else {
			t.Finalise(l)
			l.tokens = append(l.tokens, t)
		}
	} else {
		l.error("Unrecognised token.Token.")
		l.byteOffset++
	}
	l.next()
}

func (l *Lexer) error(msg string) {
	if l.errors == nil {
		l.errors = make([]util.Error, 0)
	}
	l.errors = append(l.errors, util.Error{
		LineNumber: l.line,
		Message:    msg,
	})
}
