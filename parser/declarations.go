package parser

import (
	"strconv"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

func (p *Parser) parseKeywords(targets ...lexer.TokenType) []lexer.TokenType {
	var mods []lexer.TokenType
	for p.hasTokens(1) {
		for _, t := range targets {
			if p.current().Type == t {
				return mods
			}
		}
		if p.current().Type.IsModifier() || p.current().Type.IsDeclaration() {
			mods = append(mods, p.current().Type)
			p.next()
		} else {
			p.next()
			return mods
		}
	}

	return mods
}

func parseInterfaceDeclaration(p *Parser) {

	keywords := p.parseKeywords(token.Interface)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(token.Interface)
	identifier := p.parseIdentifier()

	var inherits []*ast.PlainTypeNode

	if p.parseOptional(token.Inherits) {
		inherits = p.parsePlainTypeList()
	}

	signatures := p.parseInterfaceSignatures()

	node := ast.InterfaceDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		Modifiers:  keywords,
		Signatures: signatures,
	}

	p.scope.AddDeclaration(identifier, &node)
}

func (p *Parser) parseInterfaceSignatures() []*ast.FuncTypeNode {

	p.parseRequired(token.OpenBrace)

	for isNewLine(p) {
		parseNewLine(p)
	}

	var sigs []*ast.FuncTypeNode

	first := p.parseFuncType()
	sigs = append(sigs, first)

	for isNewLine(p) {
		parseNewLine(p)
	}

	for !p.parseOptional(token.CloseBrace) {

		if !p.isFuncType() {
			p.addError(errInvalidInterfaceProperty)
			//p.parseConstruct()
			p.next()
		} else {
			sigs = append(sigs, p.parseFuncType())
		}

		for isNewLine(p) {
			parseNewLine(p)
		}

	}
	return sigs
}

func (p *Parser) parseEnumBody() []string {

	p.parseRequired(token.OpenBrace)
	var enums []string
	// remove all new lines before the fist identifier
	for isNewLine(p) {
		parseNewLine(p)
	}

	if !p.parseOptional(token.CloseBrace) {
		// can only contain identifiers split by commas and newlines
		first := p.parseIdentifier()
		enums = append(enums, first)
		for p.parseOptional(token.Comma) {

			for isNewLine(p) {
				parseNewLine(p)
			}

			if p.current().Type == token.Identifier {
				enums = append(enums, p.parseIdentifier())
			} else {
				p.addError(errInvalidEnumProperty)
			}
		}

		for isNewLine(p) {
			parseNewLine(p)
		}

		p.parseRequired(token.CloseBrace)
	}
	return enums
}

func parseEnumDeclaration(p *Parser) {

	keywords := p.parseKeywords(token.Enum)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(token.Enum)
	identifier := p.parseIdentifier()

	var inherits []*ast.PlainTypeNode

	if p.parseOptional(token.Inherits) {
		inherits = p.parsePlainTypeList()
	}

	enums := p.parseEnumBody()

	node := ast.EnumDeclarationNode{
		Modifiers:  keywords,
		Identifier: identifier,
		Inherits:   inherits,
		Enums:      enums,
	}

	p.scope.AddDeclaration(identifier, &node)
}

func (p *Parser) parsePlainType() *ast.PlainTypeNode {

	variable := p.parseOptional(token.Ellipsis)

	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(token.Dot) {
		names = append(names, p.parseIdentifier())
	}
	return &ast.PlainTypeNode{
		Names:    names,
		Variable: variable,
	}
}

// like any list parser, but enforces that each node must be a plain type
func (p *Parser) parsePlainTypeList() []*ast.PlainTypeNode {
	var refs []*ast.PlainTypeNode
	refs = append(refs, p.parsePlainType())
	for p.parseOptional(token.Comma) {
		refs = append(refs, p.parsePlainType())
	}
	return refs
}

func parseClassDeclaration(p *Parser) {

	keywords := p.parseKeywords(token.Class)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(token.Class)

	identifier := p.parseIdentifier()

	// is and inherits can be in any order

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

	body := p.parseBracesScope()

	node := ast.ClassDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		Interfaces: interfaces,
		Modifiers:  keywords,
		Body:       body,
	}

	p.scope.AddDeclaration(identifier, &node)
}

func parseContractDeclaration(p *Parser) {

	keywords := p.parseKeywords(token.Contract)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(token.Contract)
	identifier := p.parseIdentifier()

	// is and inherits can be in any order

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

	valids := []ast.NodeType{
		ast.ClassDeclaration, ast.InterfaceDeclaration,
		ast.EventDeclaration, ast.ExplicitVarDeclaration,
		ast.TypeDeclaration, ast.EnumDeclaration,
		ast.LifecycleDeclaration, ast.FuncDeclaration,
	}

	body := p.parseBracesScope(valids...)

	node := ast.ContractDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		Interfaces: interfaces,
		Modifiers:  keywords,
		Body:       body,
	}

	p.scope.AddDeclaration(identifier, &node)
}

func (p *Parser) parseType() ast.Node {
	switch {
	case p.isArrayType():
		return p.parseArrayType()
	case p.isMapType():
		return p.parseMapType()
	case p.isFuncType():
		return p.parseFuncType()
		// TODO: should be able to do with isPlainType() but causes crashes
	case p.isNextToken(token.Identifier):
		return p.parsePlainType()
	}
	return nil
}

func (p *Parser) parseVarDeclaration() *ast.ExplicitVarDeclarationNode {

	keywords := p.parseKeywords(token.Identifier)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(token.Comma) {
		names = append(names, p.parseIdentifier())
	}

	dType := p.parseType()

	return &ast.ExplicitVarDeclarationNode{
		Modifiers:    keywords,
		DeclaredType: dType,
		Identifiers:  names,
	}
}

func (p *Parser) parseParameters() []*ast.ExplicitVarDeclarationNode {
	var params []*ast.ExplicitVarDeclarationNode
	p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		params = append(params, p.parseVarDeclaration())
		for p.parseOptional(token.Comma) {
			params = append(params, p.parseVarDeclaration())
		}
		p.parseRequired(token.CloseBracket)
	}
	return params
}

func (p *Parser) parseTypeList() []ast.Node {
	var types []ast.Node
	first := p.parseType()
	types = append(types, first)
	for p.parseOptional(token.Comma) {
		types = append(types, p.parseType())
	}
	return types
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
	if p.parseOptional(token.OpenBracket) {
		types := p.parseTypeList()
		p.parseRequired(token.CloseBracket)
		return types
	}
	if p.current().Type != token.OpenBrace {
		return p.parseTypeList()
	}
	return nil
}

func parseFuncDeclaration(p *Parser) {

	keywords := p.parseKeywords(token.Func)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(token.Func)

	identifier := p.parseIdentifier()

	params := p.parseParameters()

	results := p.parseResults()

	body := p.parseBracesScope(ast.ExplicitVarDeclaration, ast.FuncDeclaration)

	node := ast.FuncDeclarationNode{
		Identifier: identifier,
		Parameters: params,
		Results:    results,
		Modifiers:  keywords,
		Body:       body,
	}

	p.scope.AddDeclaration(identifier, &node)
}

func parseLifecycleDeclaration(p *Parser) {

	keywords := p.parseKeywords(lexer.GetLifecycles()...)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	category := p.parseRequired(lexer.GetLifecycles()...)

	params := p.parseParameters()

	body := p.parseBracesScope(ast.ExplicitVarDeclaration, ast.FuncDeclaration)

	node := ast.LifecycleDeclarationNode{
		Modifiers:  keywords,
		Category:   category,
		Parameters: params,
		Body:       body,
	}

	if node.Parameters == nil {
		node.Parameters = params
	}

	p.scope.AddDeclaration("constructor", &node)
}

func parseTypeDeclaration(p *Parser) {

	keywords := p.parseKeywords(token.Identifier)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(token.Type)
	identifier := p.parseIdentifier()

	value := p.parseType()

	n := ast.TypeDeclarationNode{
		Modifiers:  keywords,
		Identifier: identifier,
		Value:      value,
	}

	p.scope.AddDeclaration(identifier, &n)
}

func (p *Parser) parseMapType() *ast.MapTypeNode {

	variable := p.parseOptional(token.Ellipsis)

	p.parseRequired(token.Map)
	p.parseRequired(token.OpenSquare)

	key := p.parseType()

	p.parseRequired(token.CloseSquare)

	value := p.parseType()

	mapType := ast.MapTypeNode{
		Key:      key,
		Value:    value,
		Variable: variable,
	}

	return &mapType
}

func (p *Parser) parseFuncType() *ast.FuncTypeNode {

	f := ast.FuncTypeNode{}

	f.Variable = p.parseOptional(token.Ellipsis)

	p.parseRequired(token.Func)

	if p.current().Type == token.Identifier {
		f.Identifier = p.parseIdentifier()
	}
	p.parseRequired(token.OpenBracket)

	f.Parameters = p.parseFuncTypeParameters()
	p.parseRequired(token.CloseBracket)

	if p.parseOptional(token.OpenBracket) {
		f.Results = p.parseFuncTypeParameters()
		p.parseRequired(token.CloseBracket)
	} else {
		f.Results = p.parseFuncTypeParameters()
	}

	return &f
}

func (p *Parser) parseFuncTypeParameters() []ast.Node {
	// can't mix named and unnamed
	var params []ast.Node
	if isExplicitVarDeclaration(p) {
		params = append(params, p.parseVarDeclaration())
		for p.parseOptional(token.Comma) {
			if isExplicitVarDeclaration(p) {
				params = append(params, p.parseVarDeclaration())
			} else if p.isNextAType() {
				p.addError(errMixedNamedParameters)
				p.parseType()
			} else {
				// TODO: add error
				p.next()
			}
		}
	} else {
		// must be types
		params = append(params, p.parseType())
		for p.parseOptional(token.Comma) {
			t := p.parseType()
			params = append(params, t)
			// TODO: handle errors
		}
	}
	return params
}

func (p *Parser) parseArrayType() *ast.ArrayTypeNode {

	variable := p.parseOptional(token.Ellipsis)

	p.parseRequired(token.OpenSquare)

	var max int

	if p.nextTokens(token.Integer) {
		i, err := strconv.ParseInt(p.current().String(), 10, 64)
		if err != nil {
			p.addError(errInvalidArraySize)
		}
		max = int(i)
		p.next()
	}

	p.parseRequired(token.CloseSquare)

	typ := p.parseType()

	return &ast.ArrayTypeNode{
		Value:    typ,
		Variable: variable,
		Length:   max,
	}
}

func parseExplicitVarDeclaration(p *Parser) {

	// parse variable Names
	node := p.parseVarDeclaration()

	// surely needs to be a declaration in some contexts
	switch p.scope.Type() {
	case ast.FuncDeclaration, ast.LifecycleDeclaration:
		p.scope.AddSequential(node)
		break
	default:
		for _, n := range node.Identifiers {
			p.scope.AddDeclaration(n, node)
		}
	}

}

func parseEventDeclaration(p *Parser) {
	p.parseRequired(token.Event)

	name := p.parseIdentifier()

	var types = p.parseParameters()

	node := ast.EventDeclarationNode{
		Identifier: name,
		Parameters: types,
	}
	p.scope.AddDeclaration(name, &node)
}
