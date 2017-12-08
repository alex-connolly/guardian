package evm

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/parser"
)

func TestTraverseExplicitVariableDeclaration(t *testing.T) {
	ast, errs := parser.ParseString(`name uint8`)
	goutil.AssertNow(t, errs == nil, "errs should be nil")
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e := NewVM()
	errs = validator.Validate(ast, e)
	goutil.Assert(t, errs == nil, errs.Format())
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.storage)))
}

func TestTraverseExplicitVariableDeclarationFunc(t *testing.T) {
	ast, errs := parser.ParseString(`name func(a, b string) int`)
	goutil.AssertNow(t, errs == nil, "errs should be nil")
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e := NewVM()
	errs = validator.Validate(ast, e)
	goutil.Assert(t, errs == nil, errs.Format())
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.storage)))
}

func TestTraverseExplicitVariableDeclarationFixedArray(t *testing.T) {
	ast, errs := parser.ParseString(`name [3]uint8`)
	goutil.AssertNow(t, errs == nil, "errs should be nil")
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e := NewVM()
	errs = validator.Validate(ast, e)
	goutil.Assert(t, errs == nil, errs.Format())
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.storage)))
}

func TestTraverseExplicitVariableDeclarationVariableArray(t *testing.T) {
	ast, errs := parser.ParseString(`name [3]uint8`)
	goutil.AssertNow(t, errs == nil, "errs should be nil")
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e := NewVM()
	errs = validator.Validate(ast, e)
	goutil.Assert(t, errs == nil, errs.Format())
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.storage)))
}

func TestTraverseTypeDeclaration(t *testing.T) {
	ast, errs := parser.ParseString(`type Dog int`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e := NewVM()
	errs = validator.Validate(ast, e)
	goutil.Assert(t, errs == nil, errs.Format())
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 0, fmt.Sprintf("allocate a block: %d", len(e.storage)))
}

func TestTraverseClassDeclaration(t *testing.T) {
	ast, errs := parser.ParseString(`class Dog {}`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e := NewVM()
	errs = validator.Validate(ast, e)
	goutil.Assert(t, errs == nil, errs.Format())
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 0, fmt.Sprintf("allocate a block: %d", len(e.storage)))
}

func TestTraverseInterfaceDeclaration(t *testing.T) {
	ast, errs := parser.ParseString(`interface Dog {}`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e := NewVM()
	errs = validator.Validate(ast, e)
	goutil.Assert(t, errs == nil, errs.Format())
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 0, fmt.Sprintf("allocate a block: %d", len(e.storage)))
}
