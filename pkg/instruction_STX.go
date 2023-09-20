package sim6502

type STX struct{}

func (i *STX) Mnemonic() string {
	return "STX"
}

func (i *STX) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, proc.registers.X)
	return nil
}
