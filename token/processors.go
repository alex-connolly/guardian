package token

func processNewLine(b Byterable) Token {
	//b.line++
	next(b)
	return Token{
		Type: NewLine,
	}
}

func processIgnored(b Byterable) Token {
	return Token{
		Type: None,
	}
}

func processInteger(b Byterable) (t Token) {
	t.Start = b.Offset()
	t.End = b.Offset()
	t.Type = Integer
	if current(b) == '0' {
		t.End++
		next(b)
		if isEnd(b) {
			return t
		}
		if current(b) == 'x' || current(b) == 'X' {
			//hexadecimal
			t.End++
			next(b)
			for '0' <= current(b) && current(b) <= '9' {
				next(b)
				t.End++
				if isEnd(b) {
					return t
				}
			}
		}
	} else {
		for '0' <= current(b) && current(b) <= '9' {
			next(b)
			t.End++
			if isEnd(b) {
				return t
			}
		}
	}
	return t
}

func processFloat(b Byterable) (t Token) {
	// TODO: make this handle exponents
	t.Start = b.Offset()
	t.End = b.Offset()
	t.Type = Float
	decimalUsed := false
	for '0' <= current(b) && current(b) <= '9' || current(b) == '.' {
		if current(b) == '.' {
			if decimalUsed {
				return t
			}
			decimalUsed = true
		}
		next(b)
		t.End++
		if isEnd(b) {
			return t
		}
	}
	return t
}

// TODO: handle errors etc
func processCharacter(b Byterable) Token {
	t := new(Token)
	t.Start = b.Offset()
	t.End = b.Offset()
	t.Type = Character
	b1 := next(b)
	b2 := next(b)
	for b1 != b2 {
		t.End++
		b2 = next(b)
		if isEnd(b) {
			//b.error("Character literal not closed")
			t.End += 2
			return *t
		}
	}
	t.End += 2
	return *t
}

func processIdentifier(b Byterable) Token {

	t := new(Token)
	t.Start = b.Offset()
	t.End = b.Offset()
	t.Type = Identifier
	/*if isEnd(b) {
		return *t
	}*/
	for isIdentifier(b) {
		//fmt.Printf("id: %c\n", b.Bytes()[b.Offset()])
		t.End++
		next(b)
		if isEnd(b) {
			return *t
		}
	}
	return *t
}

// processes a string sequence to create a new Token.
func processString(b Byterable) Token {
	// the Start - End is the value
	// it DOES include the enclosing quotation marks
	t := new(Token)
	t.Start = b.Offset()
	t.End = b.Offset()
	t.Type = String
	b1 := next(b)
	b2 := next(b)
	for b1 != b2 {
		t.End++
		b2 = next(b)
		if isEnd(b) {
			//b.error("String literal not closed")
			t.End += 2
			return *t
		}
	}
	t.End += 2
	return *t
}

func processFixed(len int, tkn Type) processorFunc {
	return func(b Byterable) (t Token) {
		// Start and End don't matter
		t.Type = tkn
		b.SetOffset(b.Offset() + len)
		return t
	}
}
