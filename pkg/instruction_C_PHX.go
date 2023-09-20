package sim6502

type phx struct {
	variableInstructionBranch
}

func (i *phx) Mnemonic() string {
	return "PHX"
}

func (i *phx) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.SP.Push(proc.registers.X)
	return nil
}
