package sim6502

type rti struct{}

func (i *rti) Mnemonic() string {
	return "RTI"
}

func (i *rti) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.value = proc.registers.SP.Pop() & 0xDF
	pcl := proc.registers.SP.Pop()
	pch := proc.registers.SP.Pop()
	pc := ((uint16(pch) << 8) | uint16(pcl))

	proc.registers.PC.Set(pc)
	return nil
}
