package sim6502

type lda struct {
	standardCrossPage
}

func (i *lda) Mnemonic() string {
	return "LDA"
}

func (i *lda) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	proc.registers.A = data
	proc.registers.SR.setNZ(data)
	return nil
}
