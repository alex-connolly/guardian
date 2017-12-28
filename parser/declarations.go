package parser

import (
	"strconv"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
)

func parseInterfaceDeclaration(p *Parser) {

	p.parseGroupable(token.Interface, func(p *Parser) {

		identifier := p.parseIdentifier()

		generics := p.parsePossibleGenerics()

		var inherits []*ast.PlainTypeNode

		if p.parseOptional(token.Inherits) {
			inherits = p.parsePlainTypeList()
		}

		signatures := p.parseInterfaceSignatures()

		node := ast.InterfaceDeclarationNode{
			Identifier: identifier,
			Generics:   generics,
			Supers:     inherits,
			Modifiers:  p.getModifiers(),
			Signatures: signatures,
		}

		p.scope.AddDeclaration(identifier, &node)
	})

}

func (p *Parser) parseInterfaceSignatures() []*ast.FuncTypeNode {

	p.parseRequired(token.OpenBrace)

	p.ignoreNewLines()

	var sigs []*ast.FuncTypeNode

	if p.parseOptional(token.CloseBrace) {
		return sigs
	}
	sigs = append(sigs, p.parseFuncSignature())
	p.ignoreNewLines()

	for !p.parseOptional(token.CloseBrace) {
		sig := p.parseFuncSignature()
		if sig != nil {
			sigs = append(sigs, sig)
		} else {
			p.addError(errInvalidInterfaceProperty)
			//p.parseConstruct()
			p.next()
		}
		p.ignoreNewLines()

	}
	return sigs
}

func (p *Parser) parseEnumBody() []string {

	p.parseRequired(token.OpenBrace)

	var enums []string
	// remove all new lines before the fist identifier
	p.ignoreNewLines()

	if !p.parseOptional(token.CloseBrace) {
		// can only contain identifiers split by commas and newlines
		first := p.parseIdentifier()
		enums = append(enums, first)
		for p.parseOptional(token.Comma) {

			p.ignoreNewLines()

			if p.current().Type == token.Identifier {
				enums = append(enums, p.parseIdentifier())
			} else {
				p.addError(errInvalidEnumProperty)
			}
		}

		p.ignoreNewLines()

		p.parseRequired(token.CloseBrace)
	}
	return enums
}

func parseEnumDeclaration(p *Parser) {

	p.parseGroupable(token.Enum, func(p *Parser) {

		identifier := p.parseIdentifier()

		var inherits []*ast.PlainTypeNode

		if p.parseOptional(token.Inherits) {
			inherits = p.parsePlainTypeList()
		}

		p.errorUnless(token.OpenBrace)

		enums := p.parseEnumBody()

		node := ast.EnumDeclarationNode{
			Modifiers:  p.getModifiers(),
			Identifier: identifier,
			Inherits:   inherits,
			Enums:      enums,
		}

		p.scope.AddDeclaration(identifier, &node)
	})

}

func (p *Parser) parsePlainType() *ast.PlainTypeNode {

	variable := p.parseOptional(token.Ellipsis)

	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(token.Dot) {
		names = append(names, p.parseIdentifier())
	}

	var params []ast.Node
	if p.parseOptional(token.Lss) {
		params = append(params, p.parseType())
		for p.parseOptional(token.Or) {
			params = append(params, p.parseType())
		}
		p.parseRequired(token.Gtr)
	}

	return &ast.PlainTypeNode{
		Names:      names,
		Parameters: params,
		Variable:   variable,
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

func parseClassBody(p *Parser) {

	identifier := p.parseIdentifier()

	generics := p.parsePossibleGenerics()
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

	p.errorUnless(token.OpenBrace)

	body := p.parseBracesScope()

	node := ast.ClassDeclarationNode{
		Identifier: identifier,
		Generics:   generics,
		Supers:     inherits,
		Interfaces: interfaces,
		Modifiers:  p.getModifiers(),
		Body:       body,
	}

	p.scope.AddDeclaration(identifier, &node)
}

func (p *Parser) errorUnless(t token.Type) {
	for !p.isNextToken(t) {
		p.addError("yoyo")
		p.next()
	}
}

func parseClassDeclaration(p *Parser) {
	p.parseGroupable(token.Class, parseClassBody)
}

func parseContractBody(p *Parser) {
	identifier := p.parseIdentifier()

	generics := p.parsePossibleGenerics()

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

	p.errorUnless(token.OpenBrace)

	valids := []ast.NodeType{
		ast.ClassDeclaration, ast.InterfaceDeclaration,
		ast.EventDeclaration, ast.ExplicitVarDeclaration,
		ast.TypeDeclaration, ast.EnumDeclaration,
		ast.LifecycleDeclaration, ast.FuncDeclaration,
	}

	body := p.parseBracesScope(valids...)

	node := ast.ContractDeclarationNode{
		Identifier: identifier,
		Generics:   generics,
		Supers:     inherits,
		Interfaces: interfaces,
		Modifiers:  p.getModifiers(),
		Body:       body,
	}

	p.scope.AddDeclaration(identifier, &node)
}

func parseContractDeclaration(p *Parser) {
	p.parseGroupable(token.Contract, parseContractBody)
}

func (p *Parser) parseGroupable(id token.Type, declarator func(*Parser)) {
	p.parseRequired(id)
	if p.parseOptional(token.OpenBracket) {
		for !p.parseOptional(token.CloseBracket) {
			p.ignoreNewLines()
			declarator(p)
			p.ignoreNewLines()
		}
	} else {
		declarator(p)
	}
}

func (p *Parser) ignoreNewLines() {
	for isNewLine(p) {
		parseNewLine(p)
	}
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
	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(token.Comma) {
		names = append(names, p.parseIdentifier())
	}
	dType := p.parseType()
	e := &ast.ExplicitVarDeclarationNode{
		Modifiers:    p.getModifiers(),
		DeclaredType: dType,
		Identifiers:  names,
	}
	return e
}

func (p *Parser) parseParameters() []*ast.ExplicitVarDeclarationNode {
	var params []*ast.ExplicitVarDeclarationNode
	p.parseRequired(token.OpenBracket)
	p.ignoreNewLines()
	if !p.parseOptional(token.CloseBracket) {
		params = append(params, p.parseVarDeclaration())
		for p.parseOptional(token.Comma) {
			p.ignoreNewLines()
			params = append(params, p.parseVarDeclaration())
		}
		p.ignoreNewLines()
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
	if p.hasTokens(1) {
		if p.current().Type != token.OpenBrace {
			return p.parseTypeList()
		}
	}
	return nil
}

func (p *Parser) parseFuncSignature() *ast.FuncTypeNode {

	// return nil for any errors
	// TODO: maybe do this better
	f := new(ast.FuncTypeNode)

	if p.isNextToken(token.Func) {
		p.addError("can't")
		p.next()
	}

	if !p.isNextToken(token.Identifier) {
		return nil
	}

	f.Identifier = p.parseIdentifier()

	p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		f.Parameters = p.parseFuncTypeParameters()
		p.parseRequired(token.CloseBracket)
	}

	if p.parseOptional(token.OpenBracket) {
		f.Results = p.parseFuncTypeParameters()
		p.parseRequired(token.CloseBracket)
	} else {
		f.Results = p.parseFuncTypeParameters()
	}

	return f
}

func parseFuncDeclaration(p *Parser) {

	p.parseRequired(token.Func)

	generics := p.parsePossibleGenerics()

	signature := p.parseFuncSignature()

	p.errorUnless(token.OpenBrace)

	body := p.parseBracesScope(ast.ExplicitVarDeclaration, ast.FuncDeclaration)

	node := ast.FuncDeclarationNode{
		Signature: signature,
		Generics:  generics,
		Modifiers: p.getModifiers(),
		Body:      body,
	}

	p.scope.AddDeclaration(signature.Identifier, &node)
}

func parseLifecycleDeclaration(p *Parser) {

	category := p.parseRequired(token.GetLifecycles()...)

	params := p.parseParameters()

	body := p.parseBracesScope(ast.ExplicitVarDeclaration, ast.FuncDeclaration)

	node := ast.LifecycleDeclarationNode{
		Modifiers:  p.getModifiers(),
		Category:   category,
		Parameters: params,
		Body:       body,
	}

	if node.Parameters == nil {
		node.Parameters = params
	}

	p.scope.AddDeclaration("lifecycle", &node)
}

func parseTypeDeclaration(p *Parser) {

	p.parseGroupable(token.KWType, func(p *Parser) {
		identifier := p.parseIdentifier()

		value := p.parseType()

		n := ast.TypeDeclarationNode{
			Modifiers:  p.getModifiers(),
			Identifier: identifier,
			Value:      value,
		}

		p.scope.AddDeclaration(identifier, &n)
	})
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

	p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		f.Parameters = p.parseFuncTypeParameters()
		p.parseRequired(token.CloseBracket)
	}

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
		t := p.parseType()
		params = append(params, t)
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

	if p.isNextToken(token.Var) {
		p.parseGroupable(token.Var, func(p *Parser) {

			e := p.parseVarDeclaration()

			e.isConstant = false

			switch p.scope.Type() {
			case ast.FuncDeclaration, ast.LifecycleDeclaration:
				p.scope.AddSequential(e)
				break
			default:
				for _, n := range e.Identifiers {
					p.scope.AddDeclaration(n, e)
				}
			}
		})
	} else {
		p.parseGroupable(token.Const, func(p *Parser) {

			e := p.parseVarDeclaration()

			e.isConstant = true

			switch p.scope.Type() {
			case ast.FuncDeclaration, ast.LifecycleDeclaration:
				p.scope.AddSequential(e)
				break
			default:
				for _, n := range e.Identifiers {
					p.scope.AddDeclaration(n, e)
				}
			}
		})
	}

}

func parseEventDeclaration(p *Parser) {
	p.parseRequired(token.Event)

	generics := p.parsePossibleGenerics()

	name := p.parseIdentifier()

	var types = p.parseParameters()

	node := ast.EventDeclarationNode{
		Modifiers:  p.getModifiers(),
		Identifier: name,
		Generics:   generics,
		Parameters: types,
	}
	p.scope.AddDeclaration(name, &node)
}
