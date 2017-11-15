package evm

func EncodeName(name string) []byte {
	/*
		// ethereum uses the 'original keccak' (disbyte of 1 rather than 6)
		hasher := &state{rate: 136, outputLen: 32, dsbyte: 0x01}

		hasher.Write([]byte(name))

		// get the output hash
		hash := hasher.Sum(nil)
		// take the leftmost 4 bytes of the hash
		limited := hash[:4]

		return limited
	*/
	return []byte(name)
}
