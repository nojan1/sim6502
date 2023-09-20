package sim6502

type dex struct{}

func (i *dex) Mnemonic() string {
	return "DEX"
}

func (i *dex) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.X--
	proc.registers.SR.setNZ(proc.registers.X)
	return nil
}
