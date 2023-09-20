package sim6502

type SEC struct{}

func (i *SEC) Mnemonic() string {
	return "SEC"
}

func (i *SEC) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.Set(SRFlagC)
	return nil
}
