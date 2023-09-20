package sim6502

type shy struct{}

func (i *shy) Mnemonic() string {
	return "SHY"
}

func (i *shy) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	result := proc.registers.Y&uint8(data16>>8) + 1
	proc.memory.Write(data16, result)
	return nil
}
