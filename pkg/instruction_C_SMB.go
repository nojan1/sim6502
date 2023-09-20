package sim6502

type smb0 struct{ variableInstructionBranch }
type smb1 struct{ variableInstructionBranch }
type smb2 struct{ variableInstructionBranch }
type smb3 struct{ variableInstructionBranch }
type smb4 struct{ variableInstructionBranch }
type smb5 struct{ variableInstructionBranch }
type smb6 struct{ variableInstructionBranch }
type smb7 struct{ variableInstructionBranch }

func (i *smb0) Mnemonic() string { return "SMB0" }
func (i *smb1) Mnemonic() string { return "SMB1" }
func (i *smb2) Mnemonic() string { return "SMB2" }
func (i *smb3) Mnemonic() string { return "SMB3" }
func (i *smb4) Mnemonic() string { return "SMB4" }
func (i *smb5) Mnemonic() string { return "SMB5" }
func (i *smb6) Mnemonic() string { return "SMB6" }
func (i *smb7) Mnemonic() string { return "SMB7" }

func (i *smb0) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data|0x01)
	return nil
}

func (i *smb1) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data|0x02)
	return nil
}

func (i *smb2) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data|0x04)
	return nil
}

func (i *smb3) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data|0x08)
	return nil
}

func (i *smb4) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data|0x10)
	return nil
}

func (i *smb5) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data|0x20)
	return nil
}

func (i *smb6) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data|0x40)
	return nil
}

func (i *smb7) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.memory.Write(data16, data|0x80)
	return nil
}
