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

func isInteger(b Byterable) bool {
	return ('0' <= current(b) && current(b) <= '9')
}

func isFloat(b Byterable) bool {
	saved := b.Offset()
	for hasBytes(b, 1) && '0' <= current(b) && current(b) <= '9' {
		next(b)
	}
	if !hasBytes(b, 1) || current(b) != '.' {
		b.SetOffset(saved)
		return false
	}
	next(b)
	if !hasBytes(b, 1) || !('0' <= current(b) && current(b) <= '9') {
		b.SetOffset(saved)
		return false
	}
	b.SetOffset(saved)
	return true
}

func isString(b Byterable) bool {
	return (current(b) == '"')
}

func isWhitespace(b Byterable) bool {
	return (current(b) == ' ') || (current(b) == '\t')
}

func isNewLine(b Byterable) bool {
	return (current(b) == '\n')
}

func isCharacter(b Byterable) bool {
	return (current(b) == '\'')
}

func is(a string) isFunc {
	return func(b Byterable) bool {
		if b.Offset()+len(a) > len(b.Bytes()) {
			return false
		}
		//	fmt.Printf("cmp %s to %s\n", string(b.Bytes()[b.offset:b.offset+len(a)]), a)
		return string(b.Bytes()[b.Offset():b.Offset()+len(a)]) == a
	}
}

func isDistinct(a string) isFunc {
	return func(b Byterable) bool {
		if !hasBytes(b, len(a)) {
			return false
		}
		end := b.Offset() + len(a)
		if string(b.Bytes()[b.Offset():end]) != a {
			return false
		}
		if !hasBytes(b, len(a)+1) {
			return true
		}
		return !isIdentifierByte(b.Bytes()[end])
	}
}
