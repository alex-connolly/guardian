package validator

import (
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/compiler/parser"
)

func TestClassImplementsTypeValid(t *testing.T) {
	p := parser.ParseString(`
        class LightSource {}
        class Light inherits LightSource {}

        item LightSource

        constructor(){
            item = Light{}
        }
    `)
	v := ValidateScope(p.Scope)
	goutil.Assert(t, len(v.errors) == 0, v.formatErrors())
}

func TestClassImplementsMultipleTypesValid(t *testing.T) {
	p := parser.ParseString(`
        class LightSource {}
        class Object {}
        class Light inherits LightSource, Object {}

        item LightSource

        constructor(){
            item = Light{}
        }
    `)
	goutil.Assert(t, len(v.errors) == 0, v.formatErrors())
}

func TestClassDoesNotInherit(t *testing.T) {
	p := parser.ParseString(`
        class LightSource {}
        class Light {}

        item LightSource

        constructor(){
            item = Light{}
        }
    `)
	goutil.Assert(t, len(v.errors) == 1, v.formatErrors())
}

func TestClassImplementsMultipleInheritanceValid(t *testing.T) {
	p := parser.ParseString(`
		class Object {}
        class LightSource inherits Object {}
        class Light inherits LightSource {}

        item Object

        constructor(){
            item = Light{}
        }
    `)
	goutil.Assert(t, len(v.errors) == 0, v.formatErrors())
}
