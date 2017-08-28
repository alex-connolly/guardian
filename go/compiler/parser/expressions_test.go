package parser

import (
	"testing"
)

func TestParseReferenceSingle(t *testing.T) {
	p := createParser(`hello`)
	p.parseExpression()
}

func TestParseReferenceMultiple(t *testing.T) {
	p := createParser(`hello.aaa.bb`)
	p.parseExpression()
}

func TestParseConstant(t *testing.T) {
	p := createParser(`6`)
	p.parseExpression()
}
