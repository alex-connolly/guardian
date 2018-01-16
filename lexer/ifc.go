package lexer

import (
	"io/ioutil"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/util"
)

// Lexer ...
type Lexer struct {
	buffer      []byte
	byteOffset  uint
	line        uint
	column      int
	tokens      []token.Token
	tokenOffset int
	errors      util.Errors
	fileName    string
}

// Lex ...
func Lex(name string, bytes []byte) *Lexer {
	l := new(Lexer)
	l.fileName = name
	l.byteOffset = 0
	l.buffer = bytes
	l.next()
	return l
}

// LexFile ...
func LexFile(path string) *Lexer {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		l := new(Lexer)
		l.errors = []util.Error{
			util.Error{
				Message: "File does not exist",
			},
		}
		return l
	}
	return Lex(path, bytes)
}

// LexString lexes a string
func LexString(str string) *Lexer {
	return Lex("input", []byte(str))
}
