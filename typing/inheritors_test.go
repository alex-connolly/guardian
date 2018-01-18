package typing

// TODO: can't use parser here
/*
func TestClassInheritsTypeValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class LightSource {}
        class Light inherits LightSource {}

        item LightSource

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	goutil.Assert(t, len(errs) == 0, errs.Format())
}

func TestClassInheritsMultipleTypesValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class LightSource {}
        class Object {}
        class Light inherits LightSource, Object {}

        item LightSource

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	goutil.Assert(t, len(errs) == 0, errs.Format())
}

func TestClassDoesNotInherit(t *testing.T) {
	scope, _ := parser.ParseString(`
        class LightSource {}
        class Light {}

        item LightSource

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	goutil.Assert(t, len(errs) == 1, errs.Format())
}

func TestClassImplementsMultipleInheritanceValid(t *testing.T) {
	scope, _ := parser.ParseString(`
		class Object {}
        class LightSource inherits Object {}
        class Light inherits LightSource {}

        item Object

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	goutil.Assert(t, len(errs) == 0, errs.Format())
}
*/
