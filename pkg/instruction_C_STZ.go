package sim6502

type stz struct{}

func (i *stz) Mnemonic() string {
	return "STZ"
}

func (i *stz) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, 0x00)
	return nil
}
