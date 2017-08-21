package ast

func checkMnemonics(t *testing.T, is []*instruction, es []string){
    goutil.AssertNow(t, len(is) == len(es, "wrong num of instructions")
	for index, i := range is {
		goutil.Assert(t, i.Mnemonic == es[index],
			fmt.Sprintf("wrong mnemonic %d: %s, expected %s", index, is.Mnemonic, es[index]))
	}
}
