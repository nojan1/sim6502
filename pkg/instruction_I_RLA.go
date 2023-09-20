package sim6502

type rla struct{}

func (i *rla) Mnemonic() string {
	return "RLA"
}

func (i *rla) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	carryOut := data&0x80 == 0x80
	carryIn := uint8(0x00)
	if proc.registers.SR.IsSet(SRFlagC) {
		carryIn = 0x01
	}

	result := (data << 1) | carryIn
	proc.registers.SR.SetTo(SRFlagC, carryOut)

	proc.memory.Write(data16, result)
	proc.registers.A = result & proc.registers.A

	proc.registers.SR.setNZ(proc.registers.A)

	return nil
}
