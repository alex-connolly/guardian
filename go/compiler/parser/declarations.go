package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

func parseInterfaceDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknInterface)
	identifier := p.parseIdentifier()

	if p.parseOptional(lexer.TknInherits) {

	}

	p.parseRequired(lexer.TknOpenBrace)

	p.scope.Declare("", ast.InterfaceDeclarationNode{
		Identifier: identifier,
		IsAbstract: abstract,
	})
}

func parseClassDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknClass)
	identifier := p.parseIdentifier()

	if p.parseOptional(lexer.TknInherits) {

	}

	if p.parseOptional(lexer.TknIs) {

	}

	p.parseRequired(lexer.TknOpenBrace)

	p.scope.Declare("", ast.ClassDeclarationNode{
		Identifier: identifier,
		IsAbstract: abstract,
	})
}

func parseContractDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknContract)
	identifier := p.parseIdentifier()

	if p.parseOptional(lexer.TknInherits) {

	}

	if p.parseOptional(lexer.TknIs) {

	}

	p.parseRequired(lexer.TknOpenBrace)

	p.scope.Declare("", ast.ContractDeclarationNode{
		Identifier: identifier,
		IsAbstract: abstract,
	})
}

func (p *Parser) parseParameters() []ast.Node {
	return nil
}

func (p *Parser) parseResults() []ast.Node {
	return nil
}

func parseFuncDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	identifier := p.parseIdentifier()

	params := p.parseParameters()

	results := p.Parseresults()

	p.parseRequired(lexer.TknOpenBrace)

	p.validate(ast.FuncDeclaration)

	p.scope.Declare("func", ast.FuncDeclarationNode{
		Identifier: identifier,
		Parameters: params,
		Results:    results,
		IsAbstract: abstract,
	})
}

func parseTypeDeclaration(p *Parser) {

	p.parseRequired(lexer.TknType)
	identifier := p.parseIdentifier()
	//oldType := p.parseType()

	p.validate(ast.TypeDeclaration)

	p.scope.Declare("type", ast.TypeDeclarationNode{
		Identifier: identifier,
	})
}

func parseMapType(p *Parser) {

	p.parseRequired(lexer.TknMap)
	p.parseRequired(lexer.TknOpenSquare)

	key := p.parseType()

	p.parseRequired(lexer.TknCloseSquare)

	value := p.parseType()

	p.validate(ast.MapType)

	p.scope.Declare("", ast.MapTypeNode{
		Key:   key,
		Value: value,
	})
}

func parseArrayType(p *Parser) {
	p.parseRequired(lexer.TknOpenSquare)

	typ := p.parseExpression()
	//var max ast.Node

	if p.parseOptional(lexer.TknColon) {
		//	max = p.parseExpression()
	}
	p.validate(ast.ArrayType)

	p.scope.Declare("", ast.ArrayTypeNode{
		Value: typ,
	})
}
