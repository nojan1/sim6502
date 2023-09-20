package sim6502

// ProcessorModel allows for different processor models and their associated behaviors
// the default model is 6502
type ProcessorModel int

const (
	// ProcessorModel6502 represents the original 6502
	ProcessorModel6502 ProcessorModel = iota

	// ProcessorModel65C02 represents the CMOS 65C02
	// which exhibits some different behaviors with respect to the decimal flag
	// on BRK and interrupt
	ProcessorModel65C02
)
