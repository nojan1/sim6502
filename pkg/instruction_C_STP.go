package sim6502

type stp struct{}

func (i *stp) Mnemonic() string {
	return "STP"
}

func (i *stp) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.Stop()
	return nil
}
