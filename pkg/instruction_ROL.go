package sim6502

type rol struct{}

func (i *rol) Mnemonic() string {
	return "ROL"
}

func (i *rol) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	carry := uint8(0)
	if proc.registers.SR.IsSet(SRFlagC) {
		carry = uint8(1)
	}

	proc.registers.SR.SetTo(SRFlagC, data&0x80 > 0)
	result := (data << 1) | carry

	if mode == A {
		proc.registers.A = result
	} else {
		proc.memory.Write(data16, result)
	}
	proc.registers.SR.setNZ(result)
	return nil
}
