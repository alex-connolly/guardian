package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

// tests conversions of the solidity examples

func TestParseVotingExample(t *testing.T) {
	p, errs := ParseFile("tests/solc/voting.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseSimpleAuctionExample(t *testing.T) {
	p, errs := ParseFile("tests/solc/simple_auction.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseBlindAuctionExample(t *testing.T) {
	p, errs := ParseFile("tests/solc/blind_auction.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParsePurchaseExample(t *testing.T) {
	p, errs := ParseFile("tests/solc/purchase.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorBalanceChecker(t *testing.T) {
	p, errs := ParseFile("tests/solc/examples/creator_balance_checker.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorBasicIterator(t *testing.T) {
	p, errs := ParseFile("tests/solc/examples/basic_iterator.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorGreeter(t *testing.T) {
	p, errs := ParseFile("tests/solc/examples/greeter.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}
