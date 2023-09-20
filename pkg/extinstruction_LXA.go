package sim6502

type lxa struct{}

func (i *lxa) Mnemonic() string {
	return "LXA"
}

func (i *lxa) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	result := (ExtendedInstructionANEConstant | proc.registers.A) & data
	proc.registers.A = result
	proc.registers.X = result
	proc.registers.SR.setNZ(result)
	return nil
}
