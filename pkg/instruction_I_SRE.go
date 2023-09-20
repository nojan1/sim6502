package sim6502

type sre struct{}

func (i *sre) Mnemonic() string {
	return "SRE"
}

func (i *sre) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	carryOut := data&0x01 == 0x01
	intermediate := data >> 1
	proc.memory.Write(data16, intermediate)
	proc.registers.SR.SetTo(SRFlagC, carryOut)

	proc.registers.A = proc.registers.A ^ intermediate
	proc.registers.SR.setNZ(proc.registers.A)

	return nil
}
