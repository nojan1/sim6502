package sim6502

type TSX struct{}

func (i *TSX) Mnemonic() string {
	return "TSX"
}

func (i *TSX) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.X = proc.registers.SP.GetStackPointer()
	proc.registers.SR.setNZ(proc.registers.X)
	return nil
}
