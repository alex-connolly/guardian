package evm

import (
	"github.com/end-r/guardian/ast"

	"github.com/end-r/guardian/validator"
	"github.com/end-r/vmgen"
)

var builtins = map[string]validator.BytecodeGenerator{
	// arithmetic
	"addmod":  validator.SimpleInstruction("ADDMOD"),
	"mulmod":  validator.SimpleInstruction("MULMOD"),
	"balance": singleArgumentCall("BALANCE"),
	// transactional
	"transfer":     transfer,
	"delegateCall": delegateCall,
	"call":         call,
	//"callcode": callCode,
	// error-checking
	"revert":  validator.SimpleInstruction("REVERT"),
	"throw":   validator.SimpleInstruction("REVERT"),
	"require": require,
	"assert":  assert,
	// cryptographic
	"keccak256": validator.SimpleInstruction("SHA3"),
	"sha256":    nil,
	"ecrecover": nil,
	"ripemd160": nil,
	// ending
	"selfDestruct": singleArgumentCall("SELFDESTRUCT"),

	// message
	"calldata":  calldata,
	"gas":       validator.SimpleInstruction("GAS"),
	"sender":    validator.SimpleInstruction("CALLER"),
	"signature": signature,

	// block
	"timestamp": validator.SimpleInstruction("TIMESTAMP"),
	"number":    validator.SimpleInstruction("NUMBER"),
	"blockhash": blockhash,
	"coinbase":  validator.SimpleInstruction("COINBASE"),
	"gasLimit":  validator.SimpleInstruction("GASLIMIT"),
	// tx
	"gasPrice": validator.SimpleInstruction("GASPRICE"),
	"origin":   validator.SimpleInstruction("ORIGIN"),
}

func transfer(vm validator.VM) (code vmgen.Bytecode) {
	e := vm.(GuardianEVM)
	call := e.expression.(*ast.CallExpressionNode)
	// gas
	code.Concat(push(uintAsBytes(uint(2300))))
	// to
	code.Concat(e.traverse(call.Arguments[0]))
	// value
	code.Concat(e.traverse(call.Arguments[1]))
	// in offset
	code.Concat(push(uintAsBytes(uint(0))))
	// in size
	code.Concat(push(uintAsBytes(uint(0))))
	// out offset
	e.allocateMemory("transfer", 1)
	mem := e.lookupMemory("transfer")
	code.Concat(push(uintAsBytes(uint(mem.offset))))
	// out size
	code.Concat(push(uintAsBytes(uint(1))))
	code.Add("CALL")
	return code
}

func call(vm validator.VM) (code vmgen.Bytecode) {
	e := vm.(GuardianEVM)
	call := e.expression.(*ast.CallExpressionNode)
	// gas
	code.Concat(e.traverse(call.Arguments[1]))
	// recipient --> should be on the stack already
	code.Concat(e.traverse(call.Arguments[0]))
	// ether value
	code.Concat(e.traverse(call.Arguments[2]))
	// memory location of start of input data
	code.Add("PUSH")
	// length of input data
	code.Add("PUSH")
	// out offset
	e.allocateMemory("transfer", 1)
	mem := e.lookupMemory("transfer")
	code.Concat(push(uintAsBytes(uint(mem.offset))))
	// out size
	code.Concat(push(uintAsBytes(uint(1))))
	code.Add("CALL")
	return code
}

func calldata(vm validator.VM) (code vmgen.Bytecode) {
	code.Add("CALLDATA")
	return code
}

func singleArgumentCall(opcode string) validator.BytecodeGenerator {
	return func(vm validator.VM) vmgen.Bytecode {
		e := vm.(GuardianEVM)
		call := e.expression.(*ast.CallExpressionNode)

		code.Concat(e.traverse(call.Arguments[0]))
		code.Add(opcode)
	}
}

func blockhash(vm validator.VM) (code vmgen.Bytecode) {
	e := vm.(GuardianEVM)
	call := e.expression.(*ast.CallExpressionNode)

	code.Concat(e.traverse(call.Arguments[0]))
	code.Add("BALANCE")
	return code
}

func require(vm validator.VM) (code vmgen.Bytecode) {
	// code.Concat(pushMarker(2))

	e := vm.(GuardianEVM)
	call := e.expression.(*ast.CallExpressionNode)
	code.Concat(e.traverse(call.Arguments[0]))

	code.Add("JUMPI")
	code.Add("REVERT")
	return code
}

func assert(vm validator.VM) (code vmgen.Bytecode) {
	// TODO: invalid opcodes
	code.Concat(pushMarker(2))
	code.Add("JUMPI")
	code.Add("INVALID")
	return code
}

func delegateCall(vm validator.VM) (code vmgen.Bytecode) {
	code.Add("DELEGATECALL")
	return code
}

func callCode(vm validator.VM) (code vmgen.Bytecode) {
	code.Add("CALLCODE")
	return code
}

func signature(vm validator.VM) (code vmgen.Bytecode) {
	// get first four bytes of calldata
	return code
}

func length(vm validator.VM) (code vmgen.Bytecode) {
	// must be an array
	// array size is always at the first index
	evm := vm.(*GuardianEVM)
	if evm.inStorage {
		code.Add("SLOAD")
	} else {
		code.Add("MLOAD")
	}
	return code
}

func appendBuiltin(vm validator.VM) (code vmgen.Bytecode) {
	// must be an array
	// array size is always at the first index
	evm := vm.(*GuardianEVM)
	if evm.inStorage {
		//code.Add(push())
		code.Add("SLOAD")
		code.Concat(push(uintAsBytes(uint(1))))
		code.Add("ADD")
		code.Add("SSTORE")

	} else {
		code.Add("MLOAD")
	}
	return code
}
