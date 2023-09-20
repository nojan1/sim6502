package sim6502

import "errors"

type bpl struct {
	variableInstructionBranch
}

func (i *bpl) Mnemonic() string {
	return "BPL"
}

func (i *bpl) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if !proc.registers.SR.IsSet(SRFlagN) {
		if proc.processorErrorOnSelfJump && data == 0xFE { // i.e. -2
			return errors.New("self jump on BPL")
		}

		proc.registers.PC.Set(data16)
	}
	return nil
}
