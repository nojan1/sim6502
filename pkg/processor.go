package sim6502

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/cjbearman/sim6502/pkg/hex"
)

// Processor represents a 6502 series processor
// Always create a new processor instance with NewProcessor() function
type Processor struct {
	model                    ProcessorModel                 // Our processor model#
	registers                Registers                      // Our registers
	memory                   Memory                         // Our memory IMPL
	instructions             []*instruction                 // The instruction set we are using
	instructionsExecuted     uint64                         // Count of instructions executed
	breakpoints              map[uint16][]BreakpointHandler // Breakpoints that are defined
	hasBreakpoints           bool                           // Set to true if breakpoints are defined
	processorErrorOnSelfJump bool                           // Set to true triggers optional behavior
	processorTrace           bool                           // Set to true, dumps each operation to debugWriter
	processorTraceStack      bool                           // Set to true, dumps stack on each operation
	traceInterrupts          bool                           // Set to true, dumps interrupt invocations to debugWriter
	errorOnJAM               bool                           // Set to true will cause the JAM error to stop the processor
	clock                    int64                          // Clock speed (cycles per second)
	nanosPerCycle            int64                          // Calculated as the nanoseconds needed for each cycle
	lastOpCycles             uint8                          // Tracks the cycles needed by the last operation
	lastInstruction          *instruction                   // The last instruction executed
	lastPC                   uint16                         // The last PC value
	lastNumCycles            int64                          // The number of cycles total after last op
	lastRanForNanos          int64                          // The amount of nanoseconds the last run was
	stop                     bool                           // when set true, signals that we need to stop
	nmi                      bool                           // when set true, NMI line is asserted
	autoResetNMI             bool                           // Allows for auto-reset of NMI line
	nmiLastState             bool                           // set to true after we see the NMI line assertion and trigger NMI
	irq                      bool                           // when set true, IRQ line is asserted
	autoResetIRQ             bool                           // Allows for auto-reset of IRQ line
	reset                    bool                           // when set true, RST line is asserted
	fixBrokenJMP             bool                           // Fixes incorrect JMP behavior on NMOS 6502
	wait                     *waiter                        // Used to manage WAI (65C02) instruction
	debugWriter              io.Writer                      // For outputting debug information
}

// ProcessorPerformance is used to report on the performance of the processor
// The data contained within is not valid if the processor is running, it is only
// valid once the processor has been stopped and reflects the stats from the last run
type ProcessorPerformance struct {
	// RanForNanoseconds indicates the number of nanoseconds for which the processor last ran
	RanForNanoseconds int64

	// RanForCycles indicates the number of cycles for which the processor last ran
	RanForCycles int64

	// EffectiveClock indicates the achieved clock rate of the last ran, in processor cycles per second
	EffectiveClock int64

	// InstructionsExecuted is the count of instructions executed
	InstructionsExecuted uint64
}

// NewProcessor returns a new, properly initialized 6502 processor
func NewProcessor(m Memory) *Processor {
	proc := &Processor{
		memory:      m,
		wait:        &waiter{},
		debugWriter: os.Stderr,
		breakpoints: make(map[uint16][]BreakpointHandler),
	}
	proc.Init()
	proc.copyInstructionSet()
	return proc
}

// Init will initialite the processor, clearing memory, registers and
// unsetting any breakpoints
func (p *Processor) Init() *Processor {
	// Stack pointer needs memory reference in order to store stack
	p.registers.SP.mem = p.memory
	p.registers.PC.processor = p
	p.registers.SR.value = 0
	p.registers.SR.Set(SRFlagI)
	p.registers.PC.Set(0)
	p.registers.SP.SetStackPointer(0xff)
	p.registers.A = 0
	p.registers.X = 0
	p.registers.Y = 0
	p.memory.Clear()
	p.nmi = false
	p.nmiLastState = false
	p.irq = false
	p.reset = false
	return p
}

// SetModel65C02 will load 65C02 additional opcodes
// and set other specific behaviors for that processor type
func (p *Processor) SetModel65C02() *Processor {
	if p.model != ProcessorModel65C02 {
		p.model = ProcessorModel65C02
		load85C02instructions(p.instructions)
	}
	return p
}

// LoadIllegalInstructions will provide support for undocumented/unofficial/illegal instruction set
func (p *Processor) LoadIllegalInstructions() *Processor {
	loadIllegalInstructions(p.instructions)
	return p
}

// Registers can be used to retrieve a copy of the current processor registers
func (p *Processor) Registers() *Registers {
	return &p.registers
}

// Memory can be used to retrieve a copy of the current memory implementation
func (p *Processor) Memory() Memory {
	return p.memory
}

// SetClock will set the clock speed in processor cycles per second
// I.E. for a 1MHZ clock set ticksPerSecond to 1,000,000
// A value of zero is unrestricted (all instructions run at maximum possible speed
// with no regard for timing)
func (p *Processor) SetClock(cyclesPerSecond int64) *Processor {
	p.clock = cyclesPerSecond
	p.nanosPerCycle = int64(float64(1000000000) / (float64(cyclesPerSecond)))
	return p
}

// SetBreakpoint is used to set a breakpoint consisting of a user provided handler
// The handler must return it's desired breakpoint address and a function
// that will be invoked when the PC reaches the breakpoint (before the instruction is executed)
// Call only when the processor is stopped
// Multiple breakpointers/handlers can be set at a given address and will be executed
// in the order set
func (p *Processor) SetBreakpoint(breakAt uint16, handler BreakpointHandler) *Processor {
	p.breakpoints[breakAt] = append(p.breakpoints[breakAt], handler)
	p.hasBreakpoints = true
	return p
}

// ClearBreakpoints will clear all previously defined breakpoints
// Call only when the processor is stopped
func (p *Processor) ClearBreakpoints() *Processor {
	p.breakpoints = make(map[uint16][]BreakpointHandler)
	return p
}

// ClearBreakpointsAt will remove all breakpoints from the specified address
func (p *Processor) ClearBreakpointsFrom(addr uint16) *Processor {
	p.breakpoints[addr] = []BreakpointHandler{}
	return p
}

// SetDebugWriter can be used to set a writer to which debug messages
// will be sent when debug options are activated
// This defaults to STDERR
func (p *Processor) SetDebugWriter(w io.Writer) *Processor {
	p.debugWriter = w
	return p
}

// SetOption can be used to set one of several processor options
// these are generally set to true to enable debugging facilities
func (p *Processor) SetOption(opt ProcessorOption, value bool) *Processor {
	switch opt {
	case Trace:
		p.processorTrace = value
	case TraceStack:
		p.processorTraceStack = value
	case ErrorOnSelfJump:
		p.processorErrorOnSelfJump = value
	case TraceInterrupts:
		p.traceInterrupts = value
	case ErrorOnJAM:
		p.errorOnJAM = value
	case AutoResetIRQ:
		p.autoResetIRQ = value
	case AutoResetNMI:
		p.autoResetNMI = value
	case Fix6052BrokenJMP:
		p.fixBrokenJMP = value
	default:
		panic(fmt.Sprintf("undefined processor option:%d", opt))
	}
	return p
}

// Stop will stop the processor if running
func (p *Processor) Stop() {
	p.stop = true
	// It's possible we're in a wait (WAI/65C02), that must be terminated
	p.wait.notify()
}

// Load is used to load raw binary content into the processor at the
// indicated start location
func (p *Processor) Load(r io.ByteReader, start uint16) *Processor {
	i := start
	if i > 0xffff {
		panic("load data exceeded memory limits")
	}
	for {
		b, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				return p
			} else {
				panic("failed to load data")
			}
		}
		p.memory.Write(i, b)
		i++
	}
}

// LoadHex is used to load an intel .hex file into the processor
func (p *Processor) LoadHex(r io.Reader) *Processor {

	hr, err := hex.NewHexReader(r)
	if err != nil {
		panic(fmt.Errorf("error reading hex file: %v", err))
	}

	for _, chunk := range hr.Chunks {
		addr := uint16(chunk.Addr)
		for _, b := range chunk.Data {
			if addr > 0xffff {
				panic("load hex exceeded memory limits")
			}
			p.memory.Write(addr, b)
			addr++
		}
	}
	return p
}

// GetLastRunPerformance will return a performance report that can be used
// to evaluate the last run of the processor
func (p *Processor) GetLastRunPerformance() ProcessorPerformance {
	return ProcessorPerformance{
		RanForNanoseconds:    p.lastRanForNanos,
		RanForCycles:         p.lastNumCycles,
		EffectiveClock:       int64((float64(p.lastNumCycles) / float64(p.lastRanForNanos)) * 1000000000),
		InstructionsExecuted: p.instructionsExecuted,
	}
}

// IRQ will assert (set low) the IRQ line when assert = true, or clear when assert = false
// IRQs will keep generating when the I flag is not set, whilst this line is asserted
// The line will remain as set until next modified, unless AutoResetIRQ option is used
// in which case once asserted it will self clear once the IRQ is initiated
func (p *Processor) IRQ(assert bool) {
	p.irq = assert
	// The processor could be in a WAI instruction, in which case we have to kick that loose
	if assert {
		p.wait.notify()
	}
}

// IsIRQSet returns true if the IRQ line is set (low)
func (p *Processor) IsIRQSet() bool {
	return p.irq
}

// Reset simulates the reset line being pulled low (state = true) or not (state = false)
// To reset the processor you have to pull it (state = true) and then set back to high
// after a few cycles, otherwise it will be missed
func (p *Processor) Reset(state bool) {
	p.reset = state
	// The processor could be in a WAI instruction, in which case we have to kick that loose
	if state {
		p.wait.notify()
	}
}

// IsResetSet returns true if the simulated reset line is pulled low
func (p *Processor) IsResetSet() bool {
	return p.reset
}

// NMI will assert (set low) the NMI line when assert = true, or clear when assert = false
// NMIs only trigger on the descending (assertin) edge, so although the line remains asserted until
// explicitly cleared, only a single NMI will result from a transition from cleared -> asserted
// If the AutoResetNMI option is set, however, the line will be cleared as soon as the NMI initiates
func (p *Processor) NMI(state bool) {
	p.nmi = state
	// The processor could be in a WAI instruction, in which case we have to kick that loose
	if !p.nmiLastState && p.nmi {
		p.wait.notify()
	}
}

// RunFrom will start the processor running at the specified address
// An error is returned if an unrecoverable error is encountered
// Typically this would be an undefined opcode, or an error returned
// from a breakpointer handler
func (p *Processor) RunFrom(addr uint16) (err error) {
	// Stop flag must be cleared
	p.stop = false

	// Any pending interrupts or wait state from a previous run must be cleared
	p.nmiLastState = false
	p.wait.notify()

	// Set PC to desired address
	p.registers.PC.Set(addr)

	// Grab thge start time, used for managing clock rate
	startTime := time.Now()
	var cumulativeTicks int64 = 0

	for {

		if p.reset {
			// Reset line is low, we just wait for it to go back high
			// unless we've been asked to stop, in which case that takes precedence
			for p.reset && !p.stop {
				if p.traceInterrupts {
					fmt.Println("Reset line is active")
				}
				time.Sleep(10 * time.Microsecond)
			}

			// Now it went high, we load the reset vector and
			// proceed
			p.registers.PC.Set(GetVector(p.memory, VectorReset))
			p.registers.SR.Set(SRFlagI)
			if p.traceInterrupts {
				fmt.Println("Reset line went low, continue from reset vector")
			}
		}

		// Step the processor
		err, stop := p.Step()

		if err != nil || stop {
			// If we get an error or a stop signal, we are done
			p.lastRanForNanos = time.Since(startTime).Nanoseconds()
			p.lastNumCycles = cumulativeTicks
			p.stop = false
			return err
		}

		// Increment cumulative ticks
		cumulativeTicks += int64(p.lastOpCycles)

		if p.clock > 0 {
			// Clock rate is enforced, so calculate if and how long we must sleep
			// to maintain that clock rate
			currentNanosSinceStart := time.Since(startTime).Nanoseconds()
			shouldBeAtNanos := cumulativeTicks * p.nanosPerCycle
			sleepNanos := shouldBeAtNanos - currentNanosSinceStart

			if sleepNanos > 0 {
				time.Sleep(time.Duration(sleepNanos) * time.Nanosecond)
			}
		}

		// If there are NMI or IRQs set, handle them now
		if p.nmi {
			// NMIs are falling edge only, we use nmiLastState to keep track of the edge
			if !p.nmiLastState {
				p.nmiLastState = true
				p.handleNMI()
			}
		} else {
			p.nmiLastState = false
		}

		if p.irq {
			p.handleIRQ()
		}
	}
}

// handleNMI is called when an NMI is triggered to transfer program control to the
// NMI routine via the NMI vector
func (p *Processor) handleNMI() {

	// If auto reset option is established, force the NMI line low
	if p.autoResetNMI {
		p.nmi = false
		p.nmiLastState = false
	}

	pc := p.registers.PC.Current()

	// Push PCH, PCL, SR to stack
	p.registers.SR.Clear(SRFlagB)             // Clears break
	p.registers.SP.Push(uint8(pc >> 8))       // PCH
	p.registers.SP.Push(uint8(pc & 0xff))     // PCL
	p.registers.SP.Push(p.registers.SR.value) // SR
	p.registers.SR.Set(SRFlagI)               // Sets interrupt disable

	if p.model == ProcessorModel65C02 {
		// 65C02 clears the decimal flag on interrupt
		p.registers.SR.Clear(SRFlagD)
	}

	targetAddress := GetVector(p.memory, VectorNMI)
	if p.traceInterrupts {
		fmt.Printf("NMI triggered. Return addr: 0x%04x, Control transferred via NMI vector to 0x%04x\n", pc, targetAddress)
	}
	p.registers.PC.Set(targetAddress)
}

// handleIRQ is called when an IRQ is triggered to transfer program control to the
// IRQ routine via the IRQ vector, unless the I flag is set
func (p *Processor) handleIRQ() {

	// If auto reset option is established, force the NMI line low
	if p.autoResetIRQ {
		p.irq = false
	}

	if p.registers.SR.IsSet(SRFlagI) {
		if p.traceInterrupts {
			fmt.Println("IRQ triggered, masked by I flag")
		}
		return
	}
	pc := p.registers.PC.Current()

	// Push PCH, PCL, SR to stack
	p.registers.SR.Clear(SRFlagB)             // Clears break
	p.registers.SP.Push(uint8(pc >> 8))       // PCH
	p.registers.SP.Push(uint8(pc & 0xff))     // PCL
	p.registers.SP.Push(p.registers.SR.value) // SR
	p.registers.SR.Set(SRFlagI)               // Sets interrupt disable

	if p.model == ProcessorModel65C02 {
		// 65C02 clears the decimal flag on interrupt
		p.registers.SR.Clear(SRFlagD)
	}

	targetAddress := GetVector(p.memory, VectorIRQ)
	if p.traceInterrupts {
		fmt.Printf("IRQ triggered, not masked by I flag. Return addr: 0x%04x, Control transferred via IRQ vector to 0x%04x\n", pc, targetAddress)
	}
	p.registers.PC.Set(targetAddress)
}

// DumpState can be used to output the current processor state to a writer
func (p *Processor) DumpState(w io.Writer) {
	fmt.Fprintf(w, "PC: $%04x\n", p.registers.PC.Current())
	fmt.Fprintf(w, "%s\n", p.registers.SR.String())
	fmt.Fprintf(w, "%s\n", p.registers.SP.String())
	fmt.Fprintf(w, "A: $%02x, X: $%02x, Y: $%02x\n", p.registers.A, p.registers.X, p.registers.Y)
	fmt.Fprintf(w, "Vectors: IRQ=$%04x, NMI=$%04x, RST=$%04x\n", GetVector(p.memory, VectorIRQ), GetVector(p.memory, VectorNMI), GetVector(p.memory, VectorReset))
	fmt.Fprintf(w, "Instructions Executed: %d\n", p.instructionsExecuted)
}

// setPCFromResetVector will set the program counter to the value of the reset vector
func (p *Processor) setPCFromResetVector() {
	p.registers.PC.Set(uint16(p.memory.Read(uint16(VectorReset))) | (uint16(p.memory.Read(uint16(VectorReset+1))) << 8))
}

// Step executes a single instruction, returning any error that may occur
// which is typically an undefined opcode or errors resulting from breakpointer handlers
func (p *Processor) Step() (err error, stop bool) {

	// Capture current program counter
	pc := p.registers.PC.Current()

	// If breakpoints are in use, check to see if we are at one of those breakpoints
	// and if we are, instantiate the handler
	if p.hasBreakpoints {
		handlers, found := p.breakpoints[pc]
		if found {
			for _, handler := range handlers {
				err := handler.HandleBreak(p)
				if err != nil {
					// The handler returned an error, so we just return that
					return nil, true
				}
			}
		}
	}

	if p.stop {
		return nil, true
	}

	// Get the opcode, advancing PC
	opcode := p.registers.PC.Next()

	// If this opcode isn't defined, we bail with an error
	instruction := p.instructions[opcode]
	if instruction == nil {
		return fmt.Errorf("illegal instruction: 0x%02x", opcode), true
	}

	// What we do now really depends on the addressing mode

	var data uint8         // For all operations with a single byte value (pretty much everything), this is that value or the value from the referenced mem cell
	var oper1, oper2 uint8 // Up to two operand bytes used by the instruction
	var addr uint16        // Memory address, for operations involving a memory location
	var xPageBoundary bool // Set to true if an ABS_X, ABS_Y or IND_Y mode increment causes a page boundary to be crossed
	cycles := instruction.BaseCycles

	switch instruction.AddressingMode {
	// ACCUMULATOR MODE
	case A:
		// input data is sourced from the accumulator
		data = p.registers.A

	// IMPLIED MODE
	case IMPL:
		// there is no input data
		data = 0

	// IMMEDIATE MODE
	case IMMED:
		// The data is in the next operand
		oper1 = p.registers.PC.Next()
		data = oper1

	// ABSOLUTE MODE
	case ABS:
		// The data is in the memory cell referenced by the next two operands as $LLHH
		oper1 = p.registers.PC.Next()
		oper2 = p.registers.PC.Next()
		addr = uint16(oper1) | (uint16(oper2) << 8)
		data = p.memory.Read(addr)

	// ABSOLUTE_X MODE
	case ABS_X:
		// The data is in the memory cell referenced by the next two operands (LL, HH) as ($HHLL plus X)
		oper1 = p.registers.PC.Next()
		oper2 = p.registers.PC.Next()
		baseAddr := (uint16(oper1) | (uint16(oper2) << 8))
		addr = baseAddr + uint16(p.registers.X)
		xPageBoundary = !samePage(baseAddr, addr)
		data = p.memory.Read(addr)

	// ABSOLUTE_Y_MODE
	case ABS_Y:
		// The data is in the memory cell referenced by the next two operands (LL HH) as ($HHLL plus Y)
		oper1 = p.registers.PC.Next()
		oper2 = p.registers.PC.Next()
		baseAddr := (uint16(oper1) | (uint16(oper2) << 8))
		addr = baseAddr + uint16(p.registers.Y)
		xPageBoundary = !samePage(baseAddr, addr)
		data = p.memory.Read(addr)

	// ZEROPAGE mode
	case ZPG:
		// The data is in the memory cell referenced by the next operand (LL) as $00LL
		oper1 = p.registers.PC.Next()
		addr = uint16(oper1)
		data = p.memory.Read(addr)

	// ZEROPAGE_X mode
	case ZPG_X:
		// The data is in the memory cell referenced by the next operand (LL) as $00(LL+X)
		oper1 = p.registers.PC.Next()
		addr = uint16(oper1 + p.registers.X)
		data = p.memory.Read(addr)

	// ZEROPAGE_Y mode
	case ZPG_Y:
		// The data is in the memory cell referenced by the next operand (LL) as $00(LL+Y)
		oper1 = p.registers.PC.Next()
		addr = uint16(oper1 + p.registers.Y)
		data = p.memory.Read(addr)

	// INDIRECT mode
	case IND:
		// The data is in the memory cell pointed to by the memory cell referenced by the next two operands (LL, HH)
		// as $HHLL
		oper1 = p.registers.PC.Next()
		oper2 = p.registers.PC.Next()
		ind := uint16(oper1) | (uint16(oper2) << 8)                             // The address of the target address
		addr = uint16(p.memory.Read(ind)) | (uint16(p.memory.Read(ind+1)) << 8) // The actual address
		data = p.memory.Read(addr)

	// INDIRECT_X mode
	case X_IND:
		// The data is in the memory cell pointed to by the memory cell referenced by the next two operands (LL, HH)
		// as $(HHLL+X)
		oper1 = p.registers.PC.Next()
		indaddr := uint16(oper1 + p.registers.X)
		addr = uint16(p.memory.Read(indaddr)) | (uint16(p.memory.Read(indaddr+1)) << 8)
		data = p.memory.Read(addr)

	// INDIRECT_Y mode
	case IND_Y:
		// The data is in the memory cell pointed to by the memory cell referenced by the next two operands (LL, HH)
		// as $(HHLL+Y)
		oper1 = p.registers.PC.Next()
		indaddr := uint16(oper1) // zpg address
		baseAddr := uint16(p.memory.Read(indaddr)) | (uint16(p.memory.Read(indaddr+1)) << 8)
		addr = baseAddr + uint16(p.registers.Y)
		xPageBoundary = !samePage(baseAddr, addr)
		data = p.memory.Read(addr)

	// RELATIVE mode
	case REL:
		// Operand is a relative offset from the PC of the instruction byte
		// The address is calculated from that
		oper1 = p.registers.PC.Next()
		data = oper1
		addr = p.registers.PC.Current()
		if oper1&0x80 > 0 {
			addr = addr - uint16(comp2(oper1))
		} else {
			addr = addr + uint16(oper1)
		}

	// 65C02 Zero page relative mode
	case ZPG_REL:
		oper1 = p.registers.PC.Next()
		data = p.memory.Read(uint16(oper1))
		offset := p.registers.PC.Next()
		oper2 = offset
		addr = p.registers.PC.Current()
		if oper1&0x80 > 0 {
			addr = addr - uint16(comp2(offset))
		} else {
			addr = addr + uint16(offset)
		}

	// 65C02 Indirect ABS_X
	case IND_ABS_X:
		oper1 = p.registers.PC.Next()
		oper2 = p.registers.PC.Next()
		indaddr := uint16(oper1) | (uint16(oper2) << 8)
		indaddr += uint16(p.registers.X)
		addr = uint16(p.memory.Read(indaddr)) | (uint16(p.memory.Read(indaddr+1)) << 8)
		data = p.memory.Read(addr)

	// 65C02 indirect zeropage
	case ZPG_IND:
		oper1 = p.registers.PC.Next()
		indaddr := uint16(oper1)
		addr = uint16(p.memory.Read(indaddr)) | (uint16(p.memory.Read(indaddr+1)) << 8)
		data = p.memory.Read(addr)

	default:
		// Should never happen
		return fmt.Errorf("undefined addressing mode: %d", instruction.AddressingMode), true
	}

	// Some instructions have variable cycles, if a cycle implements a variable cycle
	// method, add the result of that to cycles
	if variableInstruction, ok := instruction.Impl.(variableCycleInstruction); ok {
		cycles += variableInstruction.TweakCycles(instruction.AddressingMode, pc, addr, xPageBoundary)
	}

	// If processor trace is enabled, dump information to terminal
	if p.processorTrace {
		fmt.Fprintln(p.debugWriter, FormatInstruction(pc, instruction, oper1, oper2, addr, data, p.registers.SR.String(), cycles, p))
	}

	// ditto stack trace
	if p.processorTraceStack {
		fmt.Fprintln(p.debugWriter, p.registers.SP.String())
	}

	// Increment the execution counter
	p.instructionsExecuted++

	// Execute the instruction
	result := instruction.Impl.Exec(p, instruction.AddressingMode, data, addr)

	// Mostly for debugging
	p.lastOpCycles = cycles
	p.lastInstruction = instruction
	p.lastPC = pc
	return result, false
}

// copyInstructionSet takes a copy of the master instruction set for our use
// We copy it so we can extend it if needed
func (p *Processor) copyInstructionSet() {
	p.instructions = make([]*instruction, len(instructions))
	copy(p.instructions, instructions)
}
