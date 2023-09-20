package sim6502

type ane struct{}

// Because ANE is highly unstable, operation depends on a constant which may vary
// depending on chip series, temperature and so forth
// change the constant if you want
// ref: https://www.masswerk.at/6502/6502_instruction_set.html#ANE
var ExtendedInstructionANEConstant uint8 = 0x00

func (i *ane) Mnemonic() string {
	return "ANE"
}

func (i *ane) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.A = (proc.registers.A | ExtendedInstructionANEConstant) & proc.registers.X & data
	proc.registers.SR.setNZ(proc.registers.A)
	return nil
}
