package parser

import (
	"axia/guardian/go/util"
	"testing"
)

func TestIsMapLiteral(t *testing.T) {
	p := createParser("map[key]value{}")
	util.Assert(t, p.isMapLiteral(), "map literal not found")
}

func TestIsArrayLiteral(t *testing.T) {

}

func TestIsLiteral(t *testing.T) {

}

func TestIsCompositeLiteral(t *testing.T) {

}
