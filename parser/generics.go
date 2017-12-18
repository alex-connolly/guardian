package parser

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/token"
)

func (p *Parser) parseGenericsDeclaration() {
	p.parseRequired(token.Lss)
	if p.parseOptional(token.Gtr) {
		p.addError("")
		return
	}

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
		Interfaces: interfaces,
		Inherits:   inherits,
	}
}
