package parser

import (
	"strconv"

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

	keywords := p.parseKeywords(lexer.TknInterface)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(lexer.TknInterface)
	identifier := p.parseIdentifier()

	var inherits []*ast.PlainTypeNode

	if p.parseOptional(lexer.TknInherits) {
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

	p.parseRequired(lexer.TknOpenBrace)

	for isNewLine(p) {
		parseNewLine(p)
	}

	var sigs []*ast.FuncTypeNode

	first := p.parseFuncType()
	sigs = append(sigs, first)

	for isNewLine(p) {
		parseNewLine(p)
	}

	for !p.parseOptional(lexer.TknCloseBrace) {

		if !p.isFuncType() {
			p.addError("Everything in an interface must be a func type")
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

	p.parseRequired(lexer.TknOpenBrace)
	var enums []string
	// remove all new lines before the fist identifier
	for isNewLine(p) {
		parseNewLine(p)
	}

	if !p.parseOptional(lexer.TknCloseBrace) {
		// can only contain identifiers split by commas and newlines
		first := p.parseIdentifier()
		enums = append(enums, first)
		for p.parseOptional(lexer.TknComma) {

			for isNewLine(p) {
				parseNewLine(p)
			}

			if p.current().Type == lexer.TknIdentifier {
				enums = append(enums, p.parseIdentifier())
			} else {
				p.addError("Everything in an enum must be an identifier")
			}
		}

		for isNewLine(p) {
			parseNewLine(p)
		}

		p.parseRequired(lexer.TknCloseBrace)
	}
	return enums
}

func parseEnumDeclaration(p *Parser) {

	keywords := p.parseKeywords(lexer.TknEnum)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(lexer.TknEnum)
	identifier := p.parseIdentifier()

	var inherits []*ast.PlainTypeNode

	if p.parseOptional(lexer.TknInherits) {
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

	variable := p.parseOptional(lexer.TknEllipsis)

	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(lexer.TknDot) {
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
	for p.parseOptional(lexer.TknComma) {
		refs = append(refs, p.parsePlainType())
	}
	return refs
}

func parseClassDeclaration(p *Parser) {

	keywords := p.parseKeywords(lexer.TknClass)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(lexer.TknClass)

	identifier := p.parseIdentifier()

	// is and inherits can be in any order

	var inherits, interfaces []*ast.PlainTypeNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parsePlainTypeList()
		if p.parseOptional(lexer.TknIs) {
			interfaces = p.parsePlainTypeList()
		}
	} else if p.parseOptional(lexer.TknIs) {
		interfaces = p.parsePlainTypeList()
		if p.parseOptional(lexer.TknInherits) {
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

	keywords := p.parseKeywords(lexer.TknContract)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(lexer.TknContract)
	identifier := p.parseIdentifier()

	// is and inherits can be in any order

	var inherits, interfaces []*ast.PlainTypeNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parsePlainTypeList()
		if p.parseOptional(lexer.TknIs) {
			interfaces = p.parsePlainTypeList()
		}
	} else if p.parseOptional(lexer.TknIs) {
		interfaces = p.parsePlainTypeList()
		if p.parseOptional(lexer.TknInherits) {
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
	case p.isNextToken(lexer.TknIdentifier):
		return p.parsePlainType()
	}
	return nil
}

func (p *Parser) parseVarDeclaration() *ast.ExplicitVarDeclarationNode {

	keywords := p.parseKeywords(lexer.TknIdentifier)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(lexer.TknComma) {
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

func (p *Parser) parseTypeList() []ast.Node {
	var types []ast.Node
	first := p.parseType()
	types = append(types, first)
	for p.parseOptional(lexer.TknComma) {
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

	keywords := p.parseKeywords(lexer.TknFunc)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(lexer.TknFunc)

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

	keywords := p.parseKeywords(lexer.TknIdentifier)

	if p.keywords != nil {
		keywords = append(keywords, p.keywords...)
	}

	p.parseRequired(lexer.TknType)
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

	variable := p.parseOptional(lexer.TknEllipsis)

	p.parseRequired(lexer.TknMap)
	p.parseRequired(lexer.TknOpenSquare)

	key := p.parseType()

	p.parseRequired(lexer.TknCloseSquare)

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

	f.Variable = p.parseOptional(lexer.TknEllipsis)

	p.parseRequired(lexer.TknFunc)

	if p.current().Type == lexer.TknIdentifier {
		f.Identifier = p.parseIdentifier()
	}
	p.parseRequired(lexer.TknOpenBracket)

	f.Parameters = p.parseFuncTypeParameters()
	p.parseRequired(lexer.TknCloseBracket)

	if p.parseOptional(lexer.TknOpenBracket) {
		f.Results = p.parseFuncTypeParameters()
		p.parseRequired(lexer.TknCloseBracket)
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
		for p.parseOptional(lexer.TknComma) {
			if isExplicitVarDeclaration(p) {
				params = append(params, p.parseVarDeclaration())
			} else if p.isNextAType() {
				p.addError("Mixed named and unnamed parameters")
				p.parseType()
			} else {
				// TODO: add error
				p.next()
			}
		}
	} else {
		// must be types
		params = append(params, p.parseType())
		for p.parseOptional(lexer.TknComma) {
			t := p.parseType()
			params = append(params, t)
			// TODO: handle errors
		}
	}
	return params
}

func (p *Parser) parseArrayType() *ast.ArrayTypeNode {

	variable := p.parseOptional(lexer.TknEllipsis)

	p.parseRequired(lexer.TknOpenSquare)

	var max int

	if p.nextTokens(lexer.TknInteger) {
		i, err := strconv.ParseInt(p.current().String(), 10, 64)
		if err != nil {
			p.addError("Invalid array size")
		}
		max = int(i)
		p.next()
	}

	p.parseRequired(lexer.TknCloseSquare)

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
	p.parseRequired(lexer.TknEvent)

	name := p.parseIdentifier()

	var types = p.parseParameters()

	node := ast.EventDeclarationNode{
		Identifier: name,
		Parameters: types,
	}
	p.scope.AddDeclaration(name, &node)
}
