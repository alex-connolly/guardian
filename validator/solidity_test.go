package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

// tests conversions of the solidity examples

func TestParseVotingExample(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/voting.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseSimpleAuctionExample(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/simple_auction.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseBlindAuctionExample(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/blind_auction.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParsePurchaseExample(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/purchase.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorBalanceChecker(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/examples/creator_balance_checker.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorBasicIterator(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/examples/basic_iterator.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorGreeter(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/examples/greeter.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseCrowdFunder(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/crowd_funder.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

/*
func TestParseStrings(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil,"../samples/tests/solc/examples/strings.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}*/

func TestParseDao(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/examples/digixdao/dao.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseCoreWallet(t *testing.T) {
	ast, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/examples/digixdao/core_wallet.grd")
	goutil.Assert(t, ast != nil, "ast should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}

func TestParseGoldTxFeePool(t *testing.T) {
	p, errs := ValidateFile(NewTestVM(), nil, "../samples/tests/solc/examples/digixdao/gold_tx_fee_pool.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, errs == nil, errs.Format())
}
