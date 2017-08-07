package declarations

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"testing"
)

func TestParseTypeDeclaration(t *testing.T) {
	p := parser.ParseString("type Bowl int")
	util.Assert(t, p.Scope.Type() == ast.TypeDeclaration, "not a type declaration")
}

func TestParseTypeDeclarationImported(t *testing.T) {
	p := parser.ParseString("type Bowl pkg.Object")
	util.Assert(t, p.Scope.Type() == ast.TypeDeclaration, "not a type declaration")
}
