package sim6502

type TAY struct{}

func (i *TAY) Mnemonic() string {
	return "TAY"
}

func (i *TAY) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.Y = proc.registers.A
	proc.registers.SR.setNZ(proc.registers.Y)
	return nil
}
