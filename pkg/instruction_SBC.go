package sim6502

type sbc struct {
	standardCrossPage
}

func (i *sbc) Mnemonic() string {
	return "SBC"
}

func (i *sbc) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	result, c, v, n, z := sub(
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

func sub(carryIn bool, bcd bool, val1, val2 uint8) (result uint8, c, v, n, z bool) {

	if !bcd {
		// Binary mode
		// SBC is just ADC with input reversed
		result, c, v, n, z = add(carryIn, bcd, val1, val2^0xff)
	} else {

		carry := int16(0)
		if !carryIn {
			carry = 1
		}

		v1 := fixupBCD(val1)
		v2 := fixupBCD(val2)
		b1 := (v1>>4)*10 + (v1 & 0x0f)
		b2 := (v2>>4)*10 + (v2 & 0x0f)
		br := int16(b1) - int16(b2) - carry
		if br < 0 {
			br += 100
		} else {
			c = true
		}
		result = uint8((br/10)<<4 + br%10)
		n = result&0x80 == 0x80
		z = result == 0
	}
	return
}
