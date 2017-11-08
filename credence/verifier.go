package credence

import "github.com/end-r/guardian/compiler/ast"

// VerifyInvariant ...
func VerifyInvariant() {

}

// Verifier ...
type Verifier struct {
	Invariants []Invariant
}

// Invariant ...
type Invariant struct {
	Condition func(ast.Node) bool
}
