package typing

import "axia/guardian/compiler/lexer"

type OperatorPredicates map[lexer.TokenType]Predicate

var unaryPredicates = OperatorPredicates{
	lexer.TknNot: isBoolean,
	lexer.TknXor: isInteger,
}

var binaryPredicates = OperatorPredicates{
	lexer.TknAdd: func(typ Type) bool { return isNumeric(typ) || isString(typ) },
	lexer.TknSub: isNumeric,
	lexer.TknMul: isNumeric,
	lexer.TknDiv: isNumeric,
	lexer.TknMod: isInteger,

	lexer.TknAnd: isInteger,
	lexer.TknOr:  isInteger,
	lexer.TknXor: isInteger,

	lexer.TknLogicalOr:  isBoolean,
	lexer.TknLogicalAnd: isBoolean,
}

func (c *Checker) checkOperation(op lexer.TokenType) bool {
	if pred := unaryPredicates[op]; pred != nil {
		return false
	} else {
		// TODO: record an error
		return false
	}
}

func (c *Checker) binary() bool {
	return false
}

func (c *Checker) index() bool {
	return false
}
