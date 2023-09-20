package devices

// Implements a memory mapped console output

import (
	"io"

	sim6502 "github.com/cjbearman/sim6502/pkg"
)

type Console struct {
	output   io.Writer
	location uint16
}

func (c *Console) AddressRange() []sim6502.MappedMemoryAddressRange {
	return []sim6502.MappedMemoryAddressRange{
		{Start: c.location, End: c.location},
	}
}

func (c *Console) Write(addr uint16, val uint8) {
	c.output.Write([]byte{byte(val)})
}

// NewConsole returns a new memory mapped console device
// that will output to the specified writer, any bytes written
// to the specified memory location
func NewConsole(location uint16, output io.Writer) *Console {
	return &Console{output: output, location: location}
}
