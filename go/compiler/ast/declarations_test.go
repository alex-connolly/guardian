package ast

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestContractValidation(t *testing.T) {
	n := new(ContractDeclarationNode)
	goutil.Assert(t, n.Validate(FuncDeclaration), "functions should be allowed")
	goutil.Assert(t, n.Validate(InterfaceDeclaration), "interfaces should be allowed")
	goutil.Assert(t, n.Validate(ClassDeclaration), "classes should be allowed")
	goutil.Assert(t, !n.Validate(ContractDeclaration), "contracts should not be allowed")
}

func TestInterfaceValidation(t *testing.T) {
	n := new(InterfaceDeclarationNode)
	goutil.Assert(t, n.Validate(FuncDeclaration), "functions should be allowed")
	goutil.Assert(t, !n.Validate(InterfaceDeclaration), "interfaces should not be allowed")
	goutil.Assert(t, !n.Validate(ClassDeclaration), "classes should not be allowed")
	goutil.Assert(t, !n.Validate(ContractDeclaration), "contracts should not be allowed")
}

func TestClassValidation(t *testing.T) {
	n := new(ClassDeclarationNode)
	goutil.Assert(t, n.Validate(FuncDeclaration), "functions should be allowed")
	goutil.Assert(t, !n.Validate(InterfaceDeclaration), "interfaces should not be allowed")
	goutil.Assert(t, n.Validate(ClassDeclaration), "classes not be allowed")
	goutil.Assert(t, !n.Validate(ContractDeclaration), "contracts should not be allowed")
}

func TestFunctionValidation(t *testing.T) {
	n := new(FuncDeclarationNode)
	goutil.Assert(t, n.Validate(FuncDeclaration), "functions should be allowed")
	goutil.Assert(t, !n.Validate(InterfaceDeclaration), "interfaces should not be allowed")
	goutil.Assert(t, n.Validate(ClassDeclaration), "classes not be allowed")
	goutil.Assert(t, !n.Validate(ContractDeclaration), "contracts should not be allowed")
}
