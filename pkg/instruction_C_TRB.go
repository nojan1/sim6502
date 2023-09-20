package sim6502

type trb struct{}

func (i *trb) Mnemonic() string {
	return "TRB"
}

func (i *trb) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	proc.registers.SR.SetTo(SRFlagZ, data&proc.registers.A == 0)
	result := data & (proc.registers.A ^ 0xFF)
	proc.memory.Write(data16, result)

	return nil
}
