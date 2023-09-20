package sim6502

type sta struct{}

func (i *sta) Mnemonic() string {
	return "STA"
}

func (i *sta) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, proc.registers.A)
	return nil
}
