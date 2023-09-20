package sim6502

type plx struct{}

func (i *plx) Mnemonic() string {
	return "PLX"
}

func (i *plx) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.X = proc.registers.SP.Pop()
	proc.registers.SR.setNZ(proc.registers.X)
	return nil
}
