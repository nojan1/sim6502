package sim6502

type clv struct{}

func (i *clv) Mnemonic() string {
	return "CLV"
}

func (i *clv) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.Clear(SRFlagV)
	return nil
}
