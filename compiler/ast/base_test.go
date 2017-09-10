package ast

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestDeclareNode(t *testing.T) {
	scope := ScopeNode{}
	scope.Declare("interface", new(InterfaceDeclarationNode))
	goutil.Assert(t, len(scope.Nodes("interface")) == 1, "wrong interface length expect 1")
	scope.Declare("interface", new(InterfaceDeclarationNode))
	goutil.Assert(t, len(scope.Nodes("interface")) == 2, "wrong interface length expect 2")
}
