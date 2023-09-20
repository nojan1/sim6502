package sim6502

type asl struct{}

func (i *asl) Mnemonic() string {
	return "ASL"
}

func (i *asl) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	carry := data&0x80 > 0
	result := data << 1
	proc.registers.SR.setNZ(result)
	proc.registers.SR.SetTo(SRFlagC, carry)
	if mode == A {
		proc.registers.A = result
	} else {
		proc.memory.Write(data16, result)
	}
	return nil
}
