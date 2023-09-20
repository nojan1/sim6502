package sim6502

import "errors"

type bbs0 struct{ variableInstructionBranch }
type bbs1 struct{ variableInstructionBranch }
type bbs2 struct{ variableInstructionBranch }
type bbs3 struct{ variableInstructionBranch }
type bbs4 struct{ variableInstructionBranch }
type bbs5 struct{ variableInstructionBranch }
type bbs6 struct{ variableInstructionBranch }
type bbs7 struct{ variableInstructionBranch }

func (i *bbs0) Mnemonic() string { return "BBS0" }
func (i *bbs1) Mnemonic() string { return "BBS1" }
func (i *bbs2) Mnemonic() string { return "BBS2" }
func (i *bbs3) Mnemonic() string { return "BBS3" }
func (i *bbs4) Mnemonic() string { return "BBS4" }
func (i *bbs5) Mnemonic() string { return "BBS5" }
func (i *bbs6) Mnemonic() string { return "BBS6" }
func (i *bbs7) Mnemonic() string { return "BBS7" }

func (i *bbs0) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x01 > 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBS0")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbs1) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x02 > 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBS1")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbs2) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x04 > 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBS2")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbs3) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x08 > 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBS3")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbs4) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x10 > 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBS4")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbs5) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x20 > 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBS5")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbs6) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x40 > 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBS6")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbs7) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x80 > 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBS8")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}
