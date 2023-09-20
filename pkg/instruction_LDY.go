package sim6502

type ldy struct {
	standardCrossPage
}

func (i *ldy) Mnemonic() string {
	return "LDY"
}

func (i *ldy) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.Y = data
	proc.registers.SR.setNZ(data)
	return nil
}
