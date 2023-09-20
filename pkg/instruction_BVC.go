package sim6502

import "errors"

type bvc struct {
	variableInstructionBranch
}

func (i *bvc) Mnemonic() string {
	return "BVC"
}

func (i *bvc) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if !proc.registers.SR.IsSet(SRFlagV) {
		if proc.processorErrorOnSelfJump && data == 0xFE { // i.e. -2
			return errors.New("self jump on BVC")
		}

		proc.registers.PC.Set(data16)
	}
	return nil
}
