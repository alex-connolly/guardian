package declarations

import "testing"

func TestParseTypeDeclaration(t *testing.T) {
	ParseString("type Bowl int")
	//util.Assert(t, p.scope.Type() == , "not a type declaration")
}
