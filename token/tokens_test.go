package token

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/util"

	"github.com/end-r/goutil"
)

type bytecode struct {
	bytes  []byte
	offset uint
}

func (b *bytecode) Offset() uint {
	return b.offset
}

func (b *bytecode) SetOffset(o uint) {
	b.offset = o
}

func (b *bytecode) Bytes() []byte {
	return b.bytes
}

func (b *bytecode) Location() util.Location {
	return util.Location{
		Offset: b.offset,
	}
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

func TestNextTokenDistinctElif(t *testing.T) {
	b := &bytecode{bytes: []byte("elif")}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "elif", fmt.Sprintf("wrong name: %s", p.Name))
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

func TestNextTokenLongerString(t *testing.T) {
	b := &bytecode{bytes: []byte(`"hello this is dog"`)}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "string", fmt.Sprintf("wrong name: %s", p.Name))
}

func TestNextTokenAssignment(t *testing.T) {
	b := &bytecode{bytes: []byte(`x = "hello this is dog"`)}
	p := NextProtoToken(b)
	p.Process(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "identifier", fmt.Sprintf("1 wrong name: %s", p.Name))
	// ignore
	p = NextProtoToken(b)
	p.Process(b)
	p = NextProtoToken(b)
	p.Process(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "=", fmt.Sprintf("2 wrong name: %s", p.Name))
	// ignore
	p = NextProtoToken(b)
	p.Process(b)
	p = NextProtoToken(b)
	p.Process(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "string", fmt.Sprintf("3 wrong name: %s", p.Name))
}

func TestNextTokenHexadecimal(t *testing.T) {
	byt := []byte(`0x00001`)
	b := &bytecode{bytes: byt}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "integer", fmt.Sprintf("wrong name: %s", p.Name))
	tok := p.Process(b)
	goutil.AssertLength(t, int(tok.End.Offset), len(byt))
}

func TestNextTokenLongHexadecimal(t *testing.T) {
	byt := []byte(`0x0000FFF00000`)
	b := &bytecode{bytes: byt}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "integer", fmt.Sprintf("wrong name: %s", p.Name))
	tok := p.Process(b)
	goutil.AssertLength(t, int(tok.End.Offset), len(byt))
}

func TestNextTokenSingleZero(t *testing.T) {
	byt := []byte(`0`)
	b := &bytecode{bytes: byt}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "integer", fmt.Sprintf("wrong name: %s", p.Name))
	tok := p.Process(b)
	goutil.AssertLength(t, int(tok.End.Offset), len(byt))
}

func TestNextTokenNegativeInt(t *testing.T) {
	byt := []byte(`-55`)
	b := &bytecode{bytes: byt}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "integer", fmt.Sprintf("wrong name: %s", p.Name))
	tok := p.Process(b)
	goutil.AssertLength(t, int(tok.End.Offset), len(byt))
}

func TestNextTokenNegativeFloat(t *testing.T) {
	byt := []byte(`-55.00`)
	b := &bytecode{bytes: byt}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "float", fmt.Sprintf("wrong name: %s", p.Name))
	tok := p.Process(b)
	goutil.AssertLength(t, int(tok.End.Offset), len(byt))
}

func TestNextTokenSingleCharacter(t *testing.T) {
	byt := []byte(`'a'`)
	b := &bytecode{bytes: byt}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "character", fmt.Sprintf("wrong name: %s", p.Name))
	tok := p.Process(b)
	goutil.AssertLength(t, int(tok.End.Offset), len(byt))
}

func TestNextTokenSingleCharacterNewLine(t *testing.T) {
	byt := []byte(`'a'
		`)
	b := &bytecode{bytes: byt}
	p := NextProtoToken(b)
	goutil.AssertNow(t, p != nil, "pt nil")
	goutil.AssertNow(t, p.Name == "character", fmt.Sprintf("wrong name: %s", p.Name))
	tok := p.Process(b)
	goutil.AssertLength(t, int(tok.End.Offset), len("'a'"))
}
