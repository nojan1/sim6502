package sim6502

// Memory is the interface that must be implemented by any memory provider
// The memory provider should implement the entire 64K address space as it sees fit
type Memory interface {
	// Read reads a byte from the specified address
	Read(addr uint16) uint8

	// Write writes a byte to the specified address
	Write(addr uint16, value uint8)

	// Clear must clear all memory to zero value
	Clear()
}
