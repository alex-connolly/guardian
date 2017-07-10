package parser

import "axia/guardian/go/compiler/ast"

func (p *parser) parseInterfaceDeclaration() {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknInterface)
	identifier := p.parseIdentifier()

	if p.parseOptional(lexer.TknInherits) {

	}

	p.parseRequired(lexer.TknOpenBrace)

	p.scope.Declare(ast.InterfaceDeclarationNode{
		identifier: identifier,
		isAbstract: abstract,
	})
}

func (p *parser) parseClassDeclaration() {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknClass)
	identifier := p.parseIdentifier()

	if p.parseOptional(lexer.TknInherits) {

	}

	if p.parseOptional(lexer.TknIs) {

	}

	p.parseRequired(lexer.TknOpenBrace)

	p.scope.Declare(ast.ClassDeclarationNode{
		identifier: identifier,
		isAbstract: abstract,
	})
}

func (p *parser) parseClassDeclaration() {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknContract)
	identifier := p.parseIdentifier()

	if p.parseOptional(lexer.TknInherits) {

	}

	if p.parseOptional(lexer.TknIs) {

	}

	p.parseRequired(lexer.TknOpenBrace)

	p.scope.Declare(ast.ContractDeclarationNode{
		identifier: identifier,
		isAbstract: abstract,
	})
}

func (p *parser) parseFuncDeclaration() {

	abstract := p.parseOptional(lexer.TknAbstract)
	identifier := p.parseIdentifier()

	params := p.parseParameters()

	results := p.parseResults()

	p.parseRequired(lexer.TknOpenBrace)

	valid := p.scope.Validate(ast.FuncDeclarationNode)

	if !valid {
		p.error("")
	}
	// declare anyway
	p.scope.Declare("func", ast.FuncDeclarationNode{
		identifier: identifier,
		parameters: params,
		results:    results,
		isAbstract: abstract,
	})
}

func (p *parser) parseTypeDeclaration() {

	p.parseRequired(lexer.TknType)
	identifier := p.parseIdentifier()

	oldType := p.parseType()

	valid := p.scope.Validate(ast.TypeDeclarationNode)

	if !valid {
		p.error("")
	}

	p.scope.Declare("type", ast.TypeDeclarationNode{
		identifier: identifier,
	})
}
