package sim6502

type alr struct{}

func (i *alr) Mnemonic() string {
	return "ALR"
}

func (i *alr) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	// AND operation
	proc.registers.A = proc.registers.A & data

	// LSR operation
	proc.registers.SR.SetTo(SRFlagC, proc.registers.A&0x10 == 0x10)
	proc.registers.A = proc.registers.A >> 1
	proc.registers.SR.setNZ(proc.registers.A)
	return nil
}
