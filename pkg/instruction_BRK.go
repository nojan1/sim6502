package sim6502

type brk struct{}

func (i *brk) Mnemonic() string {
	return "BRK"
}

func (i *brk) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	_ = proc.registers.PC.Next() // The break mark, ignored

	ret := proc.registers.PC.Current() - 1
	pcl := uint8(ret & 0xff)
	pch := uint8(ret >> 8)

	proc.registers.SP.Push(pch)
	proc.registers.SP.Push(pcl)
	proc.registers.SP.Push(proc.registers.SR.value | 0x30) // Set break and unused flag when pushing

	if proc.model == ProcessorModel65C02 {
		// 65C02 clears the decimal flag on BRK (and any other interrupt)
		proc.registers.SR.Clear(SRFlagD)
	}

	proc.registers.SR.Set(SRFlagI) // Set interrupt flag
	proc.registers.PC.Set(uint16(proc.memory.Read(uint16(VectorIRQ))) | (uint16(proc.memory.Read(uint16(VectorIRQ+1))) << 8))
	return nil
}
