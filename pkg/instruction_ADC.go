package sim6502

type adc struct {
	standardCrossPage
}

func (i *adc) Mnemonic() string {
	return "ADC"
}

func (i *adc) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	result, c, v, n, z := add(
		proc.registers.SR.IsSet(SRFlagC),
		proc.registers.SR.IsSet(SRFlagD),
		proc.registers.A,
		data)

	proc.registers.A = result
	proc.registers.SR.SetTo(SRFlagC, c)
	proc.registers.SR.SetTo(SRFlagN, n)
	proc.registers.SR.SetTo(SRFlagZ, z)

	if !proc.registers.SR.IsSet(SRFlagD) {
		proc.registers.SR.SetTo(SRFlagV, v)
	}

	return nil
}

// add will add val1 and val2 with carry, using bcd mode if bcd is set
// returns the appropriate result, plus suggested values for the c, v, n, z flags
// the v flag should be ignored in BCD mode
// abstracted out so it can be used in some of the extended instructions in novel ways
func add(carryIn bool, bcd bool, val1, val2 uint8) (result uint8, c, v, n, z bool) {
	carry := uint16(0)
	if carryIn {
		carry = 1
	}

	if !bcd {
		// Binary mode
		result16 := uint16(val1) + uint16(val2) + carry

		// Carry results if the operation of UNSIGNED arithmatic results in > 0xff
		c = result16 > 0xff

		result = uint8(result16 & 0xff)
		v = overflows(val1, val2, uint8(carry))
	} else {
		v1 := fixupBCD(val1)
		v2 := fixupBCD(val2)
		b1 := (v1>>4)*10 + (v1 & 0x0f)
		b2 := (v2>>4)*10 + (v2 & 0x0f)
		br := b1 + b2 + carry
		c = br > 99
		br = br % 100
		result = uint8((br/10)<<4 + br%10)
	}
	n = result&0x80 == 0x80
	z = result == 0

	return
}

func fixupBCD(bcd uint8) uint16 {
	carry := uint8(0)
	low := bcd & 0x0f
	if low > 9 {
		low += 6
		carry = 1
	}
	low = low & 0x0f

	high := (bcd >> 4) + carry
	if high > 9 {
		high += 6
	}
	high = high & 0x0f

	return uint16(high<<4 | low)
}
