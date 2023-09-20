package sim6502

type las struct {
	standardCrossPage
}

func (i *las) Mnemonic() string {
	return "LAS"
}

func (i *las) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	result := proc.registers.A & proc.registers.SP.ptr

	proc.registers.A = result
	proc.registers.X = result
	proc.registers.SP.SetStackPointer(result)
	proc.registers.SR.setNZ(result)

	return nil
}
