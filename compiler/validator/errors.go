package validator

const (
	errInvalidBinaryOpTypes = "Binary operator %s is not defined for operands %s and %s"
	errInvalidCall          = "Cannot use %s as arguments to function of type %s"
	errInvalidSubscriptable = "Type %s is not subscriptable"
	errPropertyNotFound     = "Type %s does not have property %s"
	errUnnamedReference     = "Unnamed reference %s"
)
