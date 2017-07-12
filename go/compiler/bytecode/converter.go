package bytecode

import (
	"axia/guardian/go/compiler/lexer"

	"github.com/end-r/firevm"
)

func convertOperator(t lexer.TokenType) firevm.Instruction {
	switch t {
	case lexer.TknAdd:
		return firevm.AddInstruction("ADD")
	case lexer.TknSub:
		return firevm.AddInstruction("SUB")
	case lexer.TknDiv:
		return firevm.AddInstruction("DIV")
	case lexer.TknNot:
		return firevm.AddInstruction("NOT")
	}
}
