package sim6502

type cpx struct{}

func (i *cpx) Mnemonic() string {
	return "CPX"
}

func (i *cpx) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	compareToRegister(proc, proc.registers.X, data)
	return nil
}
