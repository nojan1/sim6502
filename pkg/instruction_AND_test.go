package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstructionAND(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	proc := prepareInstructionImmediate(t, instructions, "AND", 0xF1)
	proc.registers.A = 0x2F
	err, _ := proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(0x21, proc.registers.A)
}
