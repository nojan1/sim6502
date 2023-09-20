package sim6502

type eor struct {
	standardCrossPage
}

func (i *eor) Mnemonic() string {
	return "EOR"
}

func (i *eor) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.A = proc.registers.A ^ data
	proc.registers.SR.setNZ(proc.registers.A)
	return nil
}
