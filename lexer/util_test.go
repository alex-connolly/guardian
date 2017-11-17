package lexer

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func checkTokens(t *testing.T, received []Token, expected []TokenType) {
	goutil.AssertNow(t, len(received) == len(expected), fmt.Sprintf("wrong num of tokens: a %d / e %d", len(received), len(expected)))
	for index, r := range received {
		goutil.Assert(t, r.Type == expected[index],
			fmt.Sprintf("wrong type %d: %s, expected %d", index, r.proto.name, expected[index]))
	}
}
