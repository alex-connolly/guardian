package parser

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/token"
)

func (p *Parser) parsePossibleGenerics() []*ast.GenericDeclarationNode {
	if p.parseOptional(token.Lss) {
		gens := p.parseGenerics()
		p.parseRequired(token.Gtr)
		return gens
	}
	return nil
}

func (p *Parser) parseGenerics() []*ast.GenericDeclarationNode {
	var generics []*ast.GenericDeclarationNode
	generics = append(generics, p.parseGeneric())
	for p.parseOptional(token.Comma) {
		generics = append(generics, p.parseGeneric())
	}
	return generics
}

// TODO: java ? wildcards
// upper and lower bounds?
func (p *Parser) parseGeneric() *ast.GenericDeclarationNode {
	identifier := p.parseIdentifier()
	var inherits, interfaces []*ast.PlainTypeNode

	if p.parseOptional(token.Inherits) {
		inherits = p.parsePlainTypeList()
		if p.parseOptional(token.Is) {
			interfaces = p.parsePlainTypeList()
		}
	} else if p.parseOptional(token.Is) {
		interfaces = p.parsePlainTypeList()
		if p.parseOptional(token.Inherits) {
			inherits = p.parsePlainTypeList()
		}
	}
	return &ast.GenericDeclarationNode{
		Identifier: identifier,
		Implements: interfaces,
		Inherits:   inherits,
	}
}
