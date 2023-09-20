package sim6502

import "fmt"

type SRFlag uint8

const (
	// SRFlagN is the bit address of the negative flag
	SRFlagN SRFlag = 7

	// SRFlagV is the bit address of the overflow flag
	SRFlagV SRFlag = 6

	// SRFlagU is the bit address of the unused flag
	SRFlagU SRFlag = 5

	// SRFlagB is the bit address of the break flag
	SRFlagB SRFlag = 4

	// SRFlagD is the bit addres of the decimal (BCD) flag
	SRFlagD SRFlag = 3

	// SRFlagI is the bit address of the interrupt flag
	SRFlagI SRFlag = 2

	// SRFlagZ is the bit address of the zero flag
	SRFlagZ SRFlag = 1

	// SRFlagC is the bit address of the carry flag
	SRFlagC SRFlag = 0
)

// StatusRegister defines a status register entity
type StatusRegister struct {

	// value contains the current status register value
	value uint8
}

// Set will set a specific flag in the status register
func (s *StatusRegister) Set(flag SRFlag) {
	s.value |= (1 << flag)
}

// Clear will clear a specific flag in the status register
func (s *StatusRegister) Clear(flag SRFlag) {
	mask := ^(uint8(1) << flag)
	s.value &= mask
}

// SetTo will set a specific flag in the status register to either
// on or off depending on the provided state (true=on)
func (s *StatusRegister) SetTo(flag SRFlag, state bool) {
	if state {
		s.Set(flag)
	} else {
		s.Clear(flag)
	}
}

// IsSet returns true if the specified flag is set, otherwise false
func (s *StatusRegister) IsSet(flag SRFlag) bool {
	return (s.value & (1 << flag)) > 0
}

// setNZ will set the N and Z flags based on the specified value
func (s *StatusRegister) setNZ(data uint8) {
	if data == 0 {
		s.value = s.value | 0x02
	} else {
		s.value = s.value & 0xFD
	}
	if data&0x80 > 0 {
		s.value = s.value | 0x80
	} else {
		s.value = s.value & 0x7F
	}
}

// String provides a string representation of the status register current state
func (s *StatusRegister) String() string {
	return fmt.Sprintf("N%d V%d B%d D%d I%d Z%d C%d",
		(s.value>>SRFlagN)&1,
		(s.value>>SRFlagV)&1,
		(s.value>>SRFlagB)&1,
		(s.value>>SRFlagD)&1,
		(s.value>>SRFlagI)&1,
		(s.value>>SRFlagZ)&1,
		(s.value>>SRFlagC)&1)
}

// Value returns the raw value of the status register
func (s *StatusRegister) Value() uint8 {
	return s.value
}
