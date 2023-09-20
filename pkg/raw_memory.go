package sim6502

import (
	"fmt"
	"io"
)

// RawMemory implements a simple R/W 64K address space with no special features
type RawMemory struct {
	data              [65536]uint8
	debugWrites       bool
	debugReads        bool
	debugWritesWriter io.Writer
	debugReadsWriter  io.Writer
}

// EnableReadDebugging will report every read to the provided writer
func (m *RawMemory) EnableReadDebugging(w io.Writer) {
	m.debugReadsWriter = w
	m.debugReads = true
}

// EnableWriteDebugging will report every write to the provided writer
func (m *RawMemory) EnableWriteDebugging(w io.Writer) {
	m.debugWritesWriter = w
	m.debugWrites = true
}

// DisableReadDebugging will disable debugging of reads
func (m *RawMemory) DisableReadDebugging() {
	m.debugReads = false
}

// DisableWriteDebugging will disable debugging of writes
func (m *RawMemory) DisableWriteDebugging() {
	m.debugWrites = false
}

// Clear clears all memory
func (m *RawMemory) Clear() {
	for i := 0; i < len(m.data); i++ {
		m.data[i] = 0x00
	}
}

// Write a memory location
func (m *RawMemory) Write(location uint16, value uint8) {
	if m.debugWrites {
		fmt.Fprintf(m.debugWritesWriter, "Memory Write (Raw): $%02x -> $%04x\n", value, location)
	}
	m.data[location] = value
}

// Read a memory location
func (m *RawMemory) Read(location uint16) uint8 {
	if m.debugReads {
		fmt.Fprintf(m.debugReadsWriter, "Memory Read  (Raw): $%02x <- $%04x\n", m.data[location], location)
	}

	return m.data[location]
}
