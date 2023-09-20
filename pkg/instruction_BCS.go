package sim6502

type bcs struct {
	variableInstructionBranch
}

func (i *bcs) Mnemonic() string {
	return "BCS"
}

func (i *bcs) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if proc.registers.SR.IsSet(SRFlagC) {
		proc.registers.PC.Set(data16)
	}
	return nil
}
