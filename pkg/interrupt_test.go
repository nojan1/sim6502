package sim6502

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type breakpointCapture struct {
	triggered int
	stop      bool
}

func (b *breakpointCapture) HandleBreak(proc *Processor) error {
	b.triggered++
	if b.stop {
		proc.Stop()
	}
	return nil
}

func getInterruptProcessor() *Processor {
	// We will run the following
	// $1000 LDA #0
	// $1002 LDX #0
	// $1004 PHA				// A to stack
	// $1005 PLP				// SR from stack (clear all status bits)
	// $1006 CLI				// Can be replaced with a SEI to test interrupt marking
	// $1007 BNE FE				// Loop forever

	runloop := []uint8{0xa9, 0x00, 0xa2, 0x00, 0x48, 0x28, 0x58, 0xd0, 0xfe}

	// This will be the interrupt handler
	// $1100 INX				// Increment X
	// $1101 RTI				// Return from interrupt

	intloop := []uint8{0xe8, 0x40}

	// So basically X counts the number of interrupts that have happened

	return NewProcessor(&RawMemory{}).
		SetClock(200).                          // Let's run this slowly
		Load(bytes.NewBuffer(runloop), 0x1000). // Load run loop
		Load(bytes.NewBuffer(intloop), 0x1100)  // Load interrupt loop
}

func TestIRQ(t *testing.T) {

	proc := getInterruptProcessor()
	// Set IRQ vector
	proc.Memory().Write(uint16(VectorIRQ), 0x00)   // IRQ vector low byte
	proc.Memory().Write(uint16(VectorIRQ+1), 0x11) // IRQ vector high byte

	// For debugging
	// proc.SetOption(Trace, true)

	assert := assert.New(t)
	go proc.RunFrom(0x1000)
	time.Sleep(100 * time.Millisecond)

	// Send IRQ high for 250ms, that should trigger a few interrupts
	proc.IRQ(true)
	time.Sleep(250 * time.Millisecond)
	proc.IRQ(false)

	proc.Stop()

	// Should have got at least two interrupts
	assert.Greater(proc.registers.X, uint8(2))
}

func TestIRQWithMasking(t *testing.T) {

	proc := getInterruptProcessor()
	// Set IRQ vector
	proc.Memory().Write(uint16(VectorIRQ), 0x00)   // IRQ vector low byte
	proc.Memory().Write(uint16(VectorIRQ+1), 0x11) // IRQ vector high byte

	proc.Memory().Write(0x1006, 0x78) // Change CLI to SEI, thus masking interrupts

	// For debugging
	// proc.SetOption(Trace, true)

	assert := assert.New(t)
	go proc.RunFrom(0x1000)
	time.Sleep(100 * time.Millisecond)
	// Send IRQ high for 250ms, that should trigger a few interrupts
	proc.IRQ(true)
	time.Sleep(250 * time.Millisecond)
	proc.IRQ(false)

	proc.Stop()

	// IRQ should not have happened due to SEI
	assert.EqualValues(0, proc.registers.X)
}

func TestNMI(t *testing.T) {

	proc := getInterruptProcessor()
	// Set NMI vector
	proc.Memory().Write(uint16(VectorNMI), 0x00)   // NMI vector low byte
	proc.Memory().Write(uint16(VectorNMI+1), 0x11) // NMI vector high byte
	proc.SetOption(AutoResetNMI, true)
	// For debugging
	// proc.SetOption(Trace, true)

	assert := assert.New(t)
	go proc.RunFrom(0x1000)
	time.Sleep(100 * time.Millisecond)
	for i := 1; i <= 5; i++ {
		assert.Nil(proc.NMI(true))
		time.Sleep(100 * time.Millisecond)
	}

	proc.Stop()
	// Should have got at least two interrupts
	assert.Greater(proc.registers.X, uint8(2))
}

func TestNMIWithMasking(t *testing.T) {

	proc := getInterruptProcessor()
	// Set NMI vector
	proc.Memory().Write(uint16(VectorNMI), 0x00)   // NMI vector low byte
	proc.Memory().Write(uint16(VectorNMI+1), 0x11) // M<O vector high byte

	proc.Memory().Write(0x1006, 0x78) // Change CLI to SEI, thus masking IRQs, but importantly, not NMIs
	proc.SetOption(AutoResetNMI, true)

	// For debugging
	// proc.SetOption(Trace, true)

	assert := assert.New(t)
	go proc.RunFrom(0x1000)
	time.Sleep(100 * time.Millisecond)
	for i := 1; i <= 5; i++ {
		assert.Nil(proc.NMI(true))
		time.Sleep(100 * time.Millisecond)
	}

	proc.Stop()

	// SEI should not mask the NMIs, so we should still have had some interrupts
	assert.Greater(proc.registers.X, uint8(2))
}

// returns a processor set up to ensure we can break out of a WAI (65C02) with an
func getWAIInterruptProcessor() *Processor {
	// We will run the following
	// $1000 LDA #0
	// $1002 LDX #0
	// $1004 PHA				// A to stack
	// $1005 PLP				// SR from stack (clear all status bits)
	// $1006 CLI				// Can be replaced with a SEI to test interrupt marking
	// $1007 WAI				// Enter wait state
	// $1008 BNE $FE			// Loop forever

	runloop := []uint8{0xa9, 0x00, 0xa2, 0x00, 0x48, 0x28, 0x58, 0xcb, 0xd0, 0xfe}

	// This will be the interrupt handler
	// $1100 INX				// Increment X
	// $1101 RTI				// Return from interrupt

	intloop := []uint8{0xe8, 0x40}

	// So basically X counts the number of interrupts that have happened

	return NewProcessor(&RawMemory{}).
		SetClock(200). // Let's run this slowly
		SetModel65C02().
		Load(bytes.NewBuffer(runloop), 0x1000). // Load run loop
		Load(bytes.NewBuffer(intloop), 0x1100)  // Load interrupt loop
}

func TestNMIExitsWAI(t *testing.T) {

	require := require.New(t)
	assert := assert.New(t)
	proc := getWAIInterruptProcessor()
	defer proc.Stop()

	// Set NMI vector
	proc.Memory().Write(uint16(VectorNMI), 0x00)   // NMI vector low byte
	proc.Memory().Write(uint16(VectorNMI+1), 0x11) // NMI vector high byte
	proc.SetOption(AutoResetNMI, true)

	// Set a breakpoint on the BNE loop, reaching this means we successfully
	// exited the WAI
	breakpoint := &breakpointCapture{stop: true}
	proc.SetBreakpoint(0x1008, breakpoint)

	// For debugging
	// proc.SetOption(Trace, true)

	go proc.RunFrom(0x1000)

	// Let things get established
	time.Sleep(200 * time.Millisecond)

	// We should now be waiting
	require.True(proc.wait.isWaiting())

	// Trigger an interrupt, this should break the wait
	assert.Nil(proc.NMI(true))

	// Allow things to take their course
	time.Sleep(200 * time.Millisecond)

	assert.EqualValues(1, proc.registers.X) // Interrupt handler did it's thing
	assert.Equal(1, breakpoint.triggered)   // Reached the breakpoint
	assert.False(proc.wait.isWaiting())     // No longer waiting

}

func TestIRQExitsWAI(t *testing.T) {

	require := require.New(t)
	assert := assert.New(t)
	proc := getWAIInterruptProcessor()
	defer proc.Stop()

	// Set IRQ vector
	proc.Memory().Write(uint16(VectorIRQ), 0x00)   // IRQ vector low byte
	proc.Memory().Write(uint16(VectorIRQ+1), 0x11) // IRQ vector high byte

	// Set a breakpoint on the BNE loop, reaching this means we successfully
	// exited the WAI
	breakpoint := &breakpointCapture{stop: true}
	proc.SetBreakpoint(0x1008, breakpoint)

	// For debugging
	proc.SetOption(Trace, true)

	go proc.RunFrom(0x1000)

	// Let things get established
	time.Sleep(200 * time.Millisecond)

	// We should now be waiting
	require.True(proc.wait.isWaiting())

	// Send IRQ high for 250ms, that should trigger a few interrupts
	proc.IRQ(true)
	time.Sleep(250 * time.Millisecond)
	proc.IRQ(false)

	// Allow things to take their course
	time.Sleep(200 * time.Millisecond)

	assert.NotEqualValues(0, proc.registers.X)     // Interrupt handler did it's thing
	assert.NotEqualValues(0, breakpoint.triggered) // Reached the breakpoint
	assert.False(proc.wait.isWaiting())            // No longer waiting

}

func TestIRQExitsWAIMasked(t *testing.T) {

	require := require.New(t)
	assert := assert.New(t)
	proc := getWAIInterruptProcessor()
	defer proc.Stop()

	// Set IRQ vector
	proc.Memory().Write(uint16(VectorIRQ), 0x00)   // IRQ vector low byte
	proc.Memory().Write(uint16(VectorIRQ+1), 0x11) // IRQ vector high byte

	proc.Memory().Write(0x1006, 0x78) // Change CLI to SEI, thus masking IRQs

	// Set a breakpoint on the BNE loop, reaching this means we successfully
	// exited the WAI
	breakpoint := &breakpointCapture{stop: true}
	proc.SetBreakpoint(0x1008, breakpoint)

	// For debugging
	// proc.SetOption(Trace, true)

	go proc.RunFrom(0x1000)

	// Let things get established
	time.Sleep(200 * time.Millisecond)

	// We should now be waiting
	require.True(proc.wait.isWaiting())

	// Send IRQ high for 250ms, that should trigger a few interrupts
	proc.IRQ(true)
	time.Sleep(250 * time.Millisecond)
	proc.IRQ(false)

	// Allow things to take their course
	time.Sleep(200 * time.Millisecond)

	assert.EqualValues(0, proc.registers.X)        // There were no interrupts
	assert.NotEqualValues(0, breakpoint.triggered) // Reached the breakpoint
	assert.False(proc.wait.isWaiting())            // No longer waiting
}
