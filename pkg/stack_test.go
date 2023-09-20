package sim6502

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackNominalOperation(t *testing.T) {

	assert := assert.New(t)
	proc := NewProcessor(&RawMemory{})
	proc.registers.SP.ptr = 0xff

	// Pushes the values 0..255 onto the stack
	for i := 0; i < 256; i++ {
		proc.registers.SP.Push(uint8(i))
	}

	// Checks that the values are in the proper stack locations
	for i := 0xff; i >= 0; i-- {
		expected := 0xff - i
		assert.EqualValues(expected, proc.memory.Read(StackLocation+uint16(i)))
	}

	// Pops them off and validates them
	for i := 255; i >= 0; i-- {
		val := proc.registers.SP.Pop()
		assert.EqualValues(i, val)
	}
}

func TestStackStringifier(t *testing.T) {
	assert := assert.New(t)
	proc := NewProcessor(&RawMemory{})

	proc.registers.SP.ptr = 0xff

	proc.registers.SP.Push(0x55)
	proc.registers.SP.Push(0x66)
	proc.registers.SP.Push(0x77)
	assert.Equal("[STACK: 77, 66, 55]", proc.registers.SP.String())
}
