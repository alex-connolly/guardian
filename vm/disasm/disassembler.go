package disasm

import "fmt"

func Disasm(data []byte) []vm.Operation {
	count := 1
	index := 0
	for index < len(data) {
		// first byte is always an opcode
		op := vm[data[index]]
		// currently, the only 2 argument opcode is PUSH
		switch op.code {
		case "PUSH":

			break
		}
		// depending on the opcode, process the next few bytes
		// 1 | PUSH (1, 10)
		// represent parameters as tuples
		// 1 | JUMPI (10)
		fmt.Printf("%d | %s (%d, %d)")
	}
}

func DisasmString(data string) []vm.Operation {
	Disasm([]byte(data))
}
