package parser

import (
	"testing"

	"github.com/end-r/guardian/ast"
)

func BenchmarkParseTypeDeclarationBuiltins(b *testing.B) {
	var a *ast.ScopeNode
	for n := 0; n < b.N; n++ {
		a, _ = ParseString(`
            type string []byte
            type address [20]byte

            class BuiltinMessage {
                data []byte
                gas uint
                sender address
                sig [4]byte
            }

            class BuiltinBlock {
                timestamp uint
                number uint
                coinbase address
                gaslimit uint
                blockhash func(blockNumber uint) [32]byte
            }

            class BuiltinTransaction {
                gasprice uint
                origin address
            }
        `)
	}
	result = a
}

func BenchmarkParseAssignmentBuiltins(b *testing.B) {
	var a *ast.ScopeNode
	for n := 0; n < b.N; n++ {
		a, _ = ParseString(`
            wei = 1
            kwei = 1000 * wei
            babbage = kwei
            mwei = 1000 * kwei
            lovelace = mwei
            gwei = 1000 * mwei
            shannon = gwei
            microether = 1000 * gwei
            szabo = microether
            milliether = 1000 * microether
            finney = milliether
            ether = 1000 * milliether
        `)
	}
	result = a
}

func BenchmarkParseFuncVarDeclarationBuiltins(b *testing.B) {
	var a *ast.ScopeNode
	for n := 0; n < b.N; n++ {
		a, _ = ParseString(`

            balance func(a address) uint256
            transfer func(a address, amount uint256) uint
            send func(a address, amount uint256) bool
            call func(a address) bool
            delegateCall func(a address)

            addmod func(x, y, k uint) uint
            mulmod func(x, y, k uint) uint
            keccak256 func()
            sha256 func()
            sha3 func()
            ripemd160 func()
            ecrecover func (v uint8, h, r, s bytes32) address
            selfDestruct func(recipient address) uint256

        `)
	}
	result = a
}

func BenchmarkParseVarDeclarationBuiltins(b *testing.B) {
	var a *ast.ScopeNode
	for n := 0; n < b.N; n++ {
		a, _ = ParseString(`
            msg BuiltinMessage
            block BuiltinBlock
            tx BuiltinTransaction
        `)
	}
	result = a
}

func BenchmarkParseFullBuiltins(b *testing.B) {
	var a *ast.ScopeNode
	for n := 0; n < b.N; n++ {
		a, _ = ParseString(`
            type string []byte
            type address [20]byte

            wei = 1
            kwei = 1000 * wei
            babbage = kwei
            mwei = 1000 * kwei
            lovelace = mwei
            gwei = 1000 * mwei
            shannon = gwei
            microether = 1000 * gwei
            szabo = microether
            milliether = 1000 * microether
            finney = milliether
            ether = 1000 * milliether

            // account functions
            balance func(a address) uint256
            transfer func(a address, amount uint256) uint
            send func(a address, amount uint256) bool
            call func(a address) bool
            delegateCall func(a address)

            // cryptographic functions
            addmod func(x, y, k uint) uint
            mulmod func(x, y, k uint) uint
            keccak256 func()
            sha256 func()
            sha3 func()
            ripemd160 func()
            ecrecover func (v uint8, h, r, s bytes32) address

            // contract functions
            // NO THIS KEYWORD: confusing for most programmers, unintentional bugs etc

            selfDestruct func(recipient address) uint256


            class BuiltinMessage {
                data []byte
                gas uint
                sender address
                sig [4]byte
            }

            msg BuiltinMessage

            class BuiltinBlock {
                timestamp uint
                number uint
                coinbase address
                gaslimit uint
                blockhash func(blockNumber uint) [32]byte
            }

            block BuiltinBlock

            class BuiltinTransaction {
                gasprice uint
                origin address
            }

            tx BuiltinTransaction
        `)
	}
	result = a
}
