package typing

type operand struct {
	typ Type
}

func (o *operand) assignableTo(t Type) {

	x := o.typ

	if identical(x, t) {
		return true
	}

	xu := x.Underlying()
	tu := t.Underlying()

	//

	// T is an interface implemented by V

}
