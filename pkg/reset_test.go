package sim6502

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getResetProcessor() *Processor {
	// We will run the following
	// $1000 CLI
	// $1001 BNE FE				// Loop forever

	runloop := []uint8{0x58, 0xd0, 0xfe}

	// This will be the reset handler, we will break here
	// $1100 NOP

	intloop := []uint8{0x0e}

	return NewProcessor(&RawMemory{}).
		SetClock(1000).                         // Let's run this slowly
		Load(bytes.NewBuffer(runloop), 0x1000). // Load run loop
		Load(bytes.NewBuffer(intloop), 0x1100)  // Load reset handler
}

func TestReset(t *testing.T) {

	proc := getResetProcessor()

	// Set Reset vector
	proc.Memory().Write(uint16(VectorReset), 0x00)   // Reset vector low byte
	proc.Memory().Write(uint16(VectorReset+1), 0x11) // Reset vector high byte

	// For debugging
	// proc.SetOption(Trace, true)

	assert := assert.New(t)
	go proc.RunFrom(0x1000)
	defer proc.Stop()
	time.Sleep(100 * time.Millisecond)

	breakpoint := &breakpointCapture{stop: true}
	proc.SetBreakpoint(0x1100, breakpoint)
	// Set reset high
	proc.Reset(true)
	time.Sleep(200 * time.Millisecond)

	// At this point, nothing should have happened
	assert.EqualValues(0, breakpoint.triggered)
	assert.False(proc.registers.SR.IsSet(SRFlagI))

	// Set reset low
	proc.Reset(false)
	time.Sleep(200 * time.Millisecond)

	// Should have hit our breakpoint
	assert.EqualValues(1, breakpoint.triggered)
	assert.True(proc.registers.SR.IsSet(SRFlagI))
}
