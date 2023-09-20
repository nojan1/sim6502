package sim6502

type lsr struct{}

func (i *lsr) Mnemonic() string {
	return "LSR"
}

func (i *lsr) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	carry := data&0x01 > 0
	result := data >> 1
	proc.registers.SR.setNZ(result)
	proc.registers.SR.SetTo(SRFlagC, carry)

	if mode == A {
		proc.registers.A = result
	} else {
		proc.memory.Write(data16, result)
	}
	return nil
}
