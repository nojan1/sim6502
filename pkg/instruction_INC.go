package sim6502

type inc struct{}

func (i *inc) Mnemonic() string {
	return "INC"
}

func (i *inc) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if mode == A {
		// 65C02 only
		proc.registers.A++
		proc.registers.SR.setNZ(proc.registers.A)
	} else {
		result := data + uint8(1)
		proc.registers.SR.setNZ(result)
		proc.memory.Write(data16, result)
	}
	return nil
}
