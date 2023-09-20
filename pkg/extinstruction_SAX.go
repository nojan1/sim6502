package sim6502

type sax struct{}

func (i *sax) Mnemonic() string {
	return "SAX"
}

func (i *sax) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, proc.registers.A&proc.registers.X)
	return nil
}
