package validator

const (
	errInvalidBinaryOpTypes     = "Binary operator %s is not defined for operands %s and %s"
	errInvalidFuncCall          = "Cannot use %s as arguments to function of type %s"
	errInvalidCall              = "Type %s cannot be called"
	errInvalidConstructorCall   = "No constructor signature matches call %s"
	errInvalidSubscriptable     = "Type %s is not subscriptable"
	errPropertyNotFound         = "Type %s does not have property %s"
	errUnnamedReference         = "Unnamed reference %s"
	errTypeRequired             = "%s is not a %s type"
	errCallExpressionNoFunc     = "Cannot call non-function type %s"
	errTypeNotVisible           = "Type %s is not visible"
	errInvalidAssignment        = "Cannot assign %s = %s"
	errTypecheckingLoop         = "Typechecking loop"
	errInvalidExpressionLeft    = "Cannot assign to expression"
	errStringLiteralUnsupported = "The current VM does not support string literals"
)
