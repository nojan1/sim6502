package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// For the extended instruction set, there should be an instruction defined at every location

func TestInstructionSet(t *testing.T) {
	assert := assert.New(t)
	instrs := make([]*instruction, len(instructions))
	copy(instrs, instructions)

	loadExtendedInstructions(instrs)

	for opcode, i := range instrs {
		assert.NotNil(i, fmt.Sprintf("No instruction found at opcode 0x%02x", opcode))
	}

}
