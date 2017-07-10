package parser

import "go/ast"

type parser struct {
	constructs []construct
	scope      ast.Node
	lexer      *lexer
	index      int
	errs       []string
}
