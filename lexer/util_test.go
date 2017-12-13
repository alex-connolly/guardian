package lexer

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/token"
)

func checkTokens(t *testing.T, received []token.Token, expected []token.Type) {
	goutil.AssertNow(t, len(received) == len(expected), fmt.Sprintf("wrong num of tokens: a %d / e %d", len(received), len(expected)))
	for index, r := range received {
		goutil.Assert(t, r.Type == expected[index],
			fmt.Sprintf("wrong type %d: %s, expected %d", index, r.Proto.Name, expected[index]))
	}
}
