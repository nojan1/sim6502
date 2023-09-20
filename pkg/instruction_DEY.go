package sim6502

type dey struct{}

func (i *dey) Mnemonic() string {
	return "DEY"
}

func (i *dey) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.Y--
	proc.registers.SR.setNZ(proc.registers.Y)
	return nil
}
