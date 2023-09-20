package sim6502

type isc struct{}

func (i *isc) Mnemonic() string {
	return "ISC"
}

func (i *isc) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	result := data + 1
	proc.memory.Write(data16, result)

	subRes, c, v, n, z := sub(
		proc.registers.SR.IsSet(SRFlagC),
		proc.registers.SR.IsSet(SRFlagD),
		proc.registers.A,
		result)

	proc.registers.A = subRes
	proc.registers.SR.SetTo(SRFlagC, c)
	proc.registers.SR.SetTo(SRFlagN, n)
	proc.registers.SR.SetTo(SRFlagZ, z)
	if !proc.registers.SR.IsSet(SRFlagD) {
		proc.registers.SR.SetTo(SRFlagV, v)
	}
	return nil
}
