package sim6502

import "errors"

type bbr0 struct{ variableInstructionBranch }
type bbr1 struct{ variableInstructionBranch }
type bbr2 struct{ variableInstructionBranch }
type bbr3 struct{ variableInstructionBranch }
type bbr4 struct{ variableInstructionBranch }
type bbr5 struct{ variableInstructionBranch }
type bbr6 struct{ variableInstructionBranch }
type bbr7 struct{ variableInstructionBranch }

func (i *bbr0) Mnemonic() string { return "BBR0" }
func (i *bbr1) Mnemonic() string { return "BBR1" }
func (i *bbr2) Mnemonic() string { return "BBR2" }
func (i *bbr3) Mnemonic() string { return "BBR3" }
func (i *bbr4) Mnemonic() string { return "BBR4" }
func (i *bbr5) Mnemonic() string { return "BBR5" }
func (i *bbr6) Mnemonic() string { return "BBR6" }
func (i *bbr7) Mnemonic() string { return "BBR7" }

// Problem here is that data is the value compared, not the offset. Need to use data16 to find the self jumps

func (i *bbr0) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x01 == 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBR0")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbr1) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x02 == 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBR1")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbr2) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x04 == 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBR2")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbr3) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x08 == 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBR3")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbr4) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x10 == 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBR4")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbr5) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x20 == 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBR5")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbr6) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x40 == 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBR6")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}

func (i *bbr7) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if data&0x80 == 0 {
		if proc.processorErrorOnSelfJump && data16 == proc.registers.PC.Current()-3 {
			return errors.New("self jump on BBR7")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}
