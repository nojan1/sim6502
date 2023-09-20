package sim6502

type inx struct{}

func (i *inx) Mnemonic() string {
	return "INX"
}

func (i *inx) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.X++
	proc.registers.SR.setNZ(proc.registers.X)
	return nil
}
