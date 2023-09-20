package sim6502

type plp struct{}

func (i *plp) Mnemonic() string {
	return "PLP"
}

func (i *plp) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	poppedSR := proc.registers.SP.Pop()
	proc.registers.SR.SetTo(SRFlagN, poppedSR&0x80 > 0)
	proc.registers.SR.SetTo(SRFlagV, poppedSR&0x40 > 0)
	proc.registers.SR.SetTo(SRFlagD, poppedSR&0x08 > 0)
	proc.registers.SR.SetTo(SRFlagI, poppedSR&0x04 > 0)
	proc.registers.SR.SetTo(SRFlagZ, poppedSR&0x02 > 0)
	proc.registers.SR.SetTo(SRFlagC, poppedSR&0x01 > 0)

	// N.B. the values of the Break and unused flags are not modified from the popped value
	return nil
}
