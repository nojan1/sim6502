package sim6502

type STY struct{}

func (i *STY) Mnemonic() string {
	return "STY"
}

func (i *STY) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, proc.registers.Y)
	return nil
}
