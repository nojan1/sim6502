package sim6502

import (
	"errors"
	"time"
)

type jam struct{}

func (i *jam) Mnemonic() string {
	return "JAM"
}

func (i *jam) Exec(proc *Processor, mode AddressingMode, data uint8, data16 uint16) error {

	if proc.errorOnJAM {
		return errors.New("JAM instruction encountered")
	}

	for {
		time.Sleep(100 * time.Microsecond)
		if proc.stop {
			return errors.New("stopped in JAM")
		}
	}
}
