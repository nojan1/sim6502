package sim6502

type TXS struct{}

func (i *TXS) Mnemonic() string {
	return "TXS"
}

func (i *TXS) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SP.SetStackPointer(proc.registers.X)
	return nil
}
