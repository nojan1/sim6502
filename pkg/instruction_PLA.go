package sim6502

type pla struct{}

func (i *pla) Mnemonic() string {
	return "PLA"
}

func (i *pla) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.A = proc.registers.SP.Pop()
	proc.registers.SR.setNZ(proc.registers.A)
	return nil
}
