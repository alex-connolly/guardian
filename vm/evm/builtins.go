package evm

import "github.com/end-r/vmgen"

var builtins = map[string]Builtin{
	// arithmetic
	"addmod":  simpleInstruction("ADDMOD"),
	"mulmod":  simpleInstruction("MULMOD"),
	"balance": simpleInstruction("BALANCE"),
	// transactional
	"transfer":     nil,
	"send":         send,
	"delegateCall": delegateCall,
	"call":         call,
	"callcode":     callCode,
	// error-checking
	"revert":  revert,
	"require": require,
	"assert":  assert,
	// cryptographic
	"sha3":      simpleInstruction("SHA3"),
	"keccak256": nil,
	"sha256":    nil,
	"ecrecover": nil,
	"ripemd160": nil,
	// ending
	"selfDestruct": selfDestruct,
}

type Builtin func() vmgen.Bytecode

// returns an anon func to handle simplest cases
func simpleInstruction(mnemonic string) Builtin {
	return func() (code vmgen.Bytecode) {
		code.Add(mnemonic)
		return code
	}
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
