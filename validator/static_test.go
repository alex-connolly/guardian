package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestStaticVariableAccess(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
        class Dog {
            static var name string
        }
        x = Dog.name
    `)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestNotStaticVariableAccess(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
        class Dog {
            var name string
        }
        x = Dog.name
    `)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestEnumVariableAccess(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
        enum Days { Mon }
        x = Days.Mon
    `)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}
