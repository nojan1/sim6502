package sim6502

type jsr struct{}

func (i *jsr) Mnemonic() string {
	return "JSR"
}

func (i *jsr) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	pc := proc.registers.PC.Current() - 1
	proc.registers.SP.Push(uint8(pc >> 8))   // PCH
	proc.registers.SP.Push(uint8(pc & 0xff)) // PCL

	proc.registers.PC.Set(data16)
	return nil
}
