package sim6502

import (
	"fmt"
	"strings"
)

// StackDebug can be set to true to output messages when stack operations occur
// strictly for debugging
var StackDebug = false

// StackLocation points to the base of the stack at 0x0100
// N.B. Stack grows down from PTR=0xff (i.e. 0x01FF)
const StackLocation uint16 = 0x0100 // Base of stack

// StackPointer is our implementation of the stack pointer
type StackPointer struct {
	//ptr is the current stack pointer value
	ptr uint8

	// we must be initialized with a pointer to the memory, since the stack
	// is stored in that memory
	mem Memory
}

// Push pushes a byte to the stack
func (sp *StackPointer) Push(value uint8) {
	if StackDebug {
		fmt.Printf("Stack DBG: Push 0x%02x to SP 0x%02x at addr %04x\n", value, sp.ptr, StackLocation+uint16(sp.ptr))
	}
	sp.mem.Write(StackLocation+uint16(sp.ptr), value)
	sp.ptr--
}

// Pop pops a byte off the stack
func (sp *StackPointer) Pop() uint8 {
	sp.ptr++
	result := sp.mem.Read(StackLocation + uint16(sp.ptr), true)
	if StackDebug {
		fmt.Printf("Stack DBG: Pop 0x%02x from SP 0x%02x at addr %04x\n", result, sp.ptr, StackLocation+uint16(sp.ptr))
	}
	return result
}

// SetStackPointer is used to explicitly set the stack pointer
func (sp *StackPointer) SetStackPointer(ptr uint8) {
	sp.ptr = ptr
}

// GetStackPointer retrieves the current stack pointer value
func (sp *StackPointer) GetStackPointer() uint8 {
	return sp.ptr
}

// PeekStackHead returns the byte on the stack head without altering the stack
func (sp *StackPointer) PeekStackHead() uint8 {
	return sp.mem.Read(StackLocation + uint16(sp.ptr+1), true)
}

// String provides a string dump of the stack
func (sp *StackPointer) String() string {
	sb := strings.Builder{}
	sep := ""
	sb.WriteString("[STACK: ")

	for i := sp.ptr + 1; i != 0; i++ {
		sb.WriteString(sep)
		sb.WriteString(fmt.Sprintf("%02x", sp.mem.Read(StackLocation+uint16(i), true)))
		sep = ", "
	}
	sb.WriteRune(']')
	return sb.String()
}
