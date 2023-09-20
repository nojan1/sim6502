package sim6502

type php struct{}

func (i *php) Mnemonic() string {
	return "PHP"
}

func (i *php) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SP.Push(proc.registers.SR.value | 0x30) // 00110000 - Unused and Break flag get set (pos 5,4)
	return nil
}
