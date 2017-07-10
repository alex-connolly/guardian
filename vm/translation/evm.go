package main

import "github.com/end-r/firevm"

// EVMInstruction ...
type EVMInstruction struct {
	opcode string
	params []byte
}

// Translate ...
func Translate(instructions []firevm.InstructionBlock) []EVMInstruction {
	var evms []EVMInstruction
	for i := 0; i < len(instructions); i++ {
		switch instructions[i].opcode {
		case "ADD", "MUL":
			if i < len(instructions)-1 && instructions[i].opcode == "MOD" {
				// merge consecutive add, mod instructions
				evms = append(evms, evm(instructions[i].opcode+"MOD"))
				i++
			} else {
				evms = append(evms, evm(instructions[i].opcode))
			}
			break
		case "SUB", "DIV", "MOD", "COINBASE", "TIMESTAMP", "NUMBER", "DIFFICULTY",
				"POP":
			// operations with the same EVM opcode
			evms = append(evms, evm(instructions[i].opcode))
		case "PUSH":
			// can distribute over multiple ops if needed
			for c := instructions[i].params[0]; c > 0; c -= 16 {
				if c > 16 {
					evms = append(evms, evm("PUSH16", 0))
				} else {
					evms = append(evms, evm("PUSH"+c, 0))
				}
			}
			break
		case "DUP":
			break
		case "SWAP":
			// Axia: it is possible to swap at an arbitrary depth
			// EVM: only possible to swap 16 deep
			for c := instructions[i].params[0]; c > 0; c -= 16 {

			}
			break
		case "MSTORE":
			break
		case "SLOAD":
			// will only load 32 bits of data at a time: Guardian can use 8
			for c := instructions[i].params[0]; c > 0; c -= 32 {
				evms = append(evms, evm("SLOAD"))
				if c < 32 {
					if c > 16 {
						evms = append(evms, evm("PUSH16", 0))
						evms = append(evms, evm("PUSH"+c-16))
					} else {
						evms = append(evms, evm("PUSH"+c, 0))
					}
				}
			}
			break
		case "SELFDESTRUCT":
			break
		case "PARAMS":
			evms = append(evms, evm("CALLDATA"))
			break
		case "FUELLIMIT":
			evms = append(evms, evm("GASLIMIT"))
			break
		case
		}
	}
	return evms
}

func evm(code string, params ...byte) EVMInstruction {
	return EVMInstruction{code, params}
}
