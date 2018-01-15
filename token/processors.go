package token

func processNewLine(b Byterable) Token {
	//b.line++
	next(b)
	return Token{
		Type: NewLine,
	}
}

func processIgnored(b Byterable) Token {
	next(b)
	return Token{
		Type: None,
	}
}

func markLimits(b Byterable, f func(Byterable) Token) Token {
	start := b.Location()
	tok := f(b)
	tok.Start = start
	tok.End = b.Location()
	return tok
}

func processInteger(b Byterable) (t Token) {
	return markLimits(b, func(byt Byterable) (t Token) {
		t.Type = Integer
		if current(b) == '-' {
			next(b)
		}
		if current(b) == '0' {
			next(b)
			if isEnd(b) {
				return t
			}
			if current(b) == 'x' || current(b) == 'X' {
				//hexadecimal
				next(b)
				for '0' <= current(b) && current(b) <= 'F' {
					next(b)
					if isEnd(b) {
						t.End = b.Location()
						return t
					}
				}
			}
		} else {
			for '0' <= current(b) && current(b) <= '9' {
				next(b)
				if isEnd(b) {
					t.End = b.Location()
					return t
				}
			}
		}
		return t
	})
}

func processFloat(b Byterable) (t Token) {
	// TODO: make this handle exponents
	return markLimits(b, func(byt Byterable) (t Token) {
		t.Type = Float
		decimalUsed := false
		if current(byt) == '-' {
			next(byt)
		}
		for '0' <= current(byt) && current(byt) <= '9' || current(byt) == '.' {
			if current(byt) == '.' {
				if decimalUsed {
					return t
				}
				decimalUsed = true
			}
			next(byt)
			if isEnd(byt) {
				return t
			}
		}
		return t
	})

}

// TODO: handle errors etc
func processCharacter(b Byterable) Token {
	return markLimits(b, func(byt Byterable) (t Token) {
		t.Type = Character
		b1 := next(byt)
		b2 := next(byt)
		for b1 != b2 {
			if isEnd(byt) {
				//b.error("Character literal not closed")
				//next(byt)
				return t
			}
			b2 = next(byt)
		}
		return t
	})

}

func processIdentifier(b Byterable) Token {

	return markLimits(b, func(byt Byterable) (t Token) {
		t.Start = byt.Location()
		t.Type = Identifier
		for isIdentifier(byt) {
			if isEnd(byt) {
				return t
			}
			next(byt)

		}
		return t
	})

}

// processes a string sequence to create a new Token.
func processString(b Byterable) Token {
	return markLimits(b, func(byt Byterable) (t Token) {
		// the Start - End is the value
		// it DOES include the enclosing quotation marks
		t.Type = String
		b1 := next(byt)
		if !hasBytes(b, 1) {
			return t
		}
		b2 := next(byt)
		for b1 != b2 {
			if isEnd(byt) {
				//b.error("String literal not closed")
				return t
			}
			b2 = next(byt)
		}
		return t
	})

}

func processFixed(len uint, tkn Type) processorFunc {
	return func(b Byterable) Token {
		return markLimits(b, func(byt Byterable) (t Token) {
			// Start and End don't matter
			t.Type = tkn
			byt.SetOffset(byt.Offset() + len)
			return t
		})
	}
}
