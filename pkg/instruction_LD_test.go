package sim6502

// Tests LDA, LDX, LDY operations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLDA(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	proc := prepareInstructionImmediate(t, instructions, "LDA", 8)
	err, _ := proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(8, proc.registers.A)
}

func TestLDX(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	proc := prepareInstructionImmediate(t, instructions, "LDX", 8)
	err, _ := proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(8, proc.registers.X)
}

func TestLDY(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	proc := prepareInstructionImmediate(t, instructions, "LDX", 8)
	err, _ := proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(8, proc.registers.X)
}
