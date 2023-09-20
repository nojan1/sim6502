package sim6502

// Registers defines the structure containing all processor registers
type Registers struct {
	// A is the accumulator
	A uint8

	// X is the X register
	X uint8

	// Y is the Y register
	Y uint8

	// SR is the status register
	SR StatusRegister

	// SP is the stack pointer
	SP StackPointer

	// PC is the program counter
	PC ProgramCounter
}
