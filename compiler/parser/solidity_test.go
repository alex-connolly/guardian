package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

// tests conversions of the solidity examples

func TestParseVotingExample(t *testing.T) {
	p := ParseFile("tests/solc/voting.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, p.formatErrors())
}

func TestParseSimpleAuctionExample(t *testing.T) {
	p := ParseFile("tests/solc/simple_auction.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, p.formatErrors())
}

func TestParseBlindAuctionExample(t *testing.T) {
	p := ParseFile("tests/solc/blind_auction.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, p.formatErrors())
}

func TestParsePurchaseExample(t *testing.T) {
	p := ParseFile("tests/solc/purchase.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, p.formatErrors())
}
