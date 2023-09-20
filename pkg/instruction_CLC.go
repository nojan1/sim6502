package sim6502

type clc struct{}

func (i *clc) Mnemonic() string {
	return "CLC"
}

func (i *clc) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.Clear(SRFlagC)
	return nil
}
