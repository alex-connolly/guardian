package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestParseClassTerminating(t *testing.T) {
	_, errs := ParseString(`class`)
	goutil.AssertNow(t, len(errs) == 3, errs.Format())
}

func TestParseClassNoIdentifier(t *testing.T) {
	_, errs := ParseString(`class {}`)
	goutil.AssertNow(t, len(errs) > 1, errs.Format())
}

func TestParseClassNoBraces(t *testing.T) {
	_, errs := ParseString(`class Dog`)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestParseClassNoOpenBrace(t *testing.T) {
	_, errs := ParseString(`class Dog }`)
	goutil.AssertNow(t, len(errs) > 1, errs.Format())
}

func TestParseClassNoCloseBrace(t *testing.T) {
	_, errs := ParseString(`class Dog {`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestParseVarNoVarStatement(t *testing.T) {
	_, errs := ParseString(`name string`)
	goutil.AssertNow(t, len(errs) > 0, errs.Format())
}
