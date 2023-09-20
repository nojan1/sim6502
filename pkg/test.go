package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCodeOffset uint16 = 0x1020
)

func findInstruction(instructions []*instruction, mnemonic string, mode AddressingMode) instruction {
	for _, i := range instructions {
		if i != nil {
			if i.AddressingMode == mode && i.Impl.Mnemonic() == mnemonic {
				return *i
			}
		}
	}
	panic(fmt.Sprintf("Failed to find opcode for %s, mode %s", mnemonic, mode.String()))
}

func setWord(p *Processor, addr uint16, word uint16) {
	p.memory.Write(addr, uint8(word&0xff))
	p.memory.Write(addr+1, uint8(word>>8))
}

func prepareInstructionImmediate(t *testing.T, instructions []*instruction, mnemonic string, value uint8) *Processor {
	require := require.New(t)
	instruction := findInstruction(instructions, mnemonic, IMMED)
	require.NotNil(instruction, fmt.Sprintf("Instruction %s does not have an immediate implementation", mnemonic))

	proc := NewProcessor(&RawMemory{})

	// Set up the reset vector to point at where we're going to store our code
	setWord(proc, uint16(VectorReset), testCodeOffset)
	require.EqualValues(testCodeOffset&0xff, proc.memory.Read(uint16(VectorReset)), "PCL not properly set")
	require.EqualValues(testCodeOffset>>8, proc.memory.Read(uint16(VectorReset+1)), "PCH not properly set")

	// Store our code at the desginated location
	proc.memory.Write(testCodeOffset, instruction.OpCode)
	proc.memory.Write(testCodeOffset+1, value)

	// Initialize program counter
	proc.setPCFromResetVector()

	// Processor is ready to go
	return proc
}

func prepareInstructionRelative(t *testing.T, instructions []*instruction, mnemonic string, value uint8) *Processor {
	require := require.New(t)
	instruction := findInstruction(instructions, mnemonic, REL)
	require.NotNil(instruction, fmt.Sprintf("Instruction %s does not have an relative implementation", mnemonic))

	proc := NewProcessor(&RawMemory{})

	// Set up the reset vector to point at where we're going to store our code
	setWord(proc, uint16(VectorReset), testCodeOffset)
	require.EqualValues(testCodeOffset&0xff, proc.memory.Read(uint16(VectorReset)), "PCL not properly set")
	require.EqualValues(testCodeOffset>>8, proc.memory.Read(uint16(VectorReset+1)), "PCH not properly set")

	// Store our code at the desginated location
	proc.memory.Write(testCodeOffset, instruction.OpCode)
	proc.memory.Write(testCodeOffset+1, value)

	// Initialize program counter
	proc.setPCFromResetVector()

	// Processor is ready to go
	return proc
}

func prepareInstructionAcc(t *testing.T, instructions []*instruction, mnemonic string) *Processor {
	require := require.New(t)
	instruction := findInstruction(instructions, mnemonic, A)
	require.NotNil(instruction, fmt.Sprintf("Instruction %s does not have an accumulator implementation", mnemonic))

	proc := NewProcessor(&RawMemory{})

	// Set up the reset vector to point at where we're going to store our code
	setWord(proc, uint16(VectorReset), testCodeOffset)
	require.EqualValues(testCodeOffset&0xff, proc.memory.Read(uint16(VectorReset)), "PCL not properly set")
	require.EqualValues(testCodeOffset>>8, proc.memory.Read(uint16(VectorReset+1)), "PCH not properly set")

	// Store our code at the desginated location
	proc.memory.Write(testCodeOffset, instruction.OpCode)

	// Initialize program counter
	proc.setPCFromResetVector()

	// Processor is ready to go
	return proc
}

func prepareInstructionImplied(t *testing.T, instructions []*instruction, mnemonic string) *Processor {
	require := require.New(t)
	instruction := findInstruction(instructions, mnemonic, IMPL)
	require.NotNil(instruction, fmt.Sprintf("Instruction %s does not have an implied implementation", mnemonic))

	proc := NewProcessor(&RawMemory{})

	// Set up the reset vector to point at where we're going to store our code
	setWord(proc, uint16(VectorReset), testCodeOffset)
	require.EqualValues(testCodeOffset&0xff, proc.memory.Read(uint16(VectorReset)), "PCL not properly set")
	require.EqualValues(testCodeOffset>>8, proc.memory.Read(uint16(VectorReset+1)), "PCH not properly set")

	// Store our code at the desginated location
	proc.memory.Write(testCodeOffset, instruction.OpCode)

	// Initialize program counter
	proc.setPCFromResetVector()

	// Processor is ready to go
	return proc
}

func prepareInstructionAbsolute(t *testing.T, instructions []*instruction, mnemonic string, addr uint16) *Processor {
	require := require.New(t)
	instruction := findInstruction(instructions, mnemonic, ABS)
	require.NotNil(instruction, fmt.Sprintf("Instruction %s does not have an absolute implementation", mnemonic))

	proc := NewProcessor(&RawMemory{})

	// Set up the reset vector to point at where we're going to store our code
	setWord(proc, uint16(VectorReset), testCodeOffset)
	require.EqualValues(testCodeOffset&0xff, proc.memory.Read(uint16(VectorReset)), "PCL not properly set")
	require.EqualValues(testCodeOffset>>8, proc.memory.Read(uint16(VectorReset+1)), "PCH not properly set")

	// Store our code at the desginated location
	proc.memory.Write(testCodeOffset, instruction.OpCode)
	proc.memory.Write(testCodeOffset+1, uint8(addr&0xff))
	proc.memory.Write(testCodeOffset+2, uint8(addr>>8))

	// Initialize program counter
	proc.setPCFromResetVector()

	// Processor is ready to go
	return proc
}
