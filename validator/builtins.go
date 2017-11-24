package validator

type Builtin interface {
	Name() string
	Children() map[string]Builtin
	Type() Type
}

// used to pass types into the s
type NumericType struct {
	size     int
	name     string
	integral bool
}

type ArrayType struct {
	// will be checked at runtime
	key    string
	length int
}

type BooleanType struct {
}
