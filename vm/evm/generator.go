package evm

import "github.com/end-r/vmgen"

// why are gas costs needed?
// to enable the compiler to estimate the gas usage of different execution paths
// can display to user at compile time

type EVMGenerator struct {
}

const (
	gasZero    = 0
	gasBase    = 2
	gasVeryLow = 3
	gasLow     = 5
	gasMid     = 8
	gasHigh    = 10

	gasJumpDest = 1

	gasQuickStep   = 2
	gasFastestStep = 3
	gasFastStep    = 5
	gasMidStep     = 8
	gasSlowStep    = 10
	gasExtStep     = 20

	gasContractByte = 200

	gasBlockhash = 20
)

type Instruction struct {
	opcode []byte
	cost   func(interface{}) int
}

func constantGas(gas int) func() int {
	return func(nil) int {
		return gas
	}
}

func generateDups() (im vmgen.InstructionMap) {
    number := 16
    for i := 0; i < number; i++ {
        im[fmt.Sprintf("DUP%d", i)] = vmgen.Instruction{Opcode: 0x80 + i, Cost: constantGas(gasZero)}
    }
    return im
}

func generateSwaps() (im vmgen.InstructionMap) {
    number := 16
    for i := 0; i < number; i++ {
        im[fmt.Sprintf("SWAP%d", i)] = vmgen.Instruction{Opcode: 0x90 + i, Cost: constantGas(gasZero)}
    }
    return im
}

func generatePushes() (im vmgen.InstructionMap) {
    number := 32
    for i := 0; i < number; i++ {
        im[fmt.Sprintf("SWAP%d", i)] = vmgen.Instruction{Opcode: 0x60 + i, Cost: constantGas(gasZero)}
    }
    return im
}

func generateLogs() (im vmgen.InstructionMap) {
    number := 5
    for i := 0; i < number; i++ {
        im[fmt.Sprintf("SWAP%d", i)] = vmgen.Instruction{Opcode: 0xA0 + i, Cost: constantGas(gasZero)}
    }
    return im
}



func (e EVMGenerator) Opcodes() vmgen.InstructionMap {
	m := vmgen.InstructionMap{
		"STOP":       vmgen.Instruction{Opcode: 0x00, Cost: constantGas(gasZero)},
		"ADD":        vmgen.Instruction{Opcode: 0x01, Cost: constantGas(gasVeryLow)},
		"MUL":        vmgen.Instruction{Opcode: 0x02, Cost: constantGas(gasLow)},
		"SUB":        vmgen.Instruction{Opcode: 0x03, Cost: constantGas(gasVeryLow)},
		"DIV":        vmgen.Instruction{Opcode: 0x04, Cost: constantGas(gasLow)},
		"SDIV":       vmgen.Instruction{Opcode: 0x05, Cost: constantGas(gasLow)},
		"MOD":        vmgen.Instruction{Opcode: 0x06, Cost: constantGas(gasLow)},
		"SMOD":       vmgen.Instruction{Opcode: 0x07, Cost: constantGas(gasLow)},
		"ADDMOD":     vmgen.Instruction{Opcode: 0x08, Cost: constantGas(gasMid)},
		"MULMOD":     vmgen.Instruction{Opcode: 0x09, Cost: constantGas(gasMid)},
		"EXP":        vmgen.Instruction{Opcode: 0x0A, Cost: },
		"SIGNEXTEND": vmgen.Instruction{Opcode: 0x0B, Cost: constantGas(gasLow)},

        "LT": vmgen.Instruction{Opcode: 0x10, Cost: constantGas(gasVeryLow)},
        "GT": vmgen.Instruction{Opcode: 0x11, Cost: constantGas(gasVeryLow)},
        "SLT": vmgen.Instruction{Opcode: 0x12, Cost: constantGas(gasVeryLow)},
        "SGT": vmgen.Instruction{Opcode: 0x13, Cost: constantGas(gasVeryLow)},
        "EQ": vmgen.Instruction{Opcode: 0x14, Cost: constantGas(gasVeryLow)},
        "ISZERO": vmgen.Instruction{Opcode: 0x15, Cost: constantGas(gasVeryLow)},
        "AND": vmgen.Instruction{Opcode: 0x16, Cost: constantGas(gasVeryLow)},
        "OR": vmgen.Instruction{Opcode: 0x17, Cost: constantGas(gasVeryLow)},
        "XOR": vmgen.Instruction{Opcode: 0x18, Cost: constantGas(gasVeryLow)},
        "NOT": vmgen.Instruction{Opcode: 0x19, Cost: constantGas(gasVeryLow)},
        "BYTE": vmgen.Instruction{Opcode: 0x1A, Cost: constantGas(gasVeryLow)},

        "SHA3": vmgen.Instruction{Opcode: 0x20, Cost: },

        "ADDRESS": vmgen.Instruction{Opcode: 0x30, Cost: constantGas(gasBase)},
        "BALANCE": vmgen.Instruction{Opcode: 0x31, Cost: },
        "ORIGIN": vmgen.Instruction{Opcode: 0x32, Cost: constantGas(gasBase)},
        "CALLER": vmgen.Instruction{Opcode: 0x33, Cost: constantGas(gasBase)},
        "CALLVALUE": vmgen.Instruction{Opcode: 0x34, Cost: constantGas(gasBase)},
        "CALLDATALOAD": vmgen.Instruction{Opcode: 0x35, Cost: constantGas(gasVeryLow)},
        "CALLDATASIZE": vmgen.Instruction{Opcode: 0x36, Cost: constantGas(gasBase)},
        "CALLDATACOPY": vmgen.Instruction{Opcode: 0x37, Cost: },
        "CODESIZE": vmgen.Instruction{Opcode: 0x38, Cost: },
        "CODECOPY": vmgen.Instruction{Opcode: 0x39, Cost: },
        "GASPRICE": vmgen.Instruction{Opcode: 0x3A, Cost: constantGas(gasBase)},
        "EXTCODESIZE": vmgen.Instruction{Opcode: 0x3B, Cost: constantGas(gasExtCode)},
        "EXTCODECOPY": vmgen.Instruction{Opcode: 0x3C, Cost: },

        "BLOCKHASH": vmgen.Instruction{Opcode: 0x10, Cost: },
        "COINBASE": vmgen.Instruction{Opcode: 0x10, Cost: constantGas(gasBase)},
        "TIMESTAMP": vmgen.Instruction{Opcode: 0x10, Cost: constantGas(gasBase)},
        "NUMBER": vmgen.Instruction{Opcode: 0x10, Cost: constantGas(gasBase)},
        "DIFFICULTY": vmgen.Instruction{Opcode: 0x10, Cost: constantGas(gasBase)},
        "GASLIMIT": vmgen.Instruction{Opcode: 0x10, Cost: constantGas(gasBase)},

        "POP": vmgen.Instruction{Opcode: 0x50, Cost: constantGas(gasBase)},
        "MLOAD": vmgen.Instruction{Opcode: 0x51, Cost: constantGas(gasVeryLow)},
        "MSTORE": vmgen.Instruction{Opcode: 0x52, Cost: constantGas(gasVeryLow)},
        "MSTORE8": vmgen.Instruction{Opcode: 0x53, Cost: constantGas(gasVeryLow)},
        "SLOAD": vmgen.Instruction{Opcode: 0x54, Cost: },
        "SSTORE": vmgen.Instruction{Opcode: 0x55, Cost: },
        "JUMP": vmgen.Instruction{Opcode: 0x56, Cost: constantGas(gasMid)},
        "JUMPI": vmgen.Instruction{Opcode: 0x57, Cost: constantGas(gasHigh)},
        "PC": vmgen.Instruction{Opcode: 0x58, Cost: constantGas(gasBase)},
        "MSIZE": vmgen.Instruction{Opcode: 0x59, Cost: constantGas(gasBase)},
        "GAS": vmgen.Instruction{Opcode: 0x5A, Cost: constantGas(gasBase)},
        "JUMPDEST": vmgen.Instruction{Opcode: 0x5B, Cost: },
	}

    m.AddAll(vmgen.GeneratePushes())
    m.AddAll(vmgen.GenerateDups())
    m.AddAll(vmgen.GenerateLogs())
    m.AddAll(vmgen.GenerateSwaps())

    return m
}
