package evm

import (
	"fmt"

	"github.com/end-r/vmgen"
)

func invalidBytecodeMessage(actual, expected vmgen.Bytecode) string {
	return fmt.Sprintf("Expected: %s\nActual: %s", expected.Format(), actual.Format())
}
