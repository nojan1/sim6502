package sim6502

type anc struct{}

func (i *anc) Mnemonic() string {
	return "ANC"
}

func (i *anc) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	// AND operation
	proc.registers.A = proc.registers.A & data

	// Bit 7 goes to carry
	proc.registers.SR.SetTo(SRFlagC, proc.registers.A&0x80 == 0x80)
	proc.registers.SR.setNZ(proc.registers.A)

	return nil
}
