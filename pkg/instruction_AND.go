package sim6502

type and struct {
	standardCrossPage
}

func (i *and) Mnemonic() string {
	return "AND"
}

func (i *and) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	result := proc.registers.A & data
	proc.registers.SR.setNZ(result)
	proc.registers.A = result
	return nil
}
