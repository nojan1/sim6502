package sim6502

// BreakpointHandler defines the interface to be implemented by
// the user when implementing a breakpoint
type BreakpointHandler interface {

	// HandleBreak will be called when the PC reaches the listed brekapoint address
	// but before that instruction is executed
	// The current processor state is passed to the handler for inspection and, if desired,
	// modification
	HandleBreak(proc *Processor) error
}

// BreakpointHandlerEnableTrace is a provided breakpoint handler that will
// enable trace when the breakpoint is reached
type BreakpointHandlerEnableTrace struct{}

func (b *BreakpointHandlerEnableTrace) HandleBreak(proc *Processor) {
	proc.SetOption(Trace, true)
}

// BreakpointHandlerDisableTrace is a provided breakpoint handler that will
// disable trace when the breakpoint is reached
type BreakpointHandlerDisableTrace struct{}

func (b *BreakpointHandlerDisableTrace) HandleBreak(proc *Processor) {
	proc.SetOption(Trace, false)
}
