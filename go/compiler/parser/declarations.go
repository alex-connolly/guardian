package parser

import (
	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
)

func parseInterfaceDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknInterface)
	identifier := p.parseIdentifier()

	var inherits []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
	}

	body := ast.ScopeNode{
		ValidTypes: []ast.NodeType{},
	}

	p.parseEnclosedScope(&body)

	node := ast.InterfaceDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		IsAbstract: abstract,
		Body:       body,
	}

	p.Scope.Declare("interface", node)
}

func parseEnumDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknEnum)
	identifier := p.parseIdentifier()

	var inherits []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
	}

	body := ast.ScopeNode{
		ValidTypes: []ast.NodeType{ast.Reference},
	}

	p.parseEnclosedScope(&body)

	node := ast.EnumDeclarationNode{
		IsAbstract: abstract,
		Identifier: identifier,
		Inherits:   inherits,
		Body:       body,
	}

	p.Scope.Declare("enum", node)
}

// like any list parser, but enforces that each node must be a reference
func (p *Parser) parseReferenceList() []ast.ReferenceNode {
	var refs []ast.ReferenceNode
	first := p.parseReference()
	refs = append(refs, first)
	for p.parseOptional(lexer.TknComma) {
		refs = append(refs, p.parseReference())
	}
	return refs
}

func parseClassDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknClass)
	identifier := p.parseIdentifier()

	// is and inherits can be in any order

	var inherits, interfaces []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
		if p.parseOptional(lexer.TknIs) {
			interfaces = p.parseReferenceList()
		}
	} else if p.parseOptional(lexer.TknIs) {
		interfaces = p.parseReferenceList()
		if p.parseOptional(lexer.TknInherits) {
			inherits = p.parseReferenceList()
		}
	}

	p.parseRequired(lexer.TknOpenBrace)

	body := ast.ScopeNode{
		ValidTypes: []ast.NodeType{},
	}

	p.parseEnclosedScope(&body)

	node := ast.ClassDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		Interfaces: interfaces,
		IsAbstract: abstract,
		Body:       body,
	}

	p.Scope.Declare("class", node)
}

func parseContractDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknContract)
	identifier := p.parseIdentifier()

	// is and inherits can be in any order

	var inherits []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
	}

	body := ast.ScopeNode{
		ValidTypes: []ast.NodeType{},
	}

	p.parseEnclosedScope(&body)

	node := ast.ContractDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		IsAbstract: abstract,
		Body:       body,
	}

	p.Scope.Declare("contract", node)
}

func (p *Parser) parseParameters() []ast.ExplicitVarDeclarationNode {
	var params []ast.ExplicitVarDeclarationNode
	p.parseRequired(lexer.TknOpenBracket)

	p.parseRequired(lexer.TknCloseBracket)
	return params
}

func (p *Parser) parseResults() []ast.ReferenceNode {
	// currently not supporting named return types
	// reasoning: confusing to user
	// returns can either be single
	// string {
	// or multiple
	// (string, string) {
	// or none
	// {
	if p.parseOptional(lexer.TknOpenBracket) {
		refs := p.parseReferenceList()
		p.parseOptional(lexer.TknCloseBracket)
		return refs
	}
	if p.current().Type == lexer.TknIdentifier {
		return p.parseReferenceList()
	}
	return nil
}

func parseFuncDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)

	identifier := p.parseIdentifier()

	params := p.parseParameters()

	results := p.parseResults()

	p.parseRequired(lexer.TknOpenBrace)

	p.validate(ast.FuncDeclaration)

	body := ast.ScopeNode{
		ValidTypes: []ast.NodeType{},
	}

	p.parseEnclosedScope(&body)

	node := ast.FuncDeclarationNode{
		Identifier: identifier,
		Parameters: params,
		Results:    results,
		IsAbstract: abstract,
		Body:       body,
	}

	p.Scope.Declare("func", node)
}

func parseConstructorDeclaration(p *Parser) {

	p.parseRequired(lexer.TknConstructor)

	params := p.parseParameters()

	p.parseRequired(lexer.TknOpenBrace)

	p.validate(ast.ConstructorDeclaration)

	body := ast.ScopeNode{
		ValidTypes: []ast.NodeType{},
	}

	p.parseEnclosedScope(&body)

	node := ast.ConstructorDeclarationNode{
		Parameters: params,
		Body:       body,
	}

	p.Scope.Declare("constructor", node)
}

func parseTypeDeclaration(p *Parser) {
	p.parseRequired(lexer.TknType)
	identifier := p.parseIdentifier()

	value := p.parseReference()

	n := ast.TypeDeclarationNode{
		Identifier: identifier,
		Value:      value,
	}

	p.Scope.Declare("type", n)
}

func (p *Parser) parseMapType() ast.Node {

	p.parseRequired(lexer.TknMap)
	p.parseRequired(lexer.TknOpenSquare)

	key := p.parseType()

	p.parseRequired(lexer.TknCloseSquare)

	value := p.parseType()

	p.validate(ast.MapType)

	mapType := ast.MapTypeNode{
		Key:   key,
		Value: value,
	}

	return mapType
}

func (p *Parser) parseArrayType() ast.Node {
	p.parseRequired(lexer.TknOpenSquare)

	//typ := p.parseExpression()
	//var max ast.Node

	if p.parseOptional(lexer.TknColon) {
		//	max = p.parseExpression()
	}
	p.validate(ast.ArrayType)

	arrayType := ast.ArrayTypeNode{
	//Value: typ,
	}
	return arrayType
}

func parseExplicitVarDeclaration(p *Parser) {

	// parse variable Names
	names := make([]string, 0)
	names = append(names, p.lexer.TokenString(p.current()))
	p.next()
	for p.parseOptional(lexer.TknComma) {
		names = append(names, p.lexer.TokenString(p.current()))
	}
	// parse type
	dType := p.parseReference()

	node := ast.ExplicitVarDeclarationNode{
		Identifiers:  names,
		DeclaredType: dType,
	}

	p.Scope.Declare("var", node)
}

func parseEventDeclaration(p *Parser) {
	p.parseRequired(lexer.TknEvent)
	name := p.lexer.TokenString(p.current())
	p.next()
	p.parseRequired(lexer.TknOpenBracket)
	types := p.parseReferenceList()
	p.parseRequired(lexer.TknCloseBracket)
	node := ast.EventDeclarationNode{
		Identifier: name,
		Parameters: types,
	}
	p.Scope.Declare("event", node)
}
