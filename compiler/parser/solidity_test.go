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

func TestParseBasicIterator(t *testing.T) {
	p := ParseFile("tests/solc/examples/basic_iterator.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, p.formatErrors())
}

func TestParseGreeter(t *testing.T) {
	p := ParseFile("tests/solc/examples/greeter.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, p.formatErrors())
}

func TestParseCreatorBalanceChecker(t *testing.T) {
	p := ParseFile("tests/solc/examples/creator_balance_checker.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, p.formatErrors())
}
