package parser

const (
	errInvalidInterfaceProperty   = "Everything in an interface must be a func type"
	errInvalidEnumProperty        = "Everything in an enum must be an identifier"
	errMixedNamedParameters       = "Mixed named and unnamed parameters"
	errInvalidArraySize           = "Invalid array size"
	errEmptyGroup                 = "Group declaration must apply modifiers"
	errDanglingExpression         = "Dangling expression"
	errConstantWithoutValue       = "Constants must have a value"
	errUnclosedGroup              = "Unclosed group not allowed"
	errInvalidScopeDeclaration    = "Invalid declaration in scope"
	errRequiredType               = "Required one of {%s}, found %s"
	errInvalidAnnotationParameter = "Invalid annotation parameter, must be string"
	errInvalidIncDec              = "Cannot increment or decrement in this context"
)
