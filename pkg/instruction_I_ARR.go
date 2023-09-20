package sim6502

type arr struct{}

func (i *arr) Mnemonic() string {
	return "ARR"
}

func (i *arr) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	// V flag is set according to (A AND oper) + oper
	_, _, v, _, _ := add(
		proc.registers.SR.IsSet(SRFlagC),
		proc.registers.SR.IsSet(SRFlagD),
		proc.registers.A&data,
		data)

	proc.registers.SR.SetTo(SRFlagV, v)

	// The carry is not set, but bit 7 (sign) is exchanged with the carry
	signBit := proc.registers.SR.IsSet(SRFlagN)
	carryBit := proc.registers.SR.IsSet(SRFlagC)
	proc.registers.SR.SetTo(SRFlagN, carryBit)
	proc.registers.SR.SetTo(SRFlagC, signBit)

	//A AND oper, C -> [76543210] -> C (i.e. ROR)
	var carryIn uint8 = 0
	if proc.registers.SR.IsSet(SRFlagC) {
		carryIn = 0x80
	}
	primaryResult := proc.registers.A & data
	primaryResult = primaryResult | carryIn
	proc.registers.SR.SetTo(SRFlagC, primaryResult&0x10 == 0x10)
	proc.registers.A = primaryResult >> 1

	return nil
}
