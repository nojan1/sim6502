package sim6502

type slo struct{}

func (i *slo) Mnemonic() string {
	return "SLO"
}

func (i *slo) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	// ASL operation
	carry := data&0x80 > 0
	result := data << 1
	proc.registers.SR.SetTo(SRFlagC, carry)
	proc.memory.Write(data16, result)

	// ORA operation
	proc.registers.A = proc.registers.A | data

	// Presumably the flags result from the ORA, except carry above from the ASL
	proc.registers.SR.setNZ(proc.registers.A)
	return nil
}
