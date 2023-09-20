package sim6502

type TYA struct{}

func (i *TYA) Mnemonic() string {
	return "TYA"
}

func (i *TYA) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.A = proc.registers.Y
	proc.registers.SR.setNZ(proc.registers.A)
	return nil
}
