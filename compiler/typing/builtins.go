package typing

type builtin struct {
	parameters []Type
}

var builtins = map[string]builtin{
	"append": builtin{[]Type{*Array, Type}},
}

func (c *Checker) builtin() bool {
	return false
}
