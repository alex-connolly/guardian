package ast

import (
	"axia/guardian/go/util"
	"testing"
)

func TestContractValidation(t *testing.T) {
	n := new(ContractDeclarationNode)
	util.Assert(t, n.Declare("function", new(FuncDeclarationNode)), "functions should be allowed")
	util.Assert(t, n.Declare("interface", new(InterfaceDeclarationNode)), "interfaces should be allowed")
	util.Assert(t, n.Declare("class", new(ClassDeclarationNode)), "classes should be allowed")
	util.Assert(t, !n.Declare("contract", new(ContractDeclarationNode)), "contracts should not be allowed")
}

func TestInterfaceValidation(t *testing.T) {
	n := new(InterfaceDeclarationNode)
	util.Assert(t, n.Declare("function", new(FuncDeclarationNode)), "functions should be allowed")
	util.Assert(t, !n.Declare("interface", new(InterfaceDeclarationNode)), "interfaces should not be allowed")
	util.Assert(t, !n.Declare("class", new(ClassDeclarationNode)), "classes should not be allowed")
	util.Assert(t, !n.Declare("contract", new(ContractDeclarationNode)), "contracts should not be allowed")
}

func TestClassValidation(t *testing.T) {
	n := new(InterfaceDeclarationNode)
	util.Assert(t, n.Declare("function", new(FuncDeclarationNode)), "functions should be allowed")
	util.Assert(t, !n.Declare("interface", new(InterfaceDeclarationNode)), "interfaces should not be allowed")
	util.Assert(t, n.Declare("class", new(ClassDeclarationNode)), "classes not be allowed")
	util.Assert(t, !n.Declare("contract", new(ContractDeclarationNode)), "contracts should not be allowed")
}

func TestFunctionValidation(t *testing.T) {
	n := new(InterfaceDeclarationNode)
	util.Assert(t, n.Declare("function", new(FuncDeclarationNode)), "functions should be allowed")
	util.Assert(t, !n.Declare("interface", new(InterfaceDeclarationNode)), "interfaces should not be allowed")
	util.Assert(t, n.Declare("class", new(ClassDeclarationNode)), "classes not be allowed")
	util.Assert(t, !n.Declare("contract", new(ContractDeclarationNode)), "contracts should not be allowed")
}
