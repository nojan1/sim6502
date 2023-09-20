package sim6502

import "errors"

type jmp struct{}

func (i *jmp) Mnemonic() string {
	return "JMP"
}

func (i *jmp) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
		return errors.New("self jump on JMP")
	}

	proc.registers.PC.Set(data16)
	return nil
}
