package sim6502

type sha struct{}

func (i *sha) Mnemonic() string {
	return "SHA"
}

func (i *sha) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	intermediate := proc.registers.A & proc.registers.X

	compareToRegister(proc, intermediate, data)

	// TODO: This line highly questionable
	proc.registers.X = intermediate - proc.registers.X
	return nil
}
