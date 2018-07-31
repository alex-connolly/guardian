package evm

import (
	"fmt"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/parser"
	"github.com/end-r/guardian/token"
	"github.com/end-r/guardian/typing"
	"github.com/end-r/guardian/validator"
)

func (evm GuardianEVM) Builtins() *ast.ScopeNode {
	if builtinScope == nil {
		builtinScope, _ = parser.ParseFile("builtins.grd")
	}
	return builtinScope
}

func (evm GuardianEVM) BooleanName() string {
	return "bool"
}

func (evm GuardianEVM) Literals() validator.LiteralMap {
	if litMap == nil {
		litMap = validator.LiteralMap{
			token.String:  validator.SimpleLiteral("string"),
			token.True:    validator.BooleanLiteral,
			token.False:   validator.BooleanLiteral,
			token.Integer: resolveIntegerLiteral,
			token.Float:   resolveFloatLiteral,
		}
	}
	return litMap
}

func resolveIntegerLiteral(v *validator.Validator, data string) typing.Type {
	if len(data) == len("0x")+20 {
		// this might be an address
	}
	x := typing.BitsNeeded(len(data))
	return v.SmallestNumericType(x, false)
}

func resolveFloatLiteral(v *validator.Validator, data string) typing.Type {
	// convert to float
	return typing.Unknown()
}

func (evm *GuardianEVM) Primitives() map[string]typing.Type {

	const maxSize = 256
	m := map[string]typing.Type{}

	const increment = 8
	for i := increment; i <= maxSize; i += increment {
		ints := fmt.Sprintf("int%d", i)
		uints := "u" + ints
		m[uints] = &typing.NumericType{Name: uints, BitSize: i, Signed: false, Integer: true}
		m[ints] = &typing.NumericType{Name: ints, BitSize: i, Signed: true, Integer: true}
	}

	return m
}

func (evm GuardianEVM) ValidDeclarations() []ast.NodeType {
	return ast.AllDeclarations
}

func (evm *uardianEVM) ValidExpressions() []ast.NodeType {
	return ast.AllExpressions
}

func (evm GuardianEVM) ValidStatements() []ast.NodeType {
	return ast.AllStatements
}

// TODO: context stacks
// want to validate that a modifier can only be applied to a function in a contract
var mods = []*validator.ModifierGroup{
	&validator.ModifierGroup{
		Name:       "Visibility",
		Modifiers:  []string{"external", "internal", "global"},
		RequiredOn: []ast.NodeType{ast.FuncDeclaration},
		AllowedOn:  []ast.NodeType{ast.FuncDeclaration},
		Maximum:    1,
	},
	&validator.ModifierGroup{
		Name:       "Payable",
		Modifiers:  []string{"payable", "nonpayable"},
		RequiredOn: []ast.NodeType{},
		AllowedOn:  []ast.NodeType{ast.FuncDeclaration},
		Maximum:    1,
	},
	&validator.ModifierGroup{
		Name:       "Payable",
		Modifiers:  []string{"payable", "nonpayable"},
		RequiredOn: []ast.NodeType{},
		AllowedOn:  []ast.NodeType{ast.FuncDeclaration},
		Maximum:    1,
	},
}

func (evm GuardianEVM) Modifiers() []*validator.ModifierGroup {
	return mods
}

func (evm GuardianEVM) Annotations() []*typing.Annotation {
	return nil
}

func (evm *GuardianEVM) Assignable(val *validator.Validator, left, right typing.Type, fromExpression ast.ExpressionNode) bool {
	t, _ := val.IsTypeVisible("address")
	if t.Compare(right) {
		switch left.(type) {
		case *typing.Contract:
			return true
		case *typing.Interface:
			return true
		}
	}
	if !typing.AssignableTo(left, right, true) {
		if validator.LiteralAssignable(left, right, fromExpression) {
			return true
		}
		return false
	}
	return true
}

func (evm GuardianEVM) Castable(val *validator.Validator, to, from typing.Type, fromExpression ast.ExpressionNode) bool {
	// can cast all addresses to all contracts
	t, _ := val.IsTypeVisible("address")
	if t.Compare(from) {
		switch to.(type) {
		case *typing.Contract:
			return true
		case *typing.Interface:
			return true
		}
	}
	// can cast all uints to addresses
	if t.Compare(to) {
		switch a := typing.ResolveUnderlying(from).(type) {
		case *typing.NumericType:
			if !a.Signed {
				// check size
				return true
			}
			if l, ok := fromExpression.(*ast.LiteralNode); ok {
				hasSign := (l.Data[0] == '-')
				if !hasSign {
					return true
				}
			}
		}
	}
	if typing.AssignableTo(to, from, false) {
		return true
	}
	if validator.LiteralAssignable(to, from, fromExpression) {
		return true
	}
	return false
}

func (evm GuardianEVM) BaseContract() (*ast.ContractDeclarationNode, util.Errors) {
	s, errs := parser.ParseString(`
		contract Base {

		    public var balance uint

		}
	`)
	c := s.Declarations.Next().(*ast.ContractDeclarationNode)
	return c, errs
}

func (evm GuardianEVM) BaseClass() (*ast.ContractDeclarationNode, util.Errors) {
	s, errs := parser.ParseString(`
		class Object {

		}
	`)
	c := s.Declarations.Next().(*ast.ContractDeclarationNode)
	return c, errs
}

func (evm GuardianEVM) BytecodeGenerators() map[string]validator.BytecodeGenerator {
	return builtins
}
