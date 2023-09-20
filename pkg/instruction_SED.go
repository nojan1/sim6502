package sim6502

type SED struct{}

func (i *SED) Mnemonic() string {
	return "SED"
}

func (i *SED) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.Set(SRFlagD)
	return nil
}
