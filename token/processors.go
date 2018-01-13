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
	t.Start = b.Location(b.Offset())
	end := b.Offset()
	t.Type = Integer
	if current(b) == '-' {
		end++
		next(b)
	}
	if current(b) == '0' {
		end++
		next(b)
		if isEnd(b) {
			return t
		}
		if current(b) == 'x' || current(b) == 'X' {
			//hexadecimal
			end++
			next(b)
			for '0' <= current(b) && current(b) <= 'F' {
				next(b)
				end++
				if isEnd(b) {
					return t
				}
			}
		}
	} else {
		for '0' <= current(b) && current(b) <= '9' {
			next(b)
			end++
			if isEnd(b) {
				return t
			}
		}
	}
	return t
}

func processFloat(b Byterable) (t Token) {
	// TODO: make this handle exponents
	t.Start = b.Location(b.Offset())
	end := b.Offset()
	t.Type = Float
	decimalUsed := false
	if current(b) == '-' {
		t.End = b.Location(end)
		next(b)
	}
	for '0' <= current(b) && current(b) <= '9' || current(b) == '.' {
		if current(b) == '.' {
			if decimalUsed {
				t.End = b.Location(end)
				return t
			}
			decimalUsed = true
		}
		next(b)
		end++
		if isEnd(b) {
			t.End = b.Location(end)
			return t
		}
	}
	t.End = b.Location(end)
	return t
}

// TODO: handle errors etc
func processCharacter(b Byterable) Token {
	t := new(Token)
	t.Start = b.Location(b.Offset())
	end := b.Offset()
	t.Type = Character
	b1 := next(b)
	b2 := next(b)
	for b1 != b2 {
		end++
		b2 = next(b)
		if isEnd(b) {
			//b.error("Character literal not closed")
			end += 2
			return *t
		}
	}
	end += 2
	return *t
}

func processIdentifier(b Byterable) Token {

	t := new(Token)
	t.Start = b.Location(b.Offset())
	end := b.Offset()
	t.Type = Identifier
	/*if isEnd(b) {
		return *t
	}*/
	for isIdentifier(b) {
		//fmt.Printf("id: %c\n", b.Bytes()[b.Offset()])
		end++
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
	t.Start = b.Location(b.Offset())
	end := b.Offset()
	t.Type = String
	b1 := next(b)
	b2 := next(b)
	for b1 != b2 {
		end++
		b2 = next(b)
		if isEnd(b) {
			//b.error("String literal not closed")
			end += 2
			return *t
		}
	}
	end += 2
	return *t
}

func processFixed(len uint, tkn Type) processorFunc {
	return func(b Byterable) (t Token) {
		// Start and End don't matter
		t.Type = tkn
		b.SetOffset(b.Offset() + len)
		return t
	}
}
