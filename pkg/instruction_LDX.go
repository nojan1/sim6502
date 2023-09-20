package sim6502

type ldx struct {
	standardCrossPage
}

func (i *ldx) Mnemonic() string {
	return "LDX"
}

func (i *ldx) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.X = data
	proc.registers.SR.setNZ(data)
	return nil
}
