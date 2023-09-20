package sim6502

type ply struct{}

func (i *ply) Mnemonic() string {
	return "PLY"
}

func (i *ply) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.Y = proc.registers.SP.Pop()
	proc.registers.SR.setNZ(proc.registers.Y)
	return nil
}
