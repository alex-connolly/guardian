package validator

import (
	"testing"

	"github.com/end-r/guardian/compiler/ast"
)

func TestValidateAssignment(t *testing.T) {
	v := NewValidator()
	node := new(ast.AssignmentStatementNode)
	v.validateAssignment(node)
}
