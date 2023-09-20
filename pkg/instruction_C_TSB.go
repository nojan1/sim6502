package sim6502

type tsb struct{}

func (i *tsb) Mnemonic() string {
	return "TSB"
}

func (i *tsb) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SR.SetTo(SRFlagZ, data&proc.registers.A == 0)
	result := data | proc.registers.A
	proc.memory.Write(data16, result)
	return nil
}
