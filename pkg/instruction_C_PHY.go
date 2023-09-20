package sim6502

type phy struct {
	variableInstructionBranch
}

func (i *phy) Mnemonic() string {
	return "PHY"
}

func (i *phy) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SP.Push(proc.registers.Y)
	return nil
}
