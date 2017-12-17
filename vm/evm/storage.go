package evm

import (
	"encoding/binary"

	"github.com/end-r/vmgen"
)

type storageBlock struct {
	name   string
	size   uint
	offset uint
	slot   uint
}

type memoryBlock struct {
	size   uint
	offset uint
}

const (
	wordSize = uint(256)
)

func (s storageBlock) retrieve() (code vmgen.Bytecode) {
	if s.size > wordSize {
		// over at least 2 slots
		first := wordSize - s.offset
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, first))

		remaining := s.size - first
		slot := s.slot
		for remaining >= wordSize {
			// get whole slot
			slot += 1
			code.Concat(getByteSectionOfSlot(slot, 0, wordSize))
			remaining -= wordSize
		}
		if remaining > 0 {
			// get first remaining bits from next slot
			code.Concat(getByteSectionOfSlot(slot+1, 0, remaining))
		}
	} else if s.offset+s.size > wordSize {
		// over 2 slots
		// get last wordSize - s.offset bits from first
		first := wordSize - s.offset
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, first))
		// get first s.size - (wordSize - s.offset) bits from second
		code.Concat(getByteSectionOfSlot(s.slot+1, 0, s.size-first))
	} else {
		// all within 1 slot
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, s.size))
	}
	return code
}

func getByteSectionOfSlot(slot, start, size uint) (code vmgen.Bytecode) {
	code.Add("PUSH", uintAsBytes(slot)...)
	code.Add("SLOAD")
	mask := ((1 << size) - 1) << start
	code.Add("PUSH", uintAsBytes(uint(mask))...)
	code.Add("AND")
	return code
}

func uintAsBytes(a uint) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(a))
	return bs
}

func (m memoryBlock) retrieve() (code vmgen.Bytecode) {
	code.Concat(push(uintAsBytes(m.offset)))
	/*for size := m.size; size > wordSize; size -= wordSize {
		loc := m.offset + (m.size - size)
		code.Add("PUSH", uintAsBytes(loc)...)
		code.Add("MLOAD")
	}*/
	return code
}

func (s storageBlock) store() (code vmgen.Bytecode) {
	code.Concat(push(EncodeName(s.name)))
	code.Add("SSTORE")
	return code
}

func (m memoryBlock) store() (code vmgen.Bytecode) {
	free := []byte{0x40}
	code.Add("PUSH", free...)
	code.Add("PUSH", uintAsBytes(m.offset)...)
	code.Add("MSTORE")
	return code
}

func (evm *GuardianEVM) freeMemory(name string) {
	m, ok := evm.memory[name]
	if ok {
		if evm.freedMemory == nil {
			evm.freedMemory = make([]*memoryBlock, 0)
		}
		evm.freedMemory = append(evm.freedMemory, m)
		evm.memory[name] = nil
	}
}

func (evm *GuardianEVM) allocateMemory(name string, size uint) {
	if evm.memory == nil {
		evm.memory = make(map[string]*memoryBlock)
	}
	// try to use previously reclaimed memory
	if evm.freedMemory != nil {
		for i, m := range evm.freedMemory {
			if m.size >= size {
				// we can use this block
				evm.memory[name] = m
				// remove it from the freed list
				// TODO: check remove function
				evm.freedMemory = append(evm.freedMemory[i:], evm.freedMemory[:i]...)
				return
			}
		}
	}

	block := memoryBlock{
		size:   size,
		offset: evm.memoryCursor,
	}
	evm.memoryCursor += size
	evm.memory[name] = &block
}

func (evm *GuardianEVM) allocateStorage(name string, size uint) {
	// TODO: check whether there's a way to reduce storage using some weird bin packing algo
	// with a modified heuristic to reduce the cost of extracting variables using bitshifts
	// maybe?
	if evm.storage == nil {
		evm.storage = make(map[string]*storageBlock)
	}
	block := storageBlock{
		size:   size,
		offset: evm.lastOffset,
		slot:   evm.lastSlot,
	}
	for size > wordSize {
		size -= wordSize
		evm.lastSlot += 1
	}
	evm.lastOffset += size
	evm.storage[name] = &block
}

func (evm *GuardianEVM) lookupStorage(name string) *storageBlock {
	if evm.storage == nil {
		return nil
	}
	return evm.storage[name]
}

func (evm *GuardianEVM) lookupMemory(name string) *memoryBlock {
	if evm.memory == nil {
		return nil
	}
	return evm.memory[name]
}
