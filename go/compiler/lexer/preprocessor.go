package lexer

import "fmt"

type macro struct {
	parameters []string
	tokens     []Token
}

func (l *Lexer) preprocess() {

}

func (l *Lexer) addMacros() {
	for i := 0; i < len(l.Tokens); i++ {
		switch l.Tokens[i].Type {
		case TknMacro:
			// after macro next token must be a key
			key := l.TokenString(l.Tokens[i+1])
			m := macro{}
			if l.Tokens[i].Type == TknOpenBracket {
				i++
				if l.Tokens[i].Type == TknIdentifier {
					m.parameters = make([]string, 0)
					m.parameters = append(m.parameters, l.TokenString(l.Tokens[i]))
					i++
					for l.Tokens[i].Type == TknComma {
						i++
						if l.Tokens[i].Type != TknIdentifier {
							l.error("Macro parameters must be identifiers")
							i++
						} else {
							m.parameters = append(m.parameters, l.TokenString(l.Tokens[i]))
							i++
						}
					}
				}
				if l.Tokens[i].Type != TknCloseBracket {
					l.error("Macro parameters must be closed")
				}
			}
			l.macros[key] = m
			//TODO: remove macro tokens
			break
		case TknIdentifier:
			for k, v := range l.macros {
				if k == l.TokenString(l.Tokens[i]) {
					if v.parameters == nil {
						// insert to middle of slice
						l.Tokens = append(l.Tokens[:i], append(v.tokens, l.Tokens[i:]...)...)
					} else {
						if l.Tokens[i+1].Type != TknOpenBracket {
							l.error(fmt.Sprintf("Expected parameters for macro %s", k))
						} else {

						}
					}
				}
			}
			break
		}
	}
}
