package parser

import (
	"io/ioutil"
	"log"
)

// ParseFile ...
func ParseFile(path string) *Parser {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	return ParseBytes(bytes)
}

// ParseString ...
func ParseString(data string) *Parser {
	return ParseBytes([]byte(data))
}

// ParseBytes ...
func ParseBytes(data []byte) *Parser {
	p.lexer = lexer.LexBytes(data)
	p.scope = ast.FileNode{}
	p.run()
}
