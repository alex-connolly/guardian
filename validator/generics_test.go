package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
        class List<T> {

        }

        a = List<String>()
    `)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}
