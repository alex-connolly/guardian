package validator

import "github.com/end-r/guardian/compiler/ast"

func (v *Validator) addScanned(name string, t Type) {
	if v.scope.scanned == nil {
		v.scope.scanned = make(map[string]Type)
	}
	v.scope.scanned[name] = t
}

func (v *Validator) scanDeclaration(node ast.Node) {
	switch n := node.(type) {
	case ast.ClassDeclarationNode:
		v.scanClassDeclaration(n)
		break
	case ast.ContractDeclarationNode:
		v.scanContractDeclaration(n)
		break
	case ast.EnumDeclarationNode:
		v.scanEnumDeclaration(n)
		break
	case ast.FuncDeclarationNode:
		v.scanFuncDeclaration(n)
		break
	case ast.InterfaceDeclarationNode:
		v.scanInterfaceDeclaration(n)
		break
	case ast.ExplicitVarDeclarationNode:
		v.scanVarDeclaration(n)
		break
	case ast.EventDeclarationNode:
		v.scanEventDeclaration(n)
		break
	}
}

func (v *Validator) scanVarDeclaration(node ast.ExplicitVarDeclarationNode) {
	for _, id := range node.Identifiers {
		v.scope.scanned[id] = standards[Unknown]
	}
}

func (v *Validator) scanClassDeclaration(node ast.ClassDeclarationNode) {
	// scanning only takes in the class name
	// performs ABSOLUTELY NO CHECKING
	// checking will be done at the validation stage
	c := Class{
		Name: node.Identifier,
	}
	v.addScanned(node.Identifier, c)
}

func (v *Validator) scanEnumDeclaration(node ast.EnumDeclarationNode) {
	c := Enum{
		Name: node.Identifier,
	}
	v.addScanned(node.Identifier, c)
}

func (v *Validator) scanContractDeclaration(node ast.ContractDeclarationNode) {
	c := Contract{
		Name: node.Identifier,
	}
	v.addScanned(node.Identifier, c)
}

func (v *Validator) scanInterfaceDeclaration(node ast.InterfaceDeclarationNode) {
	c := Interface{
		Name: node.Identifier,
	}
	v.addScanned(node.Identifier, c)
}

func (v *Validator) scanFuncDeclaration(node ast.FuncDeclarationNode) {
	c := Func{
		name: node.Identifier,
	}
	v.addScanned(node.Identifier, c)
}

func (v *Validator) scanEventDeclaration(node ast.EventDeclarationNode) {
	c := Event{
		Name: node.Identifier,
	}
	v.addScanned(node.Identifier, c)
}
