package sim6502

type SEI struct{}

func (i *SEI) Mnemonic() string {
	return "SEI"
}

func (i *SEI) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.Set(SRFlagI)
	return nil
}
