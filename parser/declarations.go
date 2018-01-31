package parser

import (
	"strconv"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
)

func parseInterfaceDeclaration(p *Parser) {

	p.parseGroupable(token.Interface, func(p *Parser) {

		start := p.getCurrentTokenLocation()

		identifier := p.parseIdentifier()

		generics := p.parsePossibleGenerics()

		var inherits []*ast.PlainTypeNode

		if p.parseOptional(token.Inherits) {
			inherits = p.parsePlainTypeList()
		}

		signatures := p.parseInterfaceSignatures()

		node := ast.InterfaceDeclarationNode{
			Begin:      start,
			Final:      p.getLastTokenLocation(),
			Identifier: identifier,
			Generics:   generics,
			Supers:     inherits,
			Modifiers:  p.getModifiers(),
			Signatures: signatures,
		}

		p.scope.AddDeclaration(identifier, &node)
	})

}

func (p *Parser) parseInterfaceFuncSignature() *ast.FuncTypeNode {

	f := new(ast.FuncTypeNode)

	p.ignoreNewLines()

	if !p.isNextToken(token.Identifier) {
		return nil
	}

	names := make([]string, 0)
	names = append(names, p.parseIdentifier())
	for !p.parseOptional(token.OpenBracket) {
		if p.isNextToken(token.Identifier) {
			names = append(names, p.parseIdentifier())
		} else {
			// knowingly error
			p.parseRequired(token.Identifier)
		}

	}

	f.Identifier = names[len(names)-1]
	names = names[:len(names)-1]
	for _, n := range names {
		f.Mods.AddModifier(n)
	}

	//p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		f.Parameters = p.parseFuncTypeParameters()
		p.parseRequired(token.CloseBracket)
	}

	if p.parseOptional(token.OpenBracket) {
		f.Results = p.parseFuncTypeParameters()
		p.parseRequired(token.CloseBracket)
	} else if !p.isNextToken(token.OpenBrace, token.NewLine) && p.hasTokens(1) {
		f.Results = p.parseFuncTypeParameters()
	}

	p.parseOptional(token.Semicolon)

	return f
}

func (p *Parser) parseInterfaceSignatures() []*ast.FuncTypeNode {

	p.parseRequired(token.OpenBrace)

	p.ignoreNewLines()

	var sigs []*ast.FuncTypeNode

	if p.parseOptional(token.CloseBrace) {
		return sigs
	}

	sigs = append(sigs, p.parseInterfaceFuncSignature())
	p.ignoreNewLines()
	for !p.parseOptional(token.CloseBrace) {
		sig := p.parseInterfaceFuncSignature()
		if sig != nil {
			sigs = append(sigs, sig)
		} else {
			p.addError(p.getCurrentTokenLocation(), errInvalidInterfaceProperty)
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
			} else if p.current().Type == token.CloseBrace {
				p.parseRequired(token.CloseBrace)
				return enums
			} else {
				p.addError(p.getCurrentTokenLocation(), errInvalidEnumProperty)
			}
		}

		p.ignoreNewLines()

		p.parseRequired(token.CloseBrace)
	}
	return enums
}

func parseEnumDeclaration(p *Parser) {

	p.parseGroupable(token.Enum, func(p *Parser) {

		start := p.getCurrentTokenLocation()

		identifier := p.parseIdentifier()

		var inherits []*ast.PlainTypeNode

		if p.parseOptional(token.Inherits) {
			inherits = p.parsePlainTypeList()
		}

		enums := p.parseEnumBody()

		node := ast.EnumDeclarationNode{
			Begin:      start,
			Final:      p.getLastTokenLocation(),
			Modifiers:  p.getModifiers(),
			Identifier: identifier,
			Inherits:   inherits,
			Enums:      enums,
		}

		p.scope.AddDeclaration(identifier, &node)
	})

}

func (p *Parser) parsePlainType() *ast.PlainTypeNode {

	start := p.getCurrentTokenLocation()

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
		if p.isNextToken(token.Shr) {
			p.spliceTokens(token.Gtr, token.Gtr)
		}
		p.parseRequired(token.Gtr)
	}

	return &ast.PlainTypeNode{
		Begin:      start,
		Final:      p.getLastTokenLocation(),
		Names:      names,
		Parameters: params,
		Variable:   variable,
	}
}

func (p *Parser) spliceTokens(types ...token.Type) {
	// delete the token
	tok := p.current()
	p.lexer.Tokens = append(p.lexer.Tokens[:p.index], p.lexer.Tokens[p.index+1:]...)
	// insert new tokens with all those types
	for i, t := range types {
		p.lexer.Tokens = append(p.lexer.Tokens, token.Token{})
		copy(p.lexer.Tokens[p.index+1:], p.lexer.Tokens[p.index:])
		p.lexer.Tokens[p.index] = token.Token{
			Start: tok.Start,
			End:   tok.End,
			Type:  t,
		}
		p.lexer.Tokens[p.index].Start.Offset += 1 + uint(i)
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

	start := p.getCurrentTokenLocation()

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

	body := p.parseBracesScope()

	node := ast.ClassDeclarationNode{
		Begin:      start,
		Final:      p.getLastTokenLocation(),
		Identifier: identifier,
		Generics:   generics,
		Supers:     inherits,
		Interfaces: interfaces,
		Modifiers:  p.getModifiers(),
		Body:       body,
	}

	p.scope.AddDeclaration(identifier, &node)
}

func parseClassDeclaration(p *Parser) {
	p.parseGroupable(token.Class, parseClassBody)
}

func parseContractBody(p *Parser) {
	start := p.current().Start
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

	valids := []ast.NodeType{
		ast.ClassDeclaration, ast.InterfaceDeclaration,
		ast.EventDeclaration, ast.ExplicitVarDeclaration,
		ast.TypeDeclaration, ast.EnumDeclaration,
		ast.LifecycleDeclaration, ast.FuncDeclaration,
	}

	body := p.parseBracesScope(valids...)

	node := ast.ContractDeclarationNode{
		Begin:      start,
		Final:      p.getLastTokenLocation(),
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
			parseIgnored(p)
			start := p.index
			declarator(p)
			if p.index == start {
				// prevent infinite loop
				p.next()
			}
			parseIgnored(p)
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
	start := p.current().Start
	names = append(names, p.parseIdentifier())
	for p.parseOptional(token.Comma) {
		names = append(names, p.parseIdentifier())
	}
	dType := p.parseType()
	e := &ast.ExplicitVarDeclarationNode{
		Begin:        start,
		Final:        p.getLastTokenLocation(),
		Modifiers:    p.getModifiers(),
		DeclaredType: dType,
		Identifiers:  names,
	}
	return e
}

func (p *Parser) parseIndividualParameter() *ast.ExplicitVarDeclarationNode {
	// TODO: is there an easier way of doing this

	start := p.getCurrentTokenLocation()

	// a, b, c string
	// indexed a string, b string
	// indexed a string,
	// NOTE: cannot declared multiple ids with the same params
	// indexed a, b string ==> TWO SEPARATE TYPES

	// indexed a string, b int
	// indexed, a string
	// indexed, a string, b int
	var possibleMods, names []string
	var dType ast.Node
	possibleMods = append(possibleMods, p.parseIdentifier())
	for p.isNextToken(token.Identifier) {
		possibleMods = append(possibleMods, p.parseIdentifier())
	}
	switch {
	case p.isNextToken(token.Comma):
		// last one was a type: x string,
		// or a var: a, b string
		if len(possibleMods) == 1 {
			names = append(names, possibleMods[0])
			p.parseRequired(token.Comma)

			names = append(names, p.parseIdentifier())
			for p.parseOptional(token.Comma) {
				names = append(names, p.parseIdentifier())
			}
			possibleMods = nil
		} else {
			names = possibleMods[len(possibleMods)-1:]
			possibleMods = possibleMods[:len(possibleMods)-2]
			p.index -= 1
		}
		dType = p.parseType()
		break
	case p.isNextToken(token.CloseBracket, token.Dot):

		if len(possibleMods) < 2 {

		} else {
			// last was a type/start of a type
			// one before that was a parameter name
			names = append(names, possibleMods[len(possibleMods)-2])
			possibleMods = possibleMods[:len(possibleMods)-2]
			p.index -= 1
			dType = p.parseType()
		}
		break
	default:
		// reached a type - can't have been a comma so last must have been var
		names = append(names, possibleMods[len(possibleMods)-1])
		possibleMods = possibleMods[:len(possibleMods)-1]
		dType = p.parseType()
		break
	}
	mods := p.getModifiers()
	for _, m := range possibleMods {
		mods.AddModifier(m)
	}
	e := &ast.ExplicitVarDeclarationNode{
		Begin:        start,
		Final:        p.getLastTokenLocation(),
		Modifiers:    mods,
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
		params = append(params, p.parseIndividualParameter())
		for p.parseOptional(token.Comma) {
			p.ignoreNewLines()
			params = append(params, p.parseIndividualParameter())
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

	f := new(ast.FuncTypeNode)

	if p.parseOptional(token.Func) {
		return nil
	}

	f.Generics = p.parsePossibleGenerics()

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
	} else if !p.isNextToken(token.OpenBrace, token.NewLine) && p.hasTokens(1) {
		f.Results = p.parseFuncTypeParameters()
	}

	return f
}

func parseFuncDeclaration(p *Parser) {

	start := p.getCurrentTokenLocation()

	p.parseRequired(token.Func)

	generics := p.parsePossibleGenerics()

	signature := p.parseFuncSignature()

	mods := p.getModifiers()

	var body *ast.ScopeNode
	if mods.Annotation("Builtin") == nil {
		// Builtins don't need bodies
		body = p.parseBracesScope(ast.ExplicitVarDeclaration, ast.FuncDeclaration)
	}

	node := ast.FuncDeclarationNode{
		Begin:     start,
		Final:     p.getLastTokenLocation(),
		Signature: signature,
		Generics:  generics,
		Modifiers: p.getModifiers(),
		Body:      body,
	}

	if signature != nil {
		p.scope.AddDeclaration(signature.Identifier, &node)
	}
}

func parseLifecycleDeclaration(p *Parser) {

	start := p.getCurrentTokenLocation()

	category := p.parseRequired(token.GetLifecycles()...)

	params := p.parseParameters()

	body := p.parseBracesScope(ast.ExplicitVarDeclaration, ast.FuncDeclaration)

	node := ast.LifecycleDeclarationNode{
		Begin:      start,
		Final:      p.getLastTokenLocation(),
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

		start := p.getCurrentTokenLocation()
		identifier := p.parseIdentifier()

		value := p.parseType()

		n := ast.TypeDeclarationNode{
			Begin:      start,
			Final:      p.getLastTokenLocation(),
			Modifiers:  p.getModifiers(),
			Identifier: identifier,
			Value:      value,
		}

		p.scope.AddDeclaration(identifier, &n)
	})
}

func (p *Parser) parseMapType() *ast.MapTypeNode {

	start := p.getCurrentTokenLocation()

	variable := p.parseOptional(token.Ellipsis)

	p.parseRequired(token.Map)
	p.parseRequired(token.OpenSquare)

	key := p.parseType()

	p.parseRequired(token.CloseSquare)

	value := p.parseType()

	mapType := ast.MapTypeNode{
		Begin:    start,
		Final:    p.getLastTokenLocation(),
		Key:      key,
		Value:    value,
		Variable: variable,
	}

	return &mapType
}

func (p *Parser) parseFuncType() *ast.FuncTypeNode {

	f := new(ast.FuncTypeNode)

	f.Begin = p.getCurrentTokenLocation()

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

	f.Final = p.getLastTokenLocation()

	return f
}

func (p *Parser) isNamedParameter() bool {
	return p.preserveState(func(p *Parser) bool {
		if !p.parseOptional(token.Identifier) {
			return false
		}
		for p.parseOptional(token.Comma) {
			if !p.parseOptional(token.Identifier) {
				return false
			}
		}
		return p.isNextAType()
	})
}

func (p *Parser) parseFuncTypeParameters() []ast.Node {
	// can't mix named and unnamed
	var params []ast.Node
	if p.isNamedParameter() {
		params = append(params, p.parseVarDeclaration())
		for p.parseOptional(token.Comma) {
			if p.isNamedParameter() {
				params = append(params, p.parseVarDeclaration())
			} else if p.isNextAType() {
				p.addError(p.getCurrentTokenLocation(), errMixedNamedParameters)
				p.parseType()
			} else {
				// TODO: add error
				p.next()
			}
		}
	} else if !p.isNextToken(token.CloseBracket) {
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

	start := p.getCurrentTokenLocation()

	variable := p.parseOptional(token.Ellipsis)

	p.parseRequired(token.OpenSquare)

	var max int

	if !p.parseOptional(token.CloseSquare) {
		if p.nextTokens(token.Integer) {
			i, err := strconv.ParseInt(p.current().String(p.lexer), 10, 64)
			if err != nil {
				p.addError(p.getCurrentTokenLocation(), errInvalidArraySize)
			}
			max = int(i)
			p.next()
		} else {
			p.addError(p.getCurrentTokenLocation(), errInvalidArraySize)
			p.next()
		}
		p.parseRequired(token.CloseSquare)
	} else {
		// no length specified
		variable = true
	}

	typ := p.parseType()

	return &ast.ArrayTypeNode{
		Begin:    start,
		Final:    p.getLastTokenLocation(),
		Value:    typ,
		Variable: variable,
		Length:   max,
	}
}

func (p *Parser) parseOptionallyTypedVarDeclaration() *ast.ExplicitVarDeclarationNode {
	// can be a simple list of identifiers

	var names []string
	names = append(names, p.parseIdentifier())
	for p.parseOptional(token.Comma) {
		names = append(names, p.parseIdentifier())
	}
	if p.isNextToken(token.Assign) {
		return &ast.ExplicitVarDeclarationNode{
			Modifiers:    p.getModifiers(),
			DeclaredType: nil,
			Identifiers:  names,
		}
	}
	dType := p.parseType()
	e := &ast.ExplicitVarDeclarationNode{
		Modifiers:    p.getModifiers(),
		DeclaredType: dType,
		Identifiers:  names,
	}
	return e
}

func processVarDeclaration(constant bool) func(p *Parser) {
	return func(p *Parser) {
		start := p.getCurrentTokenLocation()

		e := p.parseOptionallyTypedVarDeclaration()

		e.Begin = start

		e.IsConstant = constant

		if p.parseOptional(token.Assign) {
			e.Value = p.parseExpression()
		}

		p.parseOptional(token.Semicolon)

		e.Final = p.getLastTokenLocation()

		if e.IsConstant && e.Value == nil {
			p.addError(p.getCurrentTokenLocation(), errConstantWithoutValue)
		}

		switch p.scope.Type() {
		case ast.FuncDeclaration, ast.LifecycleDeclaration:
			p.scope.AddSequential(e)
			break
		default:
			for _, n := range e.Identifiers {
				p.scope.AddDeclaration(n, e)
			}
		}
	}
}

func parseExplicitVarDeclaration(p *Parser) {
	if p.isNextToken(token.Var) {
		p.parseGroupable(token.Var, processVarDeclaration(false))
	} else {
		p.parseGroupable(token.Const, processVarDeclaration(true))
	}
}

func parseEventDeclaration(p *Parser) {

	start := p.getCurrentTokenLocation()
	p.parseRequired(token.Event)

	generics := p.parsePossibleGenerics()

	name := p.parseIdentifier()

	var types = p.parseParameters()

	node := ast.EventDeclarationNode{
		Begin:      start,
		Final:      p.getLastTokenLocation(),
		Modifiers:  p.getModifiers(),
		Identifier: name,
		Generics:   generics,
		Parameters: types,
	}
	p.scope.AddDeclaration(name, &node)
}
