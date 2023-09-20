package sim6502

type bit struct{}

func (i *bit) Mnemonic() string {
	return "BIT"
}

func (i *bit) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if mode != IMMED { // N.B. flag behaviors differ for 65C02, hence this condition
		proc.registers.SR.SetTo(SRFlagN, data&0x80 > 0)
		proc.registers.SR.SetTo(SRFlagV, data&0x40 > 0)
	}
	res := proc.registers.A & data
	proc.registers.SR.SetTo(SRFlagZ, res == 0)
	return nil
}
