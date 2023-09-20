package sim6502

type iny struct{}

func (i *iny) Mnemonic() string {
	return "INY"
}

func (i *iny) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.Y++
	proc.registers.SR.setNZ(proc.registers.Y)
	return nil
}
