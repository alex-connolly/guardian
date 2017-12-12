package evm

import (
	"encoding/binary"
	"strconv"

	"github.com/end-r/guardian/lexer"

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
	storage            map[string]storageBlock
	memoryCursor       uint
	memory             map[string]memoryBlock
	currentlyAssigning string
}

func (e *GuardianEVM) pusher(data ...byte) (code vmgen.Bytecode) {
	if len(data) > 32 {
		// problemo
	} else {
		code.Add("PUSH"+strconv.Itoa(len(data)), data)
	}
}

type storageBlock struct {
	name   string
	size   uint
	offset uint
	slot   uint
}

type memoryBlock struct {
	size   uint
	offset uint
}

const (
	wordSize = uint(256)
)

func (s storageBlock) retrive() (code vmgen.Bytecode) {
	if s.size > wordSize {
		// over at least 2 slots
		first := wordSize - s.offset
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, first))

		remaining := s.size - first
		slot := s.slot
		for remaining >= wordSize {
			// get whole slot
			slot += 1
			code.Concat(getByteSectionOfSlot(slot, 0, wordSize))
			remaining -= wordSize
		}
		if remaining > 0 {
			// get first remaining bits from next slot
			code.Concat(getByteSectionOfSlot(slot+1, 0, remaining))
		}
	} else if s.offset+s.size > wordSize {
		// over 2 slots
		// get last wordSize - s.offset bits from first
		first := wordSize - s.offset
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, first))
		// get first s.size - (wordSize - s.offset) bits from second
		code.Concat(getByteSectionOfSlot(s.slot+1, 0, s.size-first))
	} else {
		// all within 1 slot
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, s.size))
	}
	return code
}

func getByteSectionOfSlot(slot, start, size uint) (code vmgen.Bytecode) {
	code.Add("PUSH", uintAsBytes(slot)...)
	code.Add("SLOAD")
	mask := ((1 << size) - 1) << start
	code.Add("PUSH", uintAsBytes(uint(mask))...)
	code.Add("AND")
	return code
}

func uintAsBytes(a uint) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(a))
	return bs
}

func (m memoryBlock) retrieve() (code vmgen.Bytecode) {
	for size := m.size; size > wordSize; size -= wordSize {
		loc := m.offset + (m.size - size)
		code.Add("PUSH", uintAsBytes(loc)...)
		code.Add("MLOAD")
	}
	return code
}

func (s storageBlock) store() (code vmgen.Bytecode) {
	if s.name != "" {

	} else {

	}
	return code
}

func (m memoryBlock) store() (code vmgen.Bytecode) {
	free := []byte{0x40}
	code.Add("PUSH", free...)
	code.Add("PUSH", uintAsBytes(m.offset)...)
	code.Add("MSTORE")
	return code
}

func (evm *GuardianEVM) allocateMemory(name string, size uint) {
	if evm.memory == nil {
		evm.memory = make(map[string]memoryBlock)
	}
	block := memoryBlock{
		size:   size,
		offset: evm.memoryCursor,
	}
	evm.memoryCursor += size
	evm.memory[name] = block
}

func (evm *GuardianEVM) allocateStorage(name string, size uint) {
	// TODO: check whether there's a way to using some bin packing algo
	// with a modified heuristic to reduce the cost of extracting variables using bitshifts
	// maybe?
	if evm.storage == nil {
		evm.storage = make(map[string]storageBlock)
	}
	block := storageBlock{
		size:   size,
		offset: evm.lastOffset,
		slot:   evm.lastSlot,
	}
	for size > wordSize {
		size -= wordSize
		evm.lastSlot += 1
	}
	evm.lastOffset += size
	evm.storage[name] = block
}

func (evm *GuardianEVM) lookupStorage(name string) storageBlock {
	if evm.storage == nil {

	}
	return evm.storage[name]
}

func (evm GuardianEVM) Builtins() *ast.ScopeNode {
	ast, errs := parser.ParseFile(`builtins.grd`)
	if errs != nil {

	}
	return ast
}

func (evm GuardianEVM) BooleanName() string {
	return "bool"
}

func (evm GuardianEVM) Literals() validator.LiteralMap {
	return validator.LiteralMap{
		lexer.TknString:  validator.SimpleLiteral("string"),
		lexer.TknTrue:    validator.BooleanLiteral,
		lexer.TknFalse:   validator.BooleanLiteral,
		lexer.TknInteger: resolveIntegerLiteral,
		lexer.TknFloat:   resolveFloatLiteral,
	}
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
	m = validator.OperatorMap{}

	m.Add(validator.BooleanOperator, lexer.TknGeq, lexer.TknLeq,
		lexer.TknLss, lexer.TknNeq, lexer.TknEql, lexer.TknGtr)

	m.Add(operatorAdd, lexer.TknAdd)
	m.Add(validator.BooleanOperator, lexer.TknLogicalAnd, lexer.TknLogicalOr)

	// numericalOperator with floats/ints
	m.Add(validator.BinaryNumericOperator, lexer.TknSub, lexer.TknMul, lexer.TknDiv,
		lexer.TknMod)

	// integers only
	m.Add(validator.BinaryIntegerOperator, lexer.TknShl, lexer.TknShr)

	m.Add(validator.CastOperator, lexer.TknAs)

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

func (evm GuardianEVM) Traverse(node ast.Node) (vmgen.Bytecode, util.Errors) {
	// do pre-processing/hooks etc
	code := evm.traverse(node)
	// generate the bytecode
	// finalise the bytecode
	//evm.finalise()
	return code, nil
}

type hook struct {
	name     string
	position int
	bytecode []byte
}

type callable struct {
	name     string
	position int
	bytecode []byte
}

// NewGuardianEVM ...
func NewVM() GuardianEVM {
	return GuardianEVM{}
}

// A hook conditionally jumps the code to a particular point
//

func (e GuardianEVM) finalise() {
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
	for _, d := range s.Declarations.Array() {
		code.Concat(e.traverse(d.(ast.Node)))
	}
	return code
}

func (e *GuardianEVM) inStorage() bool {
	return true
}
