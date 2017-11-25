package lexer

import (
	"io/ioutil"

	"github.com/end-r/guardian/util"
)

// Lexer ...
type Lexer struct {
	buffer      []byte
	byteOffset  int
	line        int
	column      int
	Tokens      []Token
	tokenOffset int
	errors      []util.Error
	//macros      map[string]macro
}

// Lex ...
func Lex(bytes []byte) (tokens []Token, errs []util.Error) {
	l := new(Lexer)
	l.byteOffset = 0
	l.buffer = bytes
	l.next()
	return l.Tokens, l.errors
}

// LexFile ...
func LexFile(path string) (tokens []Token, errs []util.Error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, append(errs, util.Error{
			Message: "File does not exist",
		})
	}
	return Lex(bytes)
}

// LexString lexes a string
func LexString(str string) (tokens []Token, errs []util.Error) {
	return Lex([]byte(str))
}
