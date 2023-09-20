package sim6502

type cpy struct{}

func (i *cpy) Mnemonic() string {
	return "CPY"
}

func (i *cpy) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	compareToRegister(proc, proc.registers.Y, data)
	return nil
}
