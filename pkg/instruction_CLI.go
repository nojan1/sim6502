package sim6502

type cli struct{}

func (i *cli) Mnemonic() string {
	return "CLI"
}

func (i *cli) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.Clear(SRFlagI)
	return nil
}
