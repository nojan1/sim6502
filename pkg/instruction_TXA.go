package sim6502

type TXA struct{}

func (i *TXA) Mnemonic() string {
	return "TXA"
}

func (i *TXA) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.A = proc.registers.X
	proc.registers.SR.setNZ(proc.registers.A)
	return nil
}
