package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
        class List<T> {

        }

        a = new List<string>()
    `)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestInheritanceConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class Dog {}
		class BorderCollie inherits Dog {}
        class ConstrainedList<T inherits Dog> {

        }

        a = new ConstrainedList<string>()
		b = new ConstrainedList<BorderCollie>()
    `)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestImplementationConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {}
		class BorderCollie is Dog {}
        class ConstrainedList<T is Dog> {

        }

        a = new ConstrainedList<string>()
		b = new ConstrainedList<BorderCollie>()
    `)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestMultipleInheritanceConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class Dog {}
		class BorderCollie inherits Dog {}
        class ConstrainedList<T inherits Dog|S inherits Dog> {

        }

        a = new ConstrainedList<string|int>()
		b = new ConstrainedList<BorderCollie|BorderCollie>()
    `)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())
	// one for each wrong type
}

func TestMultipleImplementationConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {}
		class BorderCollie is Dog {}
        class ConstrainedList<T is Dog|S is Dog> {

        }

        a = new ConstrainedList<string|int>()
		b = new ConstrainedList<BorderCollie|BorderCollie>()
    `)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())
	// one for each wrong type
}

func TestMultipleMixedConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {}
		class Animal {}
		class BorderCollie inherits Animal is Dog {}
		class FakeDog is Dog {}
		class Cat inherits Animal {}
        class ConstrainedList<T is Dog inherits Animal|S inherits Animal is Dog> {

        }

        a = new ConstrainedList<string|int>()
		b = new ConstrainedList<BorderCollie|BorderCollie>()
		c = new ConstrainedList<FakeDog|Cat>()
    `)
	goutil.AssertNow(t, len(errs) == 4, errs.Format())
	// one for each wrong type
}

func TestSingleGenericInSignature(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class List<T> {

		}
		class Dog<K> inherits List<K> {

		}
    `)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	// one for each wrong type
}

func TestMultipleGenericInSignature(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class List<T> {

		}
		class Dog<K|V> inherits List<K> {

		}
    `)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	// one for each wrong type
}

func TestTypesWithoutParameters(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		enum Dog {}
		x = Dog<string>
    `)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
	// one for each wrong type
}

func TestFunctionGenericParameter(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func main(){
			var a int
			var b []int
			a = len(b)
			var c []string
			a = len(c)
		}

		func <T> lengthOf(items []T) int {
			return 0
		}
    `)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFunctionTwoGenericParameters(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func main(){
			var a []int
			var b []int
			merge(a, b)
			var c []string
			// should error
			merge(a, c)
		}

		func <T> merge(items []T, others []T) {

		}
    `)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestFunctionTwoGenericParametersReturnType(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func main(){
			var a []int
			var b []int
			a = thing(a, b)
			var c []string
			// should error
			c = thing(a, b)
		}

		func <T> thing(items []T) []T {
			return items
		}
    `)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestGenericFunctionLiteralAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func main(){
			var y func(int)int
			x = func <T> (a T) {

			}
			x = y
		}

		func <T> thing(items []T) []T {
			return items
		}
    `)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}
