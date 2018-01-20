package parser

const (
	errInvalidInterfaceProperty   = "Everything in an interface must be a func type"
	errInvalidEnumProperty        = "Everything in an enum must be an identifier"
	errMixedNamedParameters       = "Mixed named and unnamed parameters"
	errInvalidArraySize           = "Invalid array size"
	errEmptyGroup                 = "Group declaration must apply modifiers"
	errDanglingExpression         = "Expression evaluated but not used"
	errConstantWithoutValue       = "Constants must have a value"
	errUnclosedGroup              = "Unclosed group not allowed"
	errInvalidScopeDeclaration    = "Invalid declaration in scope"
	errRequiredType               = "Required %s, found %s"
	errInvalidAnnotationParameter = "Invalid annotation parameter, must be string"
	errInvalidIncDec              = "Cannot increment or decrement in this context"
	errInvalidTypeAfterCast       = "Expected type after cast operator"
	errIncompleteExpression       = "Incomplete expression"
	errInvalidImportPath          = "Invalid import path: %s"
	errConsecutiveExpression      = "No terminator or operator after expression: found %s"
)
