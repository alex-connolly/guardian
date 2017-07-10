package lexer

type lexer struct {
	buffer []byte
	offset int
	line   int
	column int
	tokens []token
	errors []string
}

func (l *lexer) next() {
	if l.isEOF() {
		return
	}
	found := false
	for _, pt := range getProtoTokens() {
		if pt.identifier(l) {
			t := pt.process(l)
			if t.tkntype != TknNone {
				l.tokens = append(l.tokens, t)
			} else {
				l.offset++
			}
			found = true
			break
		}
	}
	if !found {
		if l.errors == nil {
			l.errors = make([]string, 0)
		}
		l.errors = append(l.errors, "Unrecognised token.")
		l.offset++
	}
	l.next()
}

func (l *lexer) isEOF() bool {
	return l.offset >= len(l.buffer)
}

// creates a new string from the token's value
// TODO: escaped characters
func (l *lexer) tokenString(t token) string {
	data := make([]byte, t.end-t.start)
	copy(data, l.buffer[t.start:t.end])
	return string(data)
}

func (l *lexer) nextByte() byte {
	b := l.buffer[l.offset]
	l.offset++
	return b
}

func lexString(str string) *lexer {
	return lex([]byte(str))
}

func (l *lexer) current() byte {
	return l.buffer[l.offset]
}

func lex(bytes []byte) *lexer {
	l := new(lexer)
	l.buffer = bytes
	l.next()
	return l
}

func processNewLine(l *lexer) token {
	l.line++
	return token{
		tkntype: TknNone,
	}
}

func processIgnored(l *lexer) token {
	return token{
		tkntype: TknNone,
	}
}

func processNumber(l *lexer) (t token) {
	t.start = l.offset
	t.end = l.offset
	t.tkntype = TknNumber
	for '0' <= l.buffer[l.offset] && l.buffer[l.offset] <= '9' {
		l.offset++
		t.end++
		if l.isEOF() {
			return t
		}
	}
	return t
}

func processCharacter(l *lexer) (t token) {
	t.start = l.offset
	t.end = l.offset + 2
	t.tkntype = TknCharacter
	return t
}

func processIdentifier(l *lexer) token {

	t := new(token)
	t.start = l.offset
	t.end = l.offset
	t.tkntype = TknValue
	if l.isEOF() {
		return *t
	}
	for isIdentifier(l) {
		//fmt.Printf("id: %c\n", l.buffer[l.offset])
		t.end++
		l.offset++
		if l.isEOF() {
			return *t
		}
	}
	return *t
}

// processes a string sequence to create a new token.
func processString(l *lexer) token {
	// the start - end is the value
	// it DOES include the enclosing quotation marks
	t := new(token)
	t.start = l.offset
	t.end = l.offset
	t.tkntype = TknString
	b := l.nextByte()
	b2 := l.nextByte()
	for b != b2 {
		t.end++
		b2 = l.nextByte()
		if l.isEOF() {
			t.end += 2
			return *t
		}
	}
	t.end += 2
	return *t
}
