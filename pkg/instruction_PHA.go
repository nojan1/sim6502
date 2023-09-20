package sim6502

type pha struct{}

func (i *pha) Mnemonic() string {
	return "PHA"
}

func (i *pha) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SP.Push(proc.registers.A)
	return nil
}
