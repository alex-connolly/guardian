package typing

/*
func TestClassImplementsTypeValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        interface Switchable{}
        class Light is Switchable {}

        item Switchable

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	le := scope.Declarations.Length()
	goutil.AssertNow(t, le == 4, fmt.Sprintf("wrong decl length: %d", le))
	goutil.Assert(t, len(errs) == 0, errs.Format())
}

func TestClassImplementsMultipleTypesValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        interface Switchable{}
        interface Adjustable{}
        class Light is Switchable, Adjustable {}

        item Switchable

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	goutil.Assert(t, len(errs) == 0, errs.Format())
}

func TestClassImplementsInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
        interface Switchable{}
        class Light {}

        item Switchable

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	goutil.Assert(t, len(errs) == 1, errs.Format())
}

func TestClassImplementsTypeValidInterfaceInheritance(t *testing.T) {
	scope, _ := parser.ParseString(`
		interface Adjustable{}
        interface Switchable inherits Adjustable {}
        class Light is Switchable {}

        item Adjustable

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	goutil.Assert(t, len(errs) == 0, errs.Format())
}

func TestClassImplementsTypeValidClassAndInterfaceInheritance(t *testing.T) {
	scope, _ := parser.ParseString(`

		interface Adjustable{}
        interface Switchable inherits Adjustable {}
		class Object is Switchable{}
        class Light inherits Object {}

        item Adjustable

        constructor(){
            item = Light{}
        }
    `)
	errs := Validate(scope, nil, NewTestVM())
	goutil.Assert(t, len(errs) == 0, errs.Format())
}
*/
