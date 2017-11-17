package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/goutil"
)

func TestClassImplementsTypeValid(t *testing.T) {
	p := parser.ParseString(`
        interface Switchable{}
        class Light is Switchable {}

        item Switchable

        constructor(){
            item = Light{}
        }
    `)
	v := ValidateScope(p.Scope)
	le := p.Scope.Declarations.Length()
	goutil.AssertNow(t, le == 4, fmt.Sprintf("wrong decl length: %d", le))
	goutil.Assert(t, len(v.errors) == 0, v.formatErrors())
}

func TestClassImplementsMultipleTypesValid(t *testing.T) {
	p := parser.ParseString(`
        interface Switchable{}
        interface Adjustable{}
        class Light is Switchable, Adjustable {}

        item Switchable

        constructor(){
            item = Light{}
        }
    `)
	v := ValidateScope(p.Scope)
	goutil.Assert(t, len(v.errors) == 0, v.formatErrors())
}

func TestClassImplementsInvalid(t *testing.T) {
	p := parser.ParseString(`
        interface Switchable{}
        class Light {}

        item Switchable

        constructor(){
            item = Light{}
        }
    `)
	v := ValidateScope(p.Scope)
	goutil.Assert(t, len(v.errors) == 1, v.formatErrors())
}

func TestClassImplementsTypeValidInterfaceInheritance(t *testing.T) {
	p := parser.ParseString(`
		interface Adjustable{}
        interface Switchable inherits Adjustable {}
        class Light is Switchable {}

        item Adjustable

        constructor(){
            item = Light{}
        }
    `)
	v := ValidateScope(p.Scope)
	goutil.Assert(t, len(v.errors) == 0, v.formatErrors())
}

func TestClassImplementsTypeValidClassAndInterfaceInheritance(t *testing.T) {
	p := parser.ParseString(`

		interface Adjustable{}
        interface Switchable inherits Adjustable {}
		class Object is Switchable{}
        class Light inherits Object {}

        item Adjustable

        constructor(){
            item = Light{}
        }
    `)
	v := ValidateScope(p.Scope)
	goutil.Assert(t, len(v.errors) == 0, v.formatErrors())
}
