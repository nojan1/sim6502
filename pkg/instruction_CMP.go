package sim6502

type cmp struct {
	standardCrossPage
}

func (i *cmp) Mnemonic() string {
	return "CMP"
}

func (i *cmp) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	compareToRegister(proc, proc.registers.A, data)
	return nil
}

func compareToRegister(proc *Processor, reg, val uint8) {

	// Z Flag == 0 if reg != val
	// C Flag == UNSIGNED comparison result (0 reg < val, 1 reg >= val)
	// N is set to 7th bit of reg-val

	// fmt.Printf("CMP REG %02x VAL %02X\n", reg, val)

	if reg == val {
		// fmt.Printf("equal\n")
		proc.registers.SR.Clear(SRFlagN)
		proc.registers.SR.Set(SRFlagZ)
		proc.registers.SR.Set(SRFlagC)
		return
	}
	// fmt.Printf("ne\n")
	sub := reg - val
	proc.registers.SR.SetTo(SRFlagN, sub&0x80 > 0)
	proc.registers.SR.Clear(SRFlagZ)

	regGeVal := (reg >= val)
	// fmt.Printf("%d >= %d  %v\n", reg, val, regGeVal)
	// fmt.Printf("sub res %02x", sub)

	proc.registers.SR.SetTo(SRFlagC, regGeVal)

}
