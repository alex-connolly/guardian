package parser

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/token"
)

func (p *Parser) parsePossibleGenerics() []*ast.GenericDeclarationNode {
	if p.nextTokens(token.Lss) {
		return p.parseGenerics()
	}
	return nil
}

func (p *Parser) parseGenerics() []*ast.GenericDeclarationNode {
	p.parseRequired(token.Lss)
	var generics []*ast.GenericDeclarationNode
	generics = append(generics, p.parseGeneric())
	for p.parseOptional(token.Or) {
		generics = append(generics, p.parseGeneric())
	}
	p.parseRequired(token.Gtr)
	return generics
}

// TODO: java ? wildcards
// upper and lower bounds?
// can't use simple inherits/Implements
// T inherits A, B, C
// have to split T inherits A, B,
// <T inherits A | S | R>
// <T|S|R>
func (p *Parser) parseGeneric() *ast.GenericDeclarationNode {

	g := new(ast.GenericDeclarationNode)

	g.Begin = p.getCurrentTokenLocation()

	g.Identifier = p.parseIdentifier()
	g.Implements = make([]*ast.PlainTypeNode, 0)
	g.Inherits = make([]*ast.PlainTypeNode, 0)

	if p.parseOptional(token.Inherits) {
		g.Inherits = p.parsePlainTypeList()
		if p.parseOptional(token.Is) {
			g.Implements = p.parsePlainTypeList()
		}
	} else if p.parseOptional(token.Is) {
		g.Implements = p.parsePlainTypeList()
		if p.parseOptional(token.Inherits) {
			g.Inherits = p.parsePlainTypeList()
		}
	}

	g.Final = p.getCurrentTokenLocation()

	return g
}
