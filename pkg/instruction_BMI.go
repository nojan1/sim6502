package sim6502

import "errors"

type bmi struct {
	variableInstructionBranch
}

func (i *bmi) Mnemonic() string {
	return "BMI"
}

func (i *bmi) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if proc.registers.SR.IsSet(SRFlagN) {
		if proc.processorErrorOnSelfJump && data == 0xFE { // i.e. -2
			return errors.New("self jump on BMI")
		}
		proc.registers.PC.Set(data16)
	}
	return nil
}
