package evm

import (
	"github.com/end-r/guardian/validator"
	"github.com/end-r/vmgen"
)

var builtins = map[string]validator.BytecodeGenerator{
	// arithmetic
	"addmod":  validator.SimpleInstruction("ADDMOD"),
	"mulmod":  validator.SimpleInstruction("MULMOD"),
	"balance": validator.SimpleInstruction("BALANCE"),
	// transactional
	"transfer":     nil,
	"send":         send,
	"delegateCall": delegateCall,
	"call":         call,
	"callcode":     callCode,
	// error-checking
	"revert":  validator.SimpleInstruction("REVERT"),
	"require": require,
	"assert":  assert,
	// cryptographic
	"sha3":      validator.SimpleInstruction("SHA3"),
	"keccak256": nil,
	"sha256":    nil,
	"ecrecover": nil,
	"ripemd160": nil,
	// ending
	"selfDestruct": validator.SimpleInstruction("SELFDESTRUCT"),

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

func send(vm validator.VM) (code vmgen.Bytecode) {
	// Thus the operand order is: gas, to, value, in offset, in size, out offset, out size
	code.Add("CALL")
	return code
}

func call(vm validator.VM) (code vmgen.Bytecode) {
	// Thus the operand order is: gas, to, value, in offset, in size, out offset, out size
	code.Add("CALL")
	return code
}

func calldata(vm validator.VM) (code vmgen.Bytecode) {
	code.Add("CALLDATA")
	return code
}

func blockhash(vm validator.VM) (code vmgen.Bytecode) {
	code.Add("BLOCKHASH")
	return code
}

func require(vm validator.VM) (code vmgen.Bytecode) {
	code.Concat(pushMarker(2))
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
