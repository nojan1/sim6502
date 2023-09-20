package sim6502

type tas struct{}

func (i *tas) Mnemonic() string {
	return "TAS"
}

func (i *tas) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SP.SetStackPointer(proc.registers.X & proc.registers.Y)
	proc.memory.Write(data16, (proc.registers.A&proc.registers.X&uint8(data16>>8))+1)
	return nil
}
