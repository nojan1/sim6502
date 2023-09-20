package sim6502

type rmb0 struct{ variableInstructionBranch }
type rmb1 struct{ variableInstructionBranch }
type rmb2 struct{ variableInstructionBranch }
type rmb3 struct{ variableInstructionBranch }
type rmb4 struct{ variableInstructionBranch }
type rmb5 struct{ variableInstructionBranch }
type rmb6 struct{ variableInstructionBranch }
type rmb7 struct{ variableInstructionBranch }

func (i *rmb0) Mnemonic() string { return "RMB0" }
func (i *rmb1) Mnemonic() string { return "RMB1" }
func (i *rmb2) Mnemonic() string { return "RMB2" }
func (i *rmb3) Mnemonic() string { return "RMB3" }
func (i *rmb4) Mnemonic() string { return "RMB4" }
func (i *rmb5) Mnemonic() string { return "RMB5" }
func (i *rmb6) Mnemonic() string { return "RMB6" }
func (i *rmb7) Mnemonic() string { return "RMB7" }

func (i *rmb0) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data&0xFE)
	return nil
}

func (i *rmb1) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data&0xFD)
	return nil
}

func (i *rmb2) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data&0xFB)
	return nil
}

func (i *rmb3) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data&0xF7)
	return nil
}

func (i *rmb4) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data&0xEF)
	return nil
}

func (i *rmb5) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data&0xDF)
	return nil
}

func (i *rmb6) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data&0xBF)
	return nil
}

func (i *rmb7) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data&0x7F)
	return nil
}
