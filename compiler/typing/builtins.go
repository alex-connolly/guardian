package typing

type builtin struct {
	parameters []Type
}

var builtins = map[string]builtin{

	"append":  builtin{},
	"current": builtin{},
	"assert":  builtin{},
	"require": builtin{},
	"revert":  builtin{},
	// maths
	"addmod":    builtin{},
	"mulmod":    builtin{},
	"keccak256": builtin{},
	"sha256":    builtin{},
	"sha3":      builtin{},
	"ripemd160": builtin{},
	"ecrecover": builtin{},
}

func (c *Checker) builtin() bool {
	return false
}
