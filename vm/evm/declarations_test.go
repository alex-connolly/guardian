package evm

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/goutil"
)

func TestTraverseExplicitVariableDeclaration(t *testing.T) {
	e := NewVM()
	ast, errs := validator.ValidateString(e, `var name uint8`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.storage)))
}

func TestTraverseExplicitVariableDeclarationFunc(t *testing.T) {
	e := NewVM()
	ast, errs := validator.ValidateString(e, `var name func(a, b string) int`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.storage)))
}

func TestTraverseExplicitVariableDeclarationFixedArray(t *testing.T) {
	e := NewVM()
	ast, errs := validator.ValidateString(e, `var name [3]uint8`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.storage)))
}

func TestTraverseExplicitVariableDeclarationVariableArray(t *testing.T) {
	e := NewVM()
	ast, errs := validator.ValidateString(e, `var name [3]uint8`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.storage)))
}

func TestTraverseTypeDeclaration(t *testing.T) {
	e := NewVM()
	ast, errs := validator.ValidateString(e, `type Dog int`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 0, fmt.Sprintf("allocate a block: %d", len(e.storage)))
}

func TestTraverseClassDeclaration(t *testing.T) {
	e := NewVM()
	ast, errs := validator.ValidateString(e, `class Dog {}`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 0, fmt.Sprintf("allocate a block: %d", len(e.storage)))
}

func TestTraverseInterfaceDeclaration(t *testing.T) {
	e := NewVM()
	ast, errs := validator.ValidateString(e, `interface Dog {}`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 0, fmt.Sprintf("allocate a block: %d", len(e.storage)))
}
