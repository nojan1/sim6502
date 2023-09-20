package sim6502

type wai struct{}

func (i *wai) Mnemonic() string {
	return "WAI"
}

func (i *wai) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	// Bit of a hack, but since a signal on the IRQ line will break high, and given that
	// our main processor code only breaks this wait when the IRQ line is initially set
	// then we will only enter a wait if IRQ is not set
	if !proc.irq {
		proc.wait.wait()
	}
	return nil
}
