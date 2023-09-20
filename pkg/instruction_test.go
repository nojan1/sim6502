package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstructionSetOpcodeSequence(t *testing.T) {
	assert := assert.New(t)
	for idx, op := range instructions {
		if op != nil {
			assert.EqualValues(idx, op.OpCode, fmt.Sprintf("Opcode at arr# 0x%02x mismatch, value 0x%02x", idx, op.OpCode))
		}
	}
}

func TestInstructionSetSize(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(256, len(instructions), "Should be exactly 256 entries in instruction map")
}

func TestDuplicateInstructions(t *testing.T) {

	eq := func(a, b *instruction) bool {
		return a.Impl.Mnemonic() == b.Impl.Mnemonic() && a.AddressingMode == b.AddressingMode
	}

	assert := assert.New(t)
	for i := 0; i < 255; i++ {
		for n := i + 1; n < 256; n++ {
			if instructions[i] != nil && instructions[n] != nil {
				assert.False(eq(instructions[i], instructions[n]), fmt.Sprintf("Duplicate instructions at indexes 0x%02x and 0x%02x", i, n))
			}
		}
	}
}
