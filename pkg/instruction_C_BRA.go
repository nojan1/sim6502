package sim6502

import "errors"

type bra struct {
	variableInstructionBranch
}

func (i *bra) Mnemonic() string {
	return "BRA"
}

func (i *bra) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	if proc.processorErrorOnSelfJump && data == 0xFE { // i.e. -2
		return errors.New("self jump on BRA")
	}
	proc.registers.PC.Set(data16)
	return nil
}
