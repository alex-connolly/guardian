package token

func isIdentifierByte(b byte) bool {
	return ('A' <= b && b <= 'Z') ||
		('a' <= b && b <= 'z') ||
		('0' <= b && b <= '9') ||
		(b == '_')
}

func isIdentifier(b Byterable) bool {
	return isIdentifierByte(current(b))
}

func isNumeric(b Byterable) bool {
	return ('0' <= current(b) && current(b) <= '9')
}

func isInteger(b Byterable) bool {
	saved := b.Offset()
	if hasBytes(b, 2) {
		if current(b) == '-' {
			next(b)
			flag := isNumeric(b)
			b.SetOffset(saved)
			return flag
		}
	}
	return isNumeric(b)
}

func isFloat(b Byterable) bool {
	saved := b.Offset()
	if hasBytes(b, 1) && current(b) == '-' {
		next(b)
	}
	for hasBytes(b, 1) && '0' <= current(b) && current(b) <= '9' {
		next(b)
	}
	if !hasBytes(b, 1) {
		b.SetOffset(saved)
		return false
	}
	if !hasBytes(b, 1) || current(b) != '.' {
		b.SetOffset(saved)
		return false
	}
	next(b)
	if !hasBytes(b, 1) || !isNumeric(b) {
		b.SetOffset(saved)
		return false
	}
	b.SetOffset(saved)
	return true
}

func isString(b Byterable) bool {
	return (current(b) == '"') || (current(b) == '`')
}

func isWhitespace(b Byterable) bool {
	return (current(b) == ' ') || (current(b) == '\t') || (current(b) == '\v') || (current(b) == '\f')
}

func isNewLine(b Byterable) bool {
	// Took these cases from the gravity Lexer
	// probably incomplete
	// CR: Carriage Return, U+000D (UTF-8 in hex: 0D)
	// LF: Line Feed, U+000A (UTF-8 in hex: 0A)
	// CR+LF: CR (U+000D) followed by LF (U+000A) (UTF-8 in hex: 0D0A)

	// LF
	if current(b) == 0x0A {
		return true
	}

	// CR+LF or CR
	if current(b) == 0x0D {
		if hasBytes(b, 1) && b.Bytes()[b.Offset()+1] == 0x0A {
			return true
		}
	}

	// UTF-8 cases https://en.wikipedia.org/wiki/Newline#Unicode

	// NEL: Next Line, U+0085 (UTF-8 in hex: C285)
	if current(b) == 0xC2 {
		if hasBytes(b, 1) && b.Bytes()[b.Offset()+1] == 0x85 {
			return true
		}
	}

	if current(b) == 0xE2 {
		if hasBytes(b, 1) && b.Bytes()[b.Offset()+1] == 0x85 {
			return true
		}
	}

	// and probably more not handled here
	return (current(b) == '\n')
}

func isCommentClose(b Byterable) bool {
	if hasBytes(b, 2) {
		return current(b) == '*' && b.Bytes()[b.Offset()+1] == '/'
	}
	return false
}

func isCommentOpen(b Byterable) bool {
	if hasBytes(b, 2) {
		return current(b) == '/' && b.Bytes()[b.Offset()+1] == '*'
	}
	return false
}

func isLineComment(b Byterable) bool {
	if hasBytes(b, 2) {
		return current(b) == '/' && b.Bytes()[b.Offset()+1] == '/'
	}
	return false
}

func isCharacter(b Byterable) bool {
	return (current(b) == '\'')
}
