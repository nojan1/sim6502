package sim6502

type cld struct{}

func (i *cld) Mnemonic() string {
	return "CLD"
}

func (i *cld) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.Clear(SRFlagD)
	return nil
}
