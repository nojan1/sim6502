package sim6502

type TAX struct{}

func (i *TAX) Mnemonic() string {
	return "TAX"
}

func (i *TAX) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.X = proc.registers.A
	proc.registers.SR.setNZ(proc.registers.X)
	return nil
}
