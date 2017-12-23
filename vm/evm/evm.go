package evm

import (
	"strconv"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/parser"
	"github.com/end-r/guardian/typing"
	"github.com/end-r/guardian/util"
	"github.com/end-r/guardian/validator"
	"github.com/end-r/vmgen"
)

// GuardianEVM
type GuardianEVM struct {
	VM                 *vmgen.VM
	hooks              []hook
	callables          []callable
	lastSlot           uint
	lastOffset         uint
	storage            map[string]*storageBlock
	freedMemory        []*memoryBlock
	memoryCursor       uint
	memory             map[string]*memoryBlock
	currentlyAssigning string
}

func push(data []byte) (code vmgen.Bytecode) {
	if len(data) > 32 {
		// TODO: error
	}
	m := "PUSH" + strconv.Itoa(len(data))
	code.Add(m, data...)
	return code
}

// support all offsets which can be stored in a 64 bit integer
func pushMarker(offset int) (code vmgen.Bytecode) {
	//TODO: fix
	code.AddMarker("PUSH"+strconv.Itoa(8), offset)
	return code
}

var (
	builtinScope *ast.ScopeNode
	litMap       validator.LiteralMap
	opMap        validator.OperatorMap
)

func (evm GuardianEVM) Builtins() *ast.ScopeNode {
	if builtinScope == nil {
		builtinScope, _ = parser.ParseString(`

			type string []byte
			type address [20]byte

			wei = 1
			kwei = 1000 * wei
			babbage = kwei
			mwei = 1000 * kwei
			lovelace = mwei
			gwei = 1000 * mwei
			shannon = gwei
			microether = 1000 * gwei
			szabo = microether
			milliether = 1000 * microether
			finney = milliether
			ether = 1000 * milliether

			// account functions
			balance func(a address) uint256
			transfer func(a address, amount uint256) uint
			send func(a address, amount uint256) bool
			call func(a address) bool
			delegateCall func(a address)

			// cryptographic functions
			addmod func(x, y, k uint) uint
			mulmod func(x, y, k uint) uint
			keccak256 func()
			sha256 func()
			sha3 func()
			ripemd160 func()
			ecrecover func (v uint8, h, r, s bytes32) address

			// contract functions
			// NO THIS KEYWORD: confusing for most programmers, unintentional bugs etc

			selfDestruct func(recipient address) uint256


			class BuiltinMessage {
				data []byte
				gas uint
				sender address
				sig [4]byte
			}

			msg BuiltinMessage

			class BuiltinBlock {
				timestamp uint
				number uint
				coinbase address
				gaslimit uint
				blockhash func(blockNumber uint) [32]byte
			}

			block BuiltinBlock

			class BuiltinTransaction {
				gasprice uint
				origin address
			}

			tx BuiltinTransaction

		`)
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
	x := typing.BitsNeeded(len(data))
	return v.SmallestNumericType(x, false)
}

func resolveFloatLiteral(v *validator.Validator, data string) typing.Type {
	// convert to float
	return typing.Unknown()
}

func (evm GuardianEVM) Operators() (m validator.OperatorMap) {

	if opMap != nil {
		return opMap
	}
	m = validator.OperatorMap{}

	m.Add(validator.BooleanOperator, token.Geq, token.Leq,
		token.Lss, token.Neq, token.Eql, token.Gtr, token.Or, token.And)

	m.Add(operatorAdd, token.Add)
	m.Add(validator.BooleanOperator, token.LogicalAnd, token.LogicalOr)

	// numericalOperator with floats/ints
	m.Add(validator.BinaryNumericOperator, token.Sub, token.Mul, token.Div,
		token.Mod)

	// integers only
	m.Add(validator.BinaryIntegerOperator, token.Shl, token.Shr)

	m.Add(validator.CastOperator, token.As)

	opMap = m

	return m
}

func operatorAdd(v *validator.Validator, ts ...typing.Type) typing.Type {
	switch ts[0].(type) {
	case typing.NumericType:
		return validator.BinaryNumericOperator(v, ts...)
	}
	return typing.Invalid()
}

func (evm GuardianEVM) Primitives() map[string]typing.Type {

	const maxSize = 256
	m := map[string]typing.Type{
		"int":  typing.NumericType{Name: "int", BitSize: maxSize, Signed: false, Integer: true},
		"uint": typing.NumericType{Name: "uint", BitSize: maxSize, Signed: true, Integer: true},
		"byte": typing.NumericType{Name: "byte", BitSize: 8, Signed: true, Integer: true},
	}

	const increment = 8
	for i := increment; i <= maxSize; i += increment {
		ints := "int" + string(i)
		uints := "u" + ints
		m[uints] = typing.NumericType{Name: uints, BitSize: i, Signed: false, Integer: true}
		m[ints] = typing.NumericType{Name: ints, BitSize: i, Signed: true, Integer: true}
	}

	return m
}

func (evm GuardianEVM) ValidDeclarations() []ast.NodeType {
	return ast.AllDeclarations
}

func (evm GuardianEVM) ValidExpressions() []ast.NodeType {
	return ast.AllExpressions
}

func (evm GuardianEVM) ValidStatements() []ast.NodeType {
	return ast.AllStatements
}

func (evm GuardianEVM) Traverse(node ast.Node) (vmgen.Bytecode, util.Errors) {
	// do pre-processing/hooks etc
	code := evm.traverse(node)
	// generate the bytecode
	// finalise the bytecode
	//evm.finalise()
	return code, nil
}

// NewGuardianEVM ...
func NewVM() GuardianEVM {
	return GuardianEVM{}
}

// A hook conditionally jumps the code to a particular point
//

func (e *GuardianEVM) finalise() {

	// add public functions
	for _, hook := range e.publicFuncs {
		e.hookPublicFunc()
	}
	// add private functions
	// add external functions
	// add internal functions
	// add events

	for _, hook := range e.hooks {

	}
	// number of instructions =
	/*
		for _, hook := range e.hooks {
			e.VM.AddBytecode("POP")

			e.VM.AddBytecode("EQL")
			e.VM.AddBytecode("JMPI")
		}
		// if the data matches none of the function hooks
		e.VM.AddBytecode("STOP")
		for _, callable := range e.callables {
			// add function bytecode
		}*/
}

//

// can be called from outside or inside the contract
func (e *GuardianEVM) hookPublicFunc(h *hook) {

}

// can be
func (e *GuardianEVM) hookPrivateFunc(h *hook) {

}

func (e GuardianEVM) traverse(n ast.Node) (code vmgen.Bytecode) {
	/* initialise the vm
	if e.VM == nil {
		e.VM = firevm.NewVM()
	}*/
	switch node := n.(type) {
	case *ast.ScopeNode:
		return e.traverseScope(node)
	case *ast.ClassDeclarationNode:
		return e.traverseClass(node)
	case *ast.InterfaceDeclarationNode:
		return e.traverseInterface(node)
	case *ast.EnumDeclarationNode:
		return e.traverseEnum(node)
	case *ast.EventDeclarationNode:
		return e.traverseEvent(node)
	case *ast.ExplicitVarDeclarationNode:
		return e.traverseExplicitVarDecl(node)
	case *ast.TypeDeclarationNode:
		return e.traverseType(node)
	case *ast.ContractDeclarationNode:
		return e.traverseContract(node)
	case *ast.FuncDeclarationNode:
		return e.traverseFunc(node)
	case *ast.ForStatementNode:
		return e.traverseForStatement(node)
	case *ast.AssignmentStatementNode:
		return e.traverseAssignmentStatement(node)
	case *ast.CaseStatementNode:
		return e.traverseCaseStatement(node)
	case *ast.ReturnStatementNode:
		return e.traverseReturnStatement(node)
	case *ast.IfStatementNode:
		return e.traverseIfStatement(node)
	case *ast.SwitchStatementNode:
		return e.traverseSwitchStatement(node)
	}
	return code
}

func (e *GuardianEVM) traverseScope(s *ast.ScopeNode) (code vmgen.Bytecode) {
	if s == nil {
		return code
	}
	if s.Declarations != nil {
		for _, d := range s.Declarations.Array() {
			code.Concat(e.traverse(d.(ast.Node)))
		}
	}

	return code
}

func (e *GuardianEVM) inStorage() bool {
	return false
}
