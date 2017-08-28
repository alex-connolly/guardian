package lexer

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestPreprocessMacroSimple(t *testing.T) {
	l := LexString(`
        macro VALUE 20

        name = VALUE
        `)
	goutil.Assert(t, len(l.Tokens) == 3, fmt.Sprintf("wrong token length: %d", len(l.Tokens)))
}

func TestPreprocessMacroArguments(t *testing.T) {
	l := LexString(`
        macro ASSIGN(a, b){
            a = b
        }

        ASSIGN(name, 20)
        `)
	goutil.Assert(t, len(l.Tokens) == 3, fmt.Sprintf("wrong token length: %d", len(l.Tokens)))
}

func TestPreprocessMacroNestedSimple(t *testing.T) {
	l := LexString(`
        macro VALUE 20
        macro V2 VALUE

        name = V2
        `)
	goutil.Assert(t, len(l.Tokens) == 3, fmt.Sprintf("wrong token length: %d", len(l.Tokens)))
}

func TestPreprocessMacroNestedMixed(t *testing.T) {
	l := LexString(`
        macro EQUALS =
        macro ASSIGN(a, b){
            a EQUALS b
        }

        ASSIGN(name, 20)
        `)
	goutil.Assert(t, len(l.Tokens) == 3, fmt.Sprintf("wrong token length: %d", len(l.Tokens)))
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
	goutil.Assert(t, len(l.Tokens) == 9, fmt.Sprintf("wrong token length: %d", len(l.Tokens)))
}

func TestPreprocessMacroNestedArgumentsSingleLine(t *testing.T) {
	l := LexString(`
        macro MAX(a, b) a > b ? a : b
        macro ASSIGN_MAX(a, b) a = MAX(a, b)
        ASSIGN_MAX(name, 20)
        `)
	goutil.Assert(t, len(l.Tokens) == 9, fmt.Sprintf("wrong token length: %d", len(l.Tokens)))
}
