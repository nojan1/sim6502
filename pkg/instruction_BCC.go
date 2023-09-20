package sim6502

import "errors"

type bcc struct {
	variableInstructionBranch
}

func (i *bcc) Mnemonic() string {
	return "BCC"
}

func (i *bcc) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if !proc.registers.SR.IsSet(SRFlagC) {
		if proc.processorErrorOnSelfJump && data == 0xFE { // i.e. -2
			return errors.New("self jump on BCC")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}
