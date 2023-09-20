package sim6502

type rra struct{}

func (i *rra) Mnemonic() string {
	return "RRA"
}

func (i *rra) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	carryOut := data&0x01 == 0x01
	carryIn := uint8(0x00)
	if proc.registers.SR.IsSet(SRFlagC) {
		carryIn = 0x80
	}

	result := (data >> 1) | carryIn
	proc.registers.SR.SetTo(SRFlagC, carryOut)

	proc.memory.Write(data16, result)

	result2, c, v, n, z := add(
		proc.registers.SR.IsSet(SRFlagC),
		proc.registers.SR.IsSet(SRFlagD),
		proc.registers.A,
		result)

	proc.registers.SR.SetTo(SRFlagZ, c)
	proc.registers.SR.SetTo(SRFlagZ, z)
	proc.registers.SR.SetTo(SRFlagN, n)
	if !proc.registers.SR.IsSet(SRFlagD) {
		proc.registers.SR.SetTo(SRFlagV, v)
	}

	proc.registers.A = result2

	return nil
}
