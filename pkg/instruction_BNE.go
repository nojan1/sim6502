package sim6502

import "errors"

type bne struct {
	variableInstructionBranch
}

func (i *bne) Mnemonic() string {
	return "BNE"
}

func (i *bne) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if !proc.registers.SR.IsSet(SRFlagZ) {
		// Z is not set, branch
		if proc.processorErrorOnSelfJump && data == 0xFE { // i.e. -2
			return errors.New("self jump on BNE")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}
