package sim6502

type ora struct {
	standardCrossPage
}

func (i *ora) Mnemonic() string {
	return "ORA"
}

func (i *ora) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.A = proc.registers.A | data
	proc.registers.SR.setNZ(proc.registers.A)
	return nil
}
