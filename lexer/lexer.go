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
	found := false
	for _, pt := range token.GetProtoTokens() {
		if pt.Identifier(l) {
			t := pt.Process(l)
			t.Proto = pt
			if t.Type != token.None {
				//log.Printf("Found tok type: %d", t.Type)
				t.Finalise(l)
				l.tokens = append(l.tokens, t)
			} else {
				l.byteOffset++
			}
			found = true
			break
		}
	}
	if !found {
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
