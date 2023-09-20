package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// For the illegal instruction set, there should be an instruction defined at every location

func TestInstructionSet(t *testing.T) {
	assert := assert.New(t)
	instrs := make([]*instruction, len(instructions))
	copy(instrs, instructions)

	loadIllegalInstructions(instrs)

	for opcode, i := range instrs {
		assert.NotNil(i, fmt.Sprintf("No instruction found at opcode 0x%02x", opcode))
	}
}

// Ditto for the 65C02 instruction set
func Test65C02(t *testing.T) {
	assert := assert.New(t)
	instrs := make([]*instruction, len(instructions))
	copy(instrs, instructions)

	load85C02instructions(instrs)

	for opcode, i := range instrs {
		assert.NotNil(i, fmt.Sprintf("No instruction found at opcode 0x%02x", opcode))
	}
}
