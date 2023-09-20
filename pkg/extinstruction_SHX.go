package sim6502

type shx struct{}

func (i *shx) Mnemonic() string {
	return "SHX"
}

func (i *shx) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	result := proc.registers.X&uint8(data16>>8) + 1
	proc.memory.Write(data16, result)
	return nil
}
