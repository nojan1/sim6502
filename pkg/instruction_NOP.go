package sim6502

type nop struct{}

func (i *nop) Mnemonic() string {
	return "NOP"
}

func (i *nop) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	// Special behaviors for some of the 65C02 NOPs
	if proc.model == ProcessorModel65C02 {
		opcode := proc.memory.Read(proc.registers.PC.Current() - 1, true)
		switch opcode {
		case 0x02, 0x22, 0x42, 0x62, 0x82, 0xc2, 0xe2, 0x44, 0x54, 0xd4, 0xf4:
			// These NOP codes take an extra byte
			proc.registers.PC.Next()
		case 0x5c, 0xdc, 0xfc:
			// These NOP codes take two extra bytes
			proc.registers.PC.Next()
			proc.registers.PC.Next()
		}
	}

	return nil
}
