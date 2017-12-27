package token

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

type bytecode struct {
	bytes  []byte
	offset int
}

func (b *bytecode) Offset() int {
	return b.offset
}

func (b *bytecode) SetOffset(o int) {
	b.offset = o
}

func (b *bytecode) Bytes() []byte {
	return b.bytes
}

func TestNextTokenSingleFixed(t *testing.T) {
	b := &bytecode{bytes: []byte(":")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == ":", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenDoubleFixed(t *testing.T) {
	b := &bytecode{bytes: []byte("+=")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "+=", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenTripleFixed(t *testing.T) {
	b := &bytecode{bytes: []byte("<<=")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "<<=", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenDistinctNewLine(t *testing.T) {
	b := &bytecode{bytes: []byte(`in
        `)}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "in", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenDistinctWhitespace(t *testing.T) {
	b := &bytecode{bytes: []byte("in ")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "in", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenDistinctEnding(t *testing.T) {
	b := &bytecode{bytes: []byte("in")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "in", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenDistinctFixed(t *testing.T) {
	b := &bytecode{bytes: []byte("in(")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "in", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenInt(t *testing.T) {
	b := &bytecode{bytes: []byte("6")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "integer", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenFloat(t *testing.T) {
	b := &bytecode{bytes: []byte("6.5")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "float", fmt.Sprintf("wrong name: %s", p.Name))
	b = &bytecode{bytes: []byte(".5")}
	p = NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "float", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenString(t *testing.T) {
	b := &bytecode{bytes: []byte(`"hi"`)}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "string", fmt.Sprintf("wrong name: %s", p.Name))
	b = &bytecode{bytes: []byte("`hi`")}
	p = NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "string", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenCharacter(t *testing.T) {
	b := &bytecode{bytes: []byte(`'hi'`)}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "character", fmt.Sprintf("wrong name: %s", p.Name))
}
