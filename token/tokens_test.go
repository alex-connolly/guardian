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
	goutil.AssertNow(t, p.Name == ":", "wrong name")
}

func TestNextTokenDoubleFixed(t *testing.T) {
	b := &bytecode{bytes: []byte("+=")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "+=", "wrong name")
}

func TestNextTokenTripleFixed(t *testing.T) {
	b := &bytecode{bytes: []byte("<<=")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "<<=", "wrong name")
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
	goutil.AssertNow(t, p.Name == "in", "wrong name")
}

func TestNextTokenDistinctEnding(t *testing.T) {
	b := &bytecode{bytes: []byte("in")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "in", "wrong name")
}

func TestNextTokenInt(t *testing.T) {
	b := &bytecode{bytes: []byte("6")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "integer", "wrong name")
}

func TestNextTokenFloat(t *testing.T) {
	b := &bytecode{bytes: []byte("6.5")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "float", "wrong name")
}

func TestNextTokenString(t *testing.T) {
	b := &bytecode{bytes: []byte(`"hi"`)}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "string", "wrong name")
}

func TestNextTokenCharacter(t *testing.T) {
	b := &bytecode{bytes: []byte(`'hi'`)}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "character", "wrong name")
}
