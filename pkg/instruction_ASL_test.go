package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstructionASLAccumulatorMode(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	proc := prepareInstructionAcc(t, instructions, "ASL")
	proc.registers.A = 0x40
	err, _ := proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(0x80, proc.registers.A)
	assert.False(proc.registers.SR.IsSet(SRFlagC))
	assert.False(proc.registers.SR.IsSet(SRFlagZ))
	assert.True(proc.registers.SR.IsSet(SRFlagN))

	proc = prepareInstructionAcc(t, instructions, "ASL")
	proc.registers.A = 0x80
	err, _ = proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(0x00, proc.registers.A)
	assert.True(proc.registers.SR.IsSet(SRFlagC))
	assert.True(proc.registers.SR.IsSet(SRFlagZ))
	assert.False(proc.registers.SR.IsSet(SRFlagN))

	proc = prepareInstructionAcc(t, instructions, "ASL")
	proc.registers.A = 0x00
	err, _ = proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(0x00, proc.registers.A)
	assert.False(proc.registers.SR.IsSet(SRFlagC))
	assert.True(proc.registers.SR.IsSet(SRFlagZ))
	assert.False(proc.registers.SR.IsSet(SRFlagN))
}

func TestInstructionASLAbsoluteMode(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	var addr uint16 = 0x04000

	proc := prepareInstructionAbsolute(t, instructions, "ASL", addr)
	proc.memory.Write(addr, 0x40)
	err, _ := proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(0x80, proc.memory.Read(addr, true))
	assert.False(proc.registers.SR.IsSet(SRFlagC))
	assert.False(proc.registers.SR.IsSet(SRFlagZ))
	assert.True(proc.registers.SR.IsSet(SRFlagN))

	proc = prepareInstructionAbsolute(t, instructions, "ASL", addr)
	proc.memory.Write(addr, 0x80)
	err, _ = proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(0x00, proc.memory.Read(addr, true))

	assert.True(proc.registers.SR.IsSet(SRFlagC))
	assert.True(proc.registers.SR.IsSet(SRFlagZ))
	assert.False(proc.registers.SR.IsSet(SRFlagN))

	proc = prepareInstructionAbsolute(t, instructions, "ASL", addr)
	proc.memory.Write(addr, 0x00)
	err, _ = proc.Step()
	require.Nil(err, fmt.Sprintf("error on step: %v", err))
	assert.EqualValues(0x00, proc.memory.Read(addr, true))
	assert.False(proc.registers.SR.IsSet(SRFlagC))
	assert.True(proc.registers.SR.IsSet(SRFlagZ))
	assert.False(proc.registers.SR.IsSet(SRFlagN))
}
