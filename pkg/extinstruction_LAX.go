package sim6502

type lax struct {
	standardCrossPage
}

func (i *lax) Mnemonic() string {
	return "LAX"
}

func (i *lax) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.A = data
	proc.registers.X = data
	return nil
}
