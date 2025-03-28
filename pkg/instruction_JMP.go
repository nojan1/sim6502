package sim6502

import "errors"

type jmp struct{}

func (i *jmp) Mnemonic() string {
	return "JMP"
}

func (i *jmp) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	// On NMOS 6502, JMP indirect is slightly broken
	if !proc.fixBrokenJMP && mode == IND && proc.model == ProcessorModel6502 {
		// Grab the two operands again
		// low, high bytes of indirect address
		lowind := proc.memory.Read(proc.registers.PC.Current() - 2, false)
		highind := proc.memory.Read(proc.registers.PC.Current() - 1, false)

		if lowind == 0xff {
			// Breakage mode is when the low byte of the indirect address is 0xff
			// This is the indirect address
			ind := uint16(highind)<<8 | uint16(lowind)
			lowaddr := proc.memory.Read(ind, false)
			highaddr := proc.memory.Read(ind & 0xff00, false) // instead of ind+1
			proc.registers.PC.Set(uint16(highaddr)<<8 | uint16(lowaddr))
			return nil
		}
	}

	if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
		return errors.New("self jump on JMP")
	}

	proc.registers.PC.Set(data16)
	return nil
}
