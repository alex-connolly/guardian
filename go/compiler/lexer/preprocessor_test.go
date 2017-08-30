package lexer

import (
	"testing"
)

func TestPreprocessMacroSimple(t *testing.T) {
	l := LexString(`
        macro VALUE 20

        name = VALUE
        `)
	checkTokens(t, l.Tokens, []TokenType{
		TknIdentifier, TknAssign, TknIdentifier, TknNewLine,
	})
}

func TestPreprocessMacroArguments(t *testing.T) {
	l := LexString(`
        macro ASSIGN(a, b){
            a = b
        }

        ASSIGN(name, 20)
        `)
	checkTokens(t, l.Tokens, []TokenType{
		TknIdentifier, TknAssign, TknNumber, TknNewLine,
	})
}

func TestPreprocessMacroNestedSimple(t *testing.T) {
	l := LexString(`
        macro VALUE 20
        macro V2 VALUE

        name = V2
        `)
	checkTokens(t, l.Tokens, []TokenType{
		TknIdentifier, TknAssign, TknNumber, TknNewLine,
	})
}

func TestPreprocessMacroNestedMixed(t *testing.T) {
	l := LexString(`
        macro EQUALS =
        macro ASSIGN(a, b){
            a EQUALS b
        }

        ASSIGN(name, 20)
        `)
	checkTokens(t, l.Tokens, []TokenType{
		TknIdentifier, TknAssign, TknNumber, TknNewLine,
	})
}

func TestPreprocessMacroNestedArguments(t *testing.T) {
	l := LexString(`
        macro MAX(a, b) {
            a > b ? a : b
        }
        macro ASSIGN_MAX(a, b){
            a = MAX(a, b)
        }

        ASSIGN_MAX(name, 20)
        `)
	checkTokens(t, l.Tokens, []TokenType{
		TknIdentifier, TknAssign, TknIdentifier, TknGtr, TknIdentifier, TknTernary,
		TknIdentifier, TknColon, TknIdentifier, TknNewLine,
	})
}

func TestPreprocessMacroNestedArgumentsSingleLine(t *testing.T) {
	l := LexString(`
        macro MAX(a, b) a > b ? a : b
        macro ASSIGN_MAX(a, b) a = MAX(a, b)
        ASSIGN_MAX(name, 20)
        `)
	checkTokens(t, l.Tokens, []TokenType{
		TknIdentifier, TknAssign, TknIdentifier, TknGtr, TknIdentifier, TknTernary,
		TknIdentifier, TknColon, TknIdentifier, TknNewLine,
	})
}
