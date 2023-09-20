# github.com/cjbearman/sim6502

A simulator (emulator) for 6502 and 65C02 processors, for golang

* Full support for interrupts (NMI, IRQ, RST)
* Supports memory mapped I/O via callbacks
* Supports breakpoints via callbacks
* Optional support for illegal instructions (best effort, not recommended)

# Usage

## Step 1. Create a memory implementation

First thing you'll want to do is to create a memory implementation.

### Raw memory implementation (no memory mapping support, fastest)

```golang
mem := &sim6502.RawMemory{}
```

### Memory with memory mapping support

```golang
mem := &sim6502.MappableMemory{}
```

### More advanced memory

Feel free to create your own implementation of sim6502.Memory interface for more complex use cases.

## Step 2. Create a processor

### 6502 (NMOS)
```golang
proc := sim6502.NewProcessor(mem)
```
or,
### 65C02 (CMOS)
```golang
proc := sim6502.NewProcessor(mem).SetModel65C02()
```

## Step 3. Load some code
Whilst you could use the methods of the Memory implementation to load code in any custom way, support is
provided for intel .hex files and raw binary files:

### Load Intel hex file
```golang
f, err := os.Open("/tmp/dat")
if err != nil {
    panic(err)
}
defer(f.Close())
proc.LoadHex(f)
```

Since LoadHex takes a reader, you can use non-file sources also. An invalid file will cause a panic.

### Load raw binary file

Same method as LoadHex, just use Load method instead

## Step 4. Define your clock
By default the simulator will run with a clock rate of 0, which is unrestricted. No attempt at timing, run as fast as you can.

On my M1 mac this gives me an effective clock rate of >100Mhz, which is hardly realistic, but very good for fast testing.

More reasonably, use the SetClock() method of the processor to set the number of CYCLES (not instructions, one instruction has multiple cycles) per second.

For example, for a 1Mhz clock:
```golang
proc.SetClock(1000000)
```

The simulator will meter it's operation to achieve the desired clock rate, or as close as it can get.

After stopping the processor, you can call GetLastRunPerformance() to get information about how it actually performed.


## Step 5. Run your code
```golang
addr := uint16(0x0400)  // The start address
err := proc.Run(addr)
```

Your code will run until one of the following happens:
* An invalid opcode is reached (error returned)
* Stop() method is called on processor from another goroutine (nil returned)
* STP operation enconutered (65C02 only)


## Breakpoints
For both debugging, and for - in certain circumstances - getting data into and out of the processor, you can use breakpoints.

A breakpoint is defined by implementing a struct that honors the sim6502.Breakpoint interface:
```golang
type MyBreakpointHandler struct{}
func (b *MyBreakpointHandler) HandleBreak(proc *sim6502.Processor) {
    // Code here is called when your breakpoint is hit
}
```

Once you have your handler, add it to the processor and indicate the address at which it is set.
```golang
proc.SetBreakpoint(uint64(addr), &MyBreakpointHandler{})
```

When the program counter reaches the address and BEFORE the instruction at that address is executed, your breakpoint will be called. You can use the processor passed to the breakpoint to:
* Access / modify registers and flags
* Access / modify memory
* Stop the processor
* Enable or disable tracing
* ...

You can set multiple breakpoint handlers at the same address. They will be executed in the order set.

You can clear all breakpoints with the proc.ClearBreakpoints() method

## Interrupts
There are three types of hardware interrupt (plus of cause the BRK opcode)

### IRQ
IRQs are driven by a (simulated) I/O line. When this line is asserted (technically pulled low), an IRQ will be entered unless the processor I flag is set. The IRQ routine who's address is programmed at the IRQ system vector will be entered and typically ends with an RTI instruction.

Interrupts will keep occurring until the I/O line goes low (depending of course on the I flag)

Pull the simulated IRQ line low with proc.IRQ(true), reset it to high with proc.IRQ(false)

Optionally you can proc.SetOption(sim6502.AutoResetIRQ). With this option set, when the IRQ routine is entered, the line will be reset to high (i.e. un-asserted). This is not standard behavior, but is useful for when you just want to trigger an IRQ and don't want to wait for something else before disabling the line again.

Set the IRQ vector with
```golang
sim6502.SetVector(mem, sim6502.VectorIRQ, uint16(addr))
```

### NMI
NMIs are the same as IRQs, except they are only triggered on a falling edge of the interrupt line and they cannot be masked by the processor I flag.

The proc.NMI(true) call will set the line low (asserted). If it was not already low, this will trigger a single NMI.

As with IRQ, you can use proc.SetOption(sim6502.AutoResetNMI) to have the line auto-reset after the NMI is entered.

Set the NMI vector with
```golang
sim6502.SetVector(mem, sim6502.VectorNMI, uint16(addr))
```

### RST
The reset line is a little different, when pulled low (proc.Reset(true)), the processor stops. When set back high (proc.Reset(false)), the processor restarts by jumping to the RESET vector.

Set the reset vector with
```golang
sim6502.SetVector(mem, sim6502.VectorReset, uint16(addr))
```

## Odds and ends

### Fluent composition
The Processor implementation returns itself from most methods, allowing for fluent (builder style) composition:
```golang
err := sim6502.NewProcessor(mem).
    SetModel65C02().
    SetOption(sim6502.Trace).
    SetBreakpoint(0x1234, myHandler).
    Run()
```

### Debugging
Various debugging options are availbale that will output debugging information to a writer of your choice.

```golang
proc.SetDebugWriter(os.Stderr)                 // STDERR is the default
proc.SetOption(sim6502.Trace, true)            // Outputs one trace line per operation
proc.SetOption(sim6502.TraceStack, true)       // Outputs the entire stack prior to every operation (use sparingly)
proc.SetOption(sim6502.TraceInterrupts, true)  // Outputs information on every attempted IRQ/NMI/Reset
```

There are other options availble, check the docs.

You'll also find debug options in the memory implementations.

### Illegal instructions
By default, undefined 6502 instructions are considered illegal and will cause the processor to stop with error.
You can add support for these illegal operations by calling proc.LoadExtendedInstructions()

These operations are unstable, and hard to validate. So this is really not recommended.


## Trace
The sim6502.Trace option is very useful.

Bear in mind, you don't have to have this set at all times, you can turn it on and off via breakpoints that you set.

Each trace line output looks like this:

```
$36a7: $61 $46     : ADC  ($46,X)     [A=$30 X=$0e Y=$ff SP=$fc SH=$72 SR=$50 CY=6 N0 V1 B1 D0 I0 Z0 C0 MA=$0203 MC=$01]
```
This consists of 
* The PC location for the instruction
* The raw hex for the instruction and it's operands
* The operation decoded with appropriate addressing mode
* Registers and other information in the section surrounded by square brackets

The values in the last section are:
* A: The accumulator
* X: The X register
* Y: The Y register
* SP: The stack pointer
* SH: The stack head (most recent byte in the stack)
* SR: The status register
* N, V, B, D, I, Z, C: Status of all status register flags
* MA: The memory address referenced by the instruction, if any
* MC: The memory contents referenced by the instruction, if any

Note that all values shown are PRIOR to the operation executing. If you want to see the values AFTER the operation, look at the next line.

## Testing
There are a number of unit tests (go test -v ./... as per usual).

I have also implemented the tests provided by the excellent https://github.com/Klaus2m5/6502_65C02_functional_tests/tree/master repository, all which run to successful completion. The implementation for these tests can be found at:
https://github.com/cjbearman/sim6502test


## Known issues
* Some of the timing of the 85C02 operations may be a little off at the moment, this is being worked on.
* Illegal instructions are likely pretty spotty.

