package sim6502

type dcp struct{}

func (i *dcp) Mnemonic() string {
	return "DCP"
}

func (i *dcp) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	// DEC operation (mem)
	result := data - 1
	proc.memory.Write(data16, data)

	// CMP operation
	compareToRegister(proc, proc.registers.A, result)
	return nil
}
