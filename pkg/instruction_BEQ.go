package sim6502

import "errors"

type beq struct {
	variableInstructionBranch
}

func (i *beq) Mnemonic() string {
	return "BEQ"
}

func (i *beq) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if proc.registers.SR.IsSet(SRFlagZ) {
		if proc.processorErrorOnSelfJump && data == 0xFE { // i.e. -2
			return errors.New("self jump on BEQ")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}
