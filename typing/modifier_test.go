package typing

import "testing"

func TestAddModifier(t *testing.T) {
	b := Boolean()
	AddModifier(b, "static")
}
