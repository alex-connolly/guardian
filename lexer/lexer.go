package lexer

import (
	"github.com/end-r/guardian/util"
)

func (l *Lexer) next() {
	if l.isEOF() {
		return
	}
	found := false
	for _, pt := range getProtoTokens() {
		if pt.identifier(l) {
			t := pt.process(l)
			t.proto = pt
			if t.Type != TknNone {
				//log.Printf("Found tok type: %d", t.Type)
				l.tokens = append(l.tokens, l.finalise(t))
			} else {
				l.byteOffset++
			}
			found = true
			break
		}
	}
	if !found {
		l.error("Unrecognised Token.")
		l.byteOffset++
	}
	l.next()
}

func (l *Lexer) finalise(t Token) Token {
	t.data = make([]byte, t.end-t.start)
	copy(t.data, l.buffer[t.start:t.end])
	return t
}

func (l *Lexer) isEOF() bool {
	return l.byteOffset >= len(l.buffer)
}

func (l *Lexer) nextByte() byte {
	b := l.buffer[l.byteOffset]
	l.byteOffset++
	return b
}

func (l *Lexer) current() byte {
	return l.buffer[l.byteOffset]
}

func (l *Lexer) hasBytes(offset int) bool {
	return l.byteOffset+offset <= len(l.buffer)
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
