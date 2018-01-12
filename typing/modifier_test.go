package typing

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestAddModifier(t *testing.T) {
	b := Boolean()
	AddModifier(b, "static")
	goutil.AssertNow(t, HasModifier(b, "static"), "doesn't have static")
}
