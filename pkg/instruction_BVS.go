package sim6502

import "errors"

type bvs struct {
	variableInstructionBranch
}

func (i *bvs) Mnemonic() string {
	return "BVS"
}

func (i *bvs) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if proc.registers.SR.IsSet(SRFlagV) {
		if proc.processorErrorOnSelfJump && data == 0xFE { // i.e. -2
			return errors.New("self jump on BVS")
		}

		proc.registers.PC.Set(data16)
	}
	return nil
}
