package parser

import (
	"fmt"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

// Parser ...
type Parser struct {
	scope            *ast.ScopeNode
	Expression       ast.ExpressionNode
	modifiers        [][]string
	lastModifiers    []string
	annotations      []*typing.Annotation
	index            int
	errs             util.Errors
	line             int
	simple           bool
	seenCastOperator bool
	lexer            *lexer.Lexer
}

func createParser(data string) *Parser {
	p := new(Parser)
	p.line = 1
	p.seenCastOperator = false
	p.lexer = lexer.LexString(data)
	p.scope = &ast.ScopeNode{
		ValidTypes: []ast.NodeType{
			ast.InterfaceDeclaration, ast.ClassDeclaration,
			ast.FuncDeclaration,
		},
	}
	return p
}

func (p *Parser) current() token.Token {
	return p.token(0)
}

func (p *Parser) next() {
	p.index++
}

func (p *Parser) token(offset int) token.Token {
	return p.lexer.Tokens[p.index+offset]
}

// index 0, tknlength 1
// hasTokens(0) yes always
// hasTokens(1) yes
// hasTokens(2) no
func (p *Parser) hasTokens(offset int) bool {
	return p.index+offset <= len(p.lexer.Tokens)
}

func (p *Parser) parseOptional(types ...token.Type) bool {
	if !p.hasTokens(1) {
		return false
	}
	for _, t := range types {
		if p.current().Type == t {
			p.next()
			return true
		}
	}
	if !p.hasTokens(1) {
		return false
	}
	for _, t := range types {
		if p.current().Type == t {
			p.next()
			return true
		}
	}
	return false
}

// TODO: clarify what this actually returns
func (p *Parser) parseRequired(types ...token.Type) token.Type {
	if !p.hasTokens(1) {
		p.addError(p.getLastTokenLocation(), fmt.Sprintf(errRequiredType, listTypes(types), "nothing"))
		// TODO: what should be returned here
		return token.Invalid
	}
	for _, t := range types {
		if p.isNextToken(t) {
			p.next()
			return t
		}
	}
	p.ignoreComments()
	for _, t := range types {
		if p.isNextToken(t) {
			p.next()
			return t
		}
	}
	p.addError(p.getCurrentTokenLocation(), fmt.Sprintf(errRequiredType, listTypes(types), p.current().Name()))
	// correct return type
	return token.Invalid
}

func (p *Parser) ignoreComments() {
	for p.isNextToken(token.LineComment, token.CommentOpen) {
		if p.isNextToken(token.LineComment) {
			parseSingleLineComment(p)
		} else if p.isNextToken(token.CommentOpen) {
			parseMultiLineComment(p)
		}
	}
}

func listTypes(types []token.Type) string {
	if len(types) == 1 {
		return types[0].Name()
	}
	s := ""
	for _, t := range types {
		s += t.Name()
		s += ","
	}
	// remove trailing comma
	return s[:len(s)-1]
}

func (p *Parser) getLastTokenLocation() util.Location {
	if len(p.lexer.Tokens) == 0 {
		return util.Location{
			Filename: "file_name",
			Line:     0,
			Offset:   0,
		}
	}
	if p.index >= len(p.lexer.Tokens) {
		// return end of last available token
		return p.lexer.Tokens[len(p.lexer.Tokens)-1].End
	}
	return p.token(-1).End
}

func (p *Parser) getCurrentTokenLocation() util.Location {
	if p.index >= len(p.lexer.Tokens) {
		// return start of last available token
		return p.lexer.Tokens[len(p.lexer.Tokens)-1].Start
	}
	return p.current().Start
}

func (p *Parser) parseIdentifier() string {
	if !p.hasTokens(1) {
		p.addError(p.getLastTokenLocation(), fmt.Sprintf(errRequiredType, "identifier", "nothing"))
		return ""
	}
	if p.current().Type != token.Identifier {
		p.addError(p.getCurrentTokenLocation(), fmt.Sprintf(errRequiredType, "identifier", p.current().Name()))
		return ""
	}
	s := p.current().String(p.lexer)
	p.next()
	return s
}

func (p *Parser) validate(t ast.NodeType) {
	if p.scope != nil {
		if !p.scope.IsValid(t) {
			p.addError(p.getCurrentTokenLocation(), errInvalidScopeDeclaration)
		}
	}
}

func parseNewLine(p *Parser) {
	p.line++
	p.next()
}

func parseSingleLineComment(p *Parser) {
	p.parseRequired(token.LineComment)
	for p.hasTokens(1) {
		if p.current().Type != token.NewLine {
			p.next()
		} else {
			parseNewLine(p)
			return
		}

	}

}

func parseMultiLineComment(p *Parser) {
	p.parseRequired(token.CommentOpen)
	for p.hasTokens(1) && p.current().Type != token.CommentClose {
		if p.current().Type == token.NewLine {
			p.line++
		}
		p.next()
	}
	p.parseOptional(token.CommentClose)
}

func (p *Parser) parseModifierList() []string {
	var mods []string
	for p.hasTokens(1) && p.current().Type == token.Identifier {
		mods = append(mods, p.parseIdentifier())
	}
	return mods
}

func parseModifiers(p *Parser) {
	p.lastModifiers = p.parseModifierList()
}

func parseGroup(p *Parser) {
	if p.lastModifiers == nil {
		p.addError(p.getCurrentTokenLocation(), errEmptyGroup)
		p.next()
	} else {
		p.modifiers = append(p.modifiers, p.lastModifiers)
		p.lastModifiers = nil
		p.parseRequired(token.OpenBracket)
		for p.hasTokens(1) {
			if p.current().Type == token.CloseBracket {
				p.modifiers = p.modifiers[:len(p.modifiers)-1]
				p.parseRequired(token.CloseBracket)
				return
			}
			p.parseNextConstruct()
		}
		p.addError(p.getCurrentTokenLocation(), errUnclosedGroup)
	}
}

func (p *Parser) getModifiers() typing.Modifiers {
	var mods []string
	for _, m := range p.lastModifiers {
		mods = append(mods, m)
	}
	for _, ml := range p.modifiers {
		for _, m := range ml {
			mods = append(mods, m)
		}
	}
	return typing.Modifiers{
		Modifiers:   mods,
		Annotations: p.annotations,
	}
}

func (p *Parser) addError(loc util.Location, message string) {
	err := util.Error{
		Message:  message,
		Location: loc,
	}
	p.errs = append(p.errs, err)
}

func (p *Parser) parseBracesScope(valids ...ast.NodeType) *ast.ScopeNode {
	return p.parseEnclosedScope([]token.Type{token.OpenBrace}, []token.Type{token.CloseBrace}, true, valids...)
}

func (p *Parser) parseEnclosedScope(opener, closer []token.Type, parseLast bool, valids ...ast.NodeType) *ast.ScopeNode {
	p.parseRequired(opener...)
	scope := p.parseScope(closer, valids...)
	if parseLast {
		p.parseRequired(closer...)
	}
	return scope
}

func (p *Parser) parseScope(terminators []token.Type, valids ...ast.NodeType) *ast.ScopeNode {
	scope := new(ast.ScopeNode)
	scope.Parent = p.scope
	p.scope = scope
	for p.hasTokens(1) {
		for _, t := range terminators {
			if p.current().Type == t {
				p.scope = scope.Parent
				return scope
			}
		}
		p.parseNextConstruct()
	}
	return scope
}

func (p *Parser) parseNextConstruct() {
	found := false
	for _, c := range getPrimaryConstructs() {
		if c.is(p) {
			//fmt.Printf("FOUND: %s at index %d on line %d\n", c.name, p.index, p.line)
			c.parse(p)
			found = true
			break
		}
	}
	if !found {
		// try interpreting it as an expression
		saved := *p
		expr := p.parseExpression()
		if expr == nil {
			*p = saved
			//fmt.Printf("Unrecognised construct at index %d: %s\n", p.index, p.current().String(p.lexer))
			p.addError(p.getCurrentTokenLocation(), fmt.Sprintf("Unrecognised construct: %s", p.current().String(p.lexer)))
			p.next()
		} else {
			p.parsePossibleSequentialExpression(expr)
		}
	}
}

func (p *Parser) parsePossibleSequentialExpression(expr ast.ExpressionNode) {
	switch expr.Type() {
	case ast.CallExpression:
		//fmt.Println("call")
		p.scope.AddSequential(expr)
		p.parseOptional(token.Semicolon)
		return
	case ast.Reference:
		for r := expr; r != nil; {
			switch t := r.(type) {
			case *ast.CallExpressionNode:
				p.scope.AddSequential(expr)
				return
			case *ast.ReferenceNode:
				r = t.Reference
				break
			default:
				//fmt.Printf("dangling at index %d\n", p.index)
				p.addError(p.getCurrentTokenLocation(), errDanglingExpression)
				return
			}
		}
	}
	//fmt.Printf("dangling at index %d\n", p.index)
	p.addError(p.getCurrentTokenLocation(), errDanglingExpression)
}

func (p *Parser) parseString() string {
	s := p.current().TrimmedString(p.lexer)
	p.next()
	return s
}

func parseAnnotation(p *Parser) {
	a := new(typing.Annotation)
	p.parseRequired(token.At)

	a.Name = p.parseIdentifier()

	p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		var names []string
		if !p.isNextToken(token.String) {
			p.addError(p.getCurrentTokenLocation(), errInvalidAnnotationParameter)
			p.next()
		} else {
			names = append(names, p.parseString())
		}
		for p.parseOptional(token.Comma) {
			if !p.isNextToken(token.String) {
				p.addError(p.getCurrentTokenLocation(), errInvalidAnnotationParameter)
				p.next()
			} else {
				names = append(names, p.parseString())
			}
		}
		a.Parameters = names
		p.parseRequired(token.CloseBracket)
	}
	if p.annotations == nil {
		p.annotations = make([]*typing.Annotation, 0)
	}
	p.annotations = append(p.annotations, a)
}
