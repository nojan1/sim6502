package sim6502

type usb struct{}

func (i *usb) Mnemonic() string {
	return "USB" // or USBC but sticking with three letter mnemonics
}

func (i *usb) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {
	// Same as SBC
	sbc := &sbc{}
	return sbc.Exec(proc, mode, data, data16)
}
