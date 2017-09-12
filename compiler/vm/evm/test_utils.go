package evm

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/vmgen"
)

func checkMnemonics(t *testing.T, is map[byte]*vmgen.Instruction, es []string) {
	goutil.AssertNow(t, is != nil, "instructions shouldn't be nil")
	goutil.AssertNow(t, len(is) == len(es), "wrong num of instructions")
	for index, i := range is {
		goutil.Assert(t, i.Mnemonic == es[index],
			fmt.Sprintf("wrong mnemonic %d: %s, expected %s", index, i.Mnemonic, es[index]))
	}
}
