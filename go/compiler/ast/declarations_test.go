package ast

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestContractValidation(t *testing.T) {
	n := new(ContractDeclarationNode)
	goutil.Assert(t, n.Declare("function", new(FuncDeclarationNode)), "functions should be allowed")
	goutil.Assert(t, n.Declare("interface", new(InterfaceDeclarationNode)), "interfaces should be allowed")
	goutil.Assert(t, n.Declare("class", new(ClassDeclarationNode)), "classes should be allowed")
	goutil.Assert(t, !n.Declare("contract", new(ContractDeclarationNode)), "contracts should not be allowed")
}

func TestInterfaceValidation(t *testing.T) {
	n := new(InterfaceDeclarationNode)
	goutil.Assert(t, n.Declare("function", new(FuncDeclarationNode)), "functions should be allowed")
	goutil.Assert(t, !n.Declare("interface", new(InterfaceDeclarationNode)), "interfaces should not be allowed")
	goutil.Assert(t, !n.Declare("class", new(ClassDeclarationNode)), "classes should not be allowed")
	goutil.Assert(t, !n.Declare("contract", new(ContractDeclarationNode)), "contracts should not be allowed")
}

func TestClassValidation(t *testing.T) {
	n := new(InterfaceDeclarationNode)
	goutil.Assert(t, n.Declare("function", new(FuncDeclarationNode)), "functions should be allowed")
	goutil.Assert(t, !n.Declare("interface", new(InterfaceDeclarationNode)), "interfaces should not be allowed")
	goutil.Assert(t, n.Declare("class", new(ClassDeclarationNode)), "classes not be allowed")
	goutil.Assert(t, !n.Declare("contract", new(ContractDeclarationNode)), "contracts should not be allowed")
}

func TestFunctionValidation(t *testing.T) {
	n := new(InterfaceDeclarationNode)
	goutil.Assert(t, n.Declare("function", new(FuncDeclarationNode)), "functions should be allowed")
	goutil.Assert(t, !n.Declare("interface", new(InterfaceDeclarationNode)), "interfaces should not be allowed")
	goutil.Assert(t, n.Declare("class", new(ClassDeclarationNode)), "classes not be allowed")
	goutil.Assert(t, !n.Declare("contract", new(ContractDeclarationNode)), "contracts should not be allowed")
}
