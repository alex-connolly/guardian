package lexer

type lexer struct {
	buffer []byte
	offset int
	line   int
	column int
	Tokens []Token
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
			if t.Type != TknNone {
				l.Tokens = append(l.Tokens, t)
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
		l.errors = append(l.errors, "Unrecognised Token.")
		l.offset++
	}
	l.next()
}

func (l *lexer) isEOF() bool {
	return l.offset >= len(l.buffer)
}

// creates a new string from the Token's value
// TODO: escaped characters
func (l *lexer) TokenString(t Token) string {
	data := make([]byte, t.end-t.start)
	copy(data, l.buffer[t.start:t.end])
	return string(data)
}

func (l *lexer) nextByte() byte {
	b := l.buffer[l.offset]
	l.offset++
	return b
}

func LexString(str string) ([]Token, []string) {
	return Lex([]byte(str))
}

func (l *lexer) current() byte {
	return l.buffer[l.offset]
}

// Lex ...
func Lex(bytes []byte) ([]Token, []string) {
	l := new(lexer)
	l.buffer = bytes
	l.next()
	return l.Tokens, l.errors
}

func processNewLine(l *lexer) Token {
	l.line++
	return Token{
		Type: TknNone,
	}
}

func processIgnored(l *lexer) Token {
	return Token{
		Type: TknNone,
	}
}

func processNumber(l *lexer) (t Token) {
	t.start = l.offset
	t.end = l.offset
	t.Type = TknNumber
	for '0' <= l.buffer[l.offset] && l.buffer[l.offset] <= '9' {
		l.offset++
		t.end++
		if l.isEOF() {
			return t
		}
	}
	return t
}

// TODO: handle errors etc
func processCharacter(l *lexer) (t Token) {
	t.start = l.offset
	t.end = l.offset + 2
	t.Type = TknCharacter
	return t
}

func processIdentifier(l *lexer) Token {

	t := new(Token)
	t.start = l.offset
	t.end = l.offset
	t.Type = TknIdentifier
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

// processes a string sequence to create a new Token.
func processString(l *lexer) Token {
	// the start - end is the value
	// it DOES include the enclosing quotation marks
	t := new(Token)
	t.start = l.offset
	t.end = l.offset
	t.Type = TknString
	b := l.nextByte()
	b2 := l.nextByte()
	for b != b2 {
		t.end++
		b2 = l.nextByte()
		if l.isEOF() {
			l.errors = append(l.errors, "String literal not closed")
			t.end += 2
			return *t
		}
	}
	t.end += 2
	return *t
}
