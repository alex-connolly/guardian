package lexer

import (
	"fmt"
	"log"
)

type macro struct {
	parameters []string
	tokens     []Token
}

func (l *Lexer) preprocess() {
	log.Println("hi")
	l.addMacros()
}

func (l *Lexer) macroParameters(m macro) {
	l.advance()
	if l.currentToken().Type == TknIdentifier {
		m.parameters = make([]string, 0)
		m.parameters = append(m.parameters, l.TokenString(l.currentToken()))
		l.advance()
		for l.currentToken().Type == TknComma {
			l.advance()
			if l.currentToken().Type != TknIdentifier {
				l.error("Macro parameters must be identifiers")
				l.advance()
			} else {
				m.parameters = append(m.parameters, l.TokenString(l.currentToken()))
				l.advance()
			}
		}
	}
	if l.currentToken().Type != TknCloseBracket {
		l.error("Macro parameters must be closed")
	}
}

func (l *Lexer) multiLineMacro(m macro) {
	start := l.offset
	for l.currentToken().Type != TknCloseBrace {
		l.advance()
	}
	m.tokens = l.Tokens[start:l.offset]
}

func (l *Lexer) singleLineMacro(m macro) {
	start := l.offset
	for l.currentToken().Type != TknNewLine && l.offset < len(l.Tokens) {
		l.advance()
	}
	m.tokens = l.Tokens[start:l.offset]
}

func (l *Lexer) insertToken(name string, m macro) {
	if m.parameters == nil {
		// insert to middle of slice
		l.Tokens = append(l.Tokens[:l.offset], append(m.tokens, l.Tokens[l.offset:]...)...)
	} else {
		if l.Tokens[l.offset+1].Type != TknOpenBracket {
			l.error(fmt.Sprintf("Expected parameters for macro %s", name))
		} else {
			l.advance()
			params := make(map[string][]Token)
			for _, p := range m.parameters {
				start := l.offset
				for l.currentToken().Type != TknComma {
					l.advance()
				}
				params[p] = l.Tokens[start:l.offset]
			}
			// for each token
			tokens := make([]Token, len(m.tokens))
			copy(tokens, m.tokens)
			for i, t := range tokens {
				if t.Type == TknIdentifier {
					key := l.TokenString(t)
					for k, v := range params {
						if key == k {
							tokens = append(tokens[:i], append(v, tokens[i:]...)...)
						}
					}
				}
			}
			l.Tokens = append(l.Tokens[:l.offset], append(tokens, l.Tokens[l.offset:]...)...)
		}
	}
}

func (l *Lexer) addMacros() {
	log.Printf("adding macros")
	for i := 0; i < len(l.Tokens); i++ {
		switch l.Tokens[i].Type {
		case TknMacro:
			// after macro next token must be a key
			l.advance()
			key := l.TokenString(l.currentToken())
			m := macro{}
			if l.Tokens[i].Type == TknOpenBracket {
				l.macroParameters(m)
			} else if l.Tokens[i].Type == TknOpenBrace {
				l.multiLineMacro(m)
			} else {
				l.singleLineMacro(m)
			}
			l.macros[key] = m
			break
		case TknIdentifier:
			for k, v := range l.macros {
				if k == l.TokenString(l.Tokens[i]) {
					l.insertToken(k, v)
				}
			}
			break
		}
	}
}
