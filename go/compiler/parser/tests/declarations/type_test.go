package declarations

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestParseTypeDeclaration(t *testing.T) {
	p := parser.ParseString("type Bowl int")
	util.Assert(t, p.Scope.Type() == ast.TypeDeclaration, "not a type declaration")
}

func TestParseTypeDeclarationImported(t *testing.T) {
	p := parser.ParseString("type Bowl pkg.Object")
	util.Assert(t, p.Scope.Type() == ast.TypeDeclaration, "not a type declaration")
}
