package parser

const (
	errInvalidInterfaceProperty = "Everything in an interface must be a func type"
	errInvalidEnumProperty      = "Everything in an enum must be an identifier"
	errMixedNamedParameters     = "Mixed named and unnamed parameters"
	errInvalidArraySize         = "Invalid array size"
	errEmptyGroup               = "Group declaration must apply modifiers"
	errDanglingExpression       = "Dangling expression"
)
