package parser

import (
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

func parseInterfaceDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknInterface)
	identifier := p.parseIdentifier()

	var inherits []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
	}

	signatures := p.parseInterfaceSignatures()

	node := ast.InterfaceDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		IsAbstract: abstract,
		Signatures: signatures,
	}

	p.Scope.AddDeclaration(identifier, node)
}

func (p *Parser) parseInterfaceSignatures() []ast.FuncTypeNode {
	p.parseRequired(lexer.TknOpenBrace)

	var sigs []ast.FuncTypeNode
	first := p.parseFuncType()
	sigs = append(sigs, first)
	for !p.parseOptional(lexer.TknCloseBrace) {
		if !p.isFuncType() {
			p.addError("Everything in an interface must be a func type")
			//p.parseConstruct()
			p.next()
		} else {
			sigs = append(sigs, p.parseFuncType())
		}
	}
	return sigs
}

func (p *Parser) parseEnumBody() []string {
	p.parseRequired(lexer.TknOpenBrace)

	// can only contain identifiers
	var enums []string
	first := p.parseIdentifier()
	enums = append(enums, first)
	for p.parseOptional(lexer.TknComma) {
		if p.current().Type != lexer.TknIdentifier {
			p.addError("Everything in an enum must be an identifier")
		} else {
			enums = append(enums, p.parseIdentifier())
		}
	}
	p.parseRequired(lexer.TknCloseBrace)
	return enums
}

func parseEnumDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknEnum)
	identifier := p.parseIdentifier()

	var inherits []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
	}

	enums := p.parseEnumBody()

	node := ast.EnumDeclarationNode{
		IsAbstract: abstract,
		Identifier: identifier,
		Inherits:   inherits,
		Enums:      enums,
	}

	p.Scope.AddDeclaration(identifier, node)
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

	body := p.parseEnclosedScope()

	node := ast.ClassDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		Interfaces: interfaces,
		IsAbstract: abstract,
		Body:       body,
	}

	p.Scope.AddDeclaration(identifier, node)
}

func parseContractDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknContract)
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

	valids := []ast.NodeType{
		ast.ClassDeclaration, ast.InterfaceDeclaration,
		ast.EventDeclaration, ast.ExplicitVarDeclaration,
		ast.TypeDeclaration, ast.EnumDeclaration,
		ast.LifecycleDeclaration, ast.FuncDeclaration,
	}

	body := p.parseEnclosedScope(valids...)

	node := ast.ContractDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		Interfaces: interfaces,
		IsAbstract: abstract,
		Body:       body,
	}

	p.Scope.AddDeclaration(identifier, node)
}

func (p *Parser) parseType() ast.Node {
	switch {
	case p.isArrayType():
		return p.parseArrayType()
	case p.isMapType():
		return p.parseMapType()
	case p.isFuncType():
		return p.parseFuncType()
	case p.current().Type == lexer.TknIdentifier:
		return p.parseReference()
	}
	return nil
}

func (p *Parser) parseTypeList() []ast.Node {
	var types []ast.Node
	first := p.parseType()
	types = append(types, first)
	for p.parseOptional(lexer.TknComma) {
		types = append(types, p.parseType())
	}
	return types
}

func (p *Parser) parseVarDeclaration() ast.ExplicitVarDeclarationNode {
	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(lexer.TknComma) {
		names = append(names, p.parseIdentifier())
	}

	dType := p.parseType()

	return ast.ExplicitVarDeclarationNode{
		DeclaredType: dType,
		Identifiers:  names,
	}
}

func (p *Parser) parseParameters() []ast.ExplicitVarDeclarationNode {
	var params []ast.ExplicitVarDeclarationNode
	p.parseRequired(lexer.TknOpenBracket)
	if !p.parseOptional(lexer.TknCloseBracket) {
		params = append(params, p.parseVarDeclaration())
		for p.parseOptional(lexer.TknComma) {
			params = append(params, p.parseVarDeclaration())
		}
		p.parseRequired(lexer.TknCloseBracket)
	}
	return params
}

// currently not supporting named return types
// reasoning: confusing to user
// returns can either be single
// string {
// or multiple
// (string, string) {
// or none
// {
func (p *Parser) parseResults() []ast.Node {
	if p.parseOptional(lexer.TknOpenBracket) {
		types := p.parseTypeList()
		p.parseRequired(lexer.TknCloseBracket)
		return types
	}
	if p.current().Type != lexer.TknOpenBrace {
		return p.parseTypeList()
	}
	return nil
}

func parseFuncDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)

	p.parseRequired(lexer.TknFunc)

	identifier := p.parseIdentifier()

	params := p.parseParameters()

	results := p.parseResults()

	body := p.parseEnclosedScope(ast.ExplicitVarDeclaration, ast.FuncDeclaration)

	node := ast.FuncDeclarationNode{
		Identifier: identifier,
		Parameters: params,
		Results:    results,
		IsAbstract: abstract,
		Body:       body,
	}

	p.Scope.AddDeclaration(identifier, node)
}

func parseLifecycleDeclaration(p *Parser) {

	category := p.parseRequired(lexer.GetLifecycles()...)

	params := p.parseParameters()

	body := p.parseEnclosedScope(ast.ExplicitVarDeclaration, ast.FuncDeclaration)

	node := ast.LifecycleDeclarationNode{
		Category:   category,
		Parameters: params,
		Body:       body,
	}

	p.Scope.AddDeclaration("constructor", node)
}

func parseTypeDeclaration(p *Parser) {
	p.parseRequired(lexer.TknType)
	identifier := p.parseIdentifier()

	value := p.parseType()

	n := ast.TypeDeclarationNode{
		Identifier: identifier,
		Value:      value,
	}

	p.Scope.AddDeclaration(identifier, n)
}

func (p *Parser) parseMapType() ast.MapTypeNode {

	p.parseRequired(lexer.TknMap)
	p.parseRequired(lexer.TknOpenSquare)

	key := p.parseType()

	p.parseRequired(lexer.TknCloseSquare)

	value := p.parseType()

	mapType := ast.MapTypeNode{
		Key:   key,
		Value: value,
	}

	return mapType
}

func (p *Parser) parseFuncType() ast.FuncTypeNode {

	f := ast.FuncTypeNode{}
	p.parseRequired(lexer.TknFunc)
	p.parseRequired(lexer.TknOpenBracket)
	if !p.parseOptional(lexer.TknCloseBracket) {
		f.Parameters = p.parseTypeList()
		p.parseRequired(lexer.TknCloseBracket)
	}
	if p.parseOptional(lexer.TknOpenBracket) {
		f.Results = p.parseTypeList()
		p.parseRequired(lexer.TknCloseBracket)
	}
	f.Results = p.parseTypeList()
	return f
}

func (p *Parser) parseArrayType() ast.ArrayTypeNode {
	p.parseRequired(lexer.TknOpenSquare)

	typ := p.parseType()

	p.parseRequired(lexer.TknCloseSquare)

	/*if p.parseOptional(lexer.TknColon) {
		//	max = p.parseExpression()
	}*/

	return ast.ArrayTypeNode{
		Value: typ,
	}
}

func parseExplicitVarDeclaration(p *Parser) {
	// parse variable Names
	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(lexer.TknComma) {
		names = append(names, p.parseIdentifier())
	}
	// parse type
	typ := p.parseType()

	node := ast.ExplicitVarDeclarationNode{
		Identifiers:  names,
		DeclaredType: typ,
	}
	// surely needs to be a declaration in some contexts
	p.Scope.AddSequential(node)
}

func parseEventDeclaration(p *Parser) {
	p.parseRequired(lexer.TknEvent)
	name := p.lexer.TokenString(p.current())
	p.next()
	p.parseRequired(lexer.TknOpenBracket)
	var types []ast.ReferenceNode
	if !p.parseOptional(lexer.TknCloseBracket) {
		types = p.parseReferenceList()
		p.parseRequired(lexer.TknCloseBracket)
	}

	node := ast.EventDeclarationNode{
		Identifier: name,
		Parameters: types,
	}
	p.Scope.AddDeclaration(name, node)
}
