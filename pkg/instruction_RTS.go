package sim6502

type rts struct{}

func (i *rts) Mnemonic() string {
	return "RTS"
}

func (i *rts) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	pcl := proc.registers.SP.Pop()
	pch := proc.registers.SP.Pop()
	pc := (uint16(pch) << 8) | uint16(pcl)

	proc.registers.PC.Set(pc + 1)
	return nil
}
