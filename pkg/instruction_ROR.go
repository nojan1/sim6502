package sim6502

type ror struct{}

func (i *ror) Mnemonic() string {
	return "ROR"
}

func (i *ror) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	carry := uint8(0)
	if proc.registers.SR.IsSet(SRFlagC) {
		carry = uint8(0x80)
	}

	proc.registers.SR.SetTo(SRFlagC, data&0x01 > 0)
	result := (data >> 1) | carry

	if mode == A {
		proc.registers.A = result
	} else {
		proc.memory.Write(data16, result)
	}
	proc.registers.SR.setNZ(result)
	return nil
}
