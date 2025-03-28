package sim6502

import (
	"fmt"
	"io"
	"reflect"
)

// MappedMemoryAddressRange provides an address range used for mapping memory
// Start must be <= End
type MappedMemoryAddressRange struct {
	Start, End uint16
}

// MappedMemoryHandler represents an implementation of a memory mapped device
type MappedMemoryHandler interface {
	// AddressRange reports the memory mapped address range supported by the device
	AddressRange() []MappedMemoryAddressRange
}

// MappedMemoryWriteHandler is a mapped device that received writes from the CPU to it's memory location(s)
type MappedMemoryWriteHandler interface {
	MappedMemoryHandler
	// Write is called when the CPU writes to the specified address
	Write(addr uint16, val uint8)
}

// MappedMemoryReadHandler is a mapped device that receives provides data to the CPU when it's memory location(s) are read
type MappedMemoryReadHandler interface {
	MappedMemoryHandler
	// Read is called when the CPU reads the specified
	Read(addr uint16, internal bool) uint8
}

// MappableMemory is a raw 64K address space in which you may map additional handlers
// to provide either read or write functionality at specific addresses
type MappableMemory struct {
	data              [65536]uint8
	device            [65536]MappedMemoryHandler
	debugWrites       bool
	debugReads        bool
	debugWritesWriter io.Writer
	debugReadsWriter  io.Writer
}

// EnableReadDebugging will report every read to the provided writer
func (m *MappableMemory) EnableReadDebugging(w io.Writer) {
	m.debugReadsWriter = w
	m.debugReads = true
}

// EnableWriteDebugging will report every write to the provided writer
func (m *MappableMemory) EnableWriteDebugging(w io.Writer) {
	m.debugWritesWriter = w
	m.debugWrites = true
}

// DisableReadDebugging will disable debugging of reads
func (m *MappableMemory) DisableReadDebugging() {
	m.debugReads = false
}

// DisableWriteDebugging will disable debugging of writes
func (m *MappableMemory) DisableWriteDebugging() {
	m.debugWrites = false
}

// Clear clears the memory
func (m *MappableMemory) Clear() {
	for i := 0; i < len(m.data); i++ {
		m.data[i] = 0x00
	}
}

// Map will map a new handler into the memory space
// The MappedMemoryHandler must implement either or both of
// MappedMemoryReadHandler and/or MappedMemoryWriteHandler
// otherwise it won't do anything
func (m *MappableMemory) Map(handler MappedMemoryHandler) {
	for _, ar := range handler.AddressRange() {
		for i := ar.Start; i <= ar.End; i++ {
			m.device[i] = handler
		}
	}
}

// Write memory
func (m *MappableMemory) Write(location uint16, value uint8) {
	// Is this memory mapped for output

	if oh, ok := m.device[location].(MappedMemoryWriteHandler); ok {
		if m.debugWrites {
			fmt.Fprintf(m.debugWritesWriter, "Memory Write (Mapped: %s): $%02x -> $%04x\n", reflect.TypeOf(oh).Name(), value, location)
		}
		oh.Write(location, value)
		return
	}

	// Nah, just raw memory
	if m.debugWrites {
		fmt.Fprintf(m.debugWritesWriter, "Memory Write (Raw): $%02x -> $%04x\n", value, location)
	}
	m.data[location] = value
}

// Read memory
func (m *MappableMemory) Read(location uint16, internal bool) uint8 {
	// Is this memory mapped for input
	if ih, ok := m.device[location].(MappedMemoryReadHandler); ok {
		val := ih.Read(location, internal)
		if m.debugReads {
			fmt.Fprintf(m.debugReadsWriter, "Memory Read  (Mapped: %s): $%02x <- $%04x\n", reflect.TypeOf(ih).Name(), val, location)
		}

		return val
	}

	// Nah, just raw memory
	if m.debugReads {
		fmt.Fprintf(m.debugReadsWriter, "Memory Read  (Raw): $%02x <- $%04x\n", m.data[location], location)
	}

	return m.data[location]
}
