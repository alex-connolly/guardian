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
	"selfDestruct": selfDestruct,
}

func send() (code vmgen.Bytecode) {
	// Thus the operand order is: gas, to, value, in offset, in size, out offset, out size
	code.Add("CALL")
	return code
}

func call() (code vmgen.Bytecode) {
	// Thus the operand order is: gas, to, value, in offset, in size, out offset, out size
	code.Add("CALL")
	return code
}

func revert() (code vmgen.Bytecode) {
	code.Add("REVERT")
	return code
}

func require() (code vmgen.Bytecode) {
	code.Concat(pushMarker(2))
	code.Add("JUMPI")
	code.Add("REVERT")
	return code
}

func assert() (code vmgen.Bytecode) {
	// TODO: invalid opcodes
	code.Concat(pushMarker(2))
	code.Add("JUMPI")
	code.Add("INVALID")
	return code
}

func selfDestruct() (code vmgen.Bytecode) {
	code.Add("SELFDESTRUCT")
	return code
}

func delegateCall() (code vmgen.Bytecode) {
	code.Add("DELEGATECALL")
	return code
}

func callCode() (code vmgen.Bytecode) {
	code.Add("CALLCODE")
	return code
}
