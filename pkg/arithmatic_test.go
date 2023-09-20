package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinaryADCWithoutCarry(t *testing.T) {

	require := require.New(t)
	assert := assert.New(t)
	for a := 0; a <= 0xff; a++ {

		for b := 0; b <= 0xff; b++ {

			// Add a (accum) to b (operand)

			proc := prepareInstructionImmediate(t, instructions, "ADC", uint8(b))
			proc.registers.A = uint8(a)
			proc.registers.SR.Clear(SRFlagC)
			proc.registers.SR.Clear(SRFlagD)
			require.Nil(proc.Step())

			expectedResult := uint8((a + b) & 0xff)
			result16 := a + b
			expectedZFlag := expectedResult == 0
			expectedNFlag := (expectedResult & 0x80) > 0
			expectedCFlag := result16 > 0xff

			expectedVFlag := false
			// If both numbers are positive...
			if a&0x80 == 0 && b&0x80 == 0 {
				// Overflow occurrs if the result > 0x7F (number not positive)
				if expectedResult > 0x7f {
					expectedVFlag = true
				}
			}
			// If both numbers are negative...
			if a&0x80 == 0x80 && b&0x80 == 0x80 {
				// Overflow occurss if the result is < 0x80 (number not negative)
				if expectedResult < 0x80 {
					expectedVFlag = true
				}
			}

			assert.Equal(expectedResult, proc.registers.A, fmt.Sprintf("%d+%d should = %d", a, b, expectedResult))
			assert.Equal(expectedZFlag, proc.registers.SR.IsSet(SRFlagZ), fmt.Sprintf("%d+%d=%d, expected Z flag incorrect", a, b, expectedResult))
			assert.Equal(expectedNFlag, proc.registers.SR.IsSet(SRFlagN), fmt.Sprintf("%d+%d=%d, expected N flag incorrect", a, b, expectedResult))
			assert.Equal(expectedCFlag, proc.registers.SR.IsSet(SRFlagC), fmt.Sprintf("%d+%d=%d, expected C flag incorrect", a, b, expectedResult))
			assert.Equal(expectedVFlag, proc.registers.SR.IsSet(SRFlagV), fmt.Sprintf("%d+%d=%d, expected V flag incorrect", a, b, expectedResult))

			proc.registers.SR.Clear(SRFlagD) // Decimal mode
		}
	}
}

func TestBinaryADCWithCarry(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	for a := 0; a <= 0xff; a++ {

		for b := 0; b <= 0xff; b++ {

			// Add a (accum) to b (operand)

			proc := prepareInstructionImmediate(t, instructions, "ADC", uint8(b))
			proc.registers.A = uint8(a)
			proc.registers.SR.Set(SRFlagC)
			require.Nil(proc.Step())

			expectedResult := uint8((a + b + 1) & 0xff)
			result16 := a + b + 1
			expectedZFlag := expectedResult == 0
			expectedNFlag := (expectedResult & 0x80) > 0
			expectedCFlag := result16 > 0xff

			expectedVFlag := false
			// If both numbers are positive...
			if a&0x80 == 0 && b&0x80 == 0 {
				// Overflow occurrs if the result > 0x7F (number not positive)
				if expectedResult > 0x7f {
					expectedVFlag = true
				}
			}
			// If both numbers are negative...
			if a&0x80 == 0x80 && b&0x80 == 0x80 {
				// Overflow occurss if the result is < 0x80 (number not negative)
				if expectedResult < 0x80 {
					expectedVFlag = true
				}
			}

			assert.Equal(expectedResult, proc.registers.A, fmt.Sprintf("%d+%d should = %d", a, b, expectedResult))
			assert.Equal(expectedZFlag, proc.registers.SR.IsSet(SRFlagZ), fmt.Sprintf("%d+%d=%d, expected Z flag incorrect", a, b, expectedResult))
			assert.Equal(expectedNFlag, proc.registers.SR.IsSet(SRFlagN), fmt.Sprintf("%d+%d=%d, expected N flag incorrect", a, b, expectedResult))
			assert.Equal(expectedCFlag, proc.registers.SR.IsSet(SRFlagC), fmt.Sprintf("%d+%d=%d, expected C flag incorrect", a, b, expectedResult))
			assert.Equal(expectedVFlag, proc.registers.SR.IsSet(SRFlagV), fmt.Sprintf("%d+%d=%d, expected V flag incorrect", a, b, expectedResult))

			proc.registers.SR.Clear(SRFlagD) // Decimal mode

		}
	}

}

func TestFromBCD(t *testing.T) {

	assert := assert.New(t)
	for i := 0; i <= 99; i++ {
		high := uint8(i / 10)
		low := uint8(i % 10)

		bcd := (high << 4) | low

		assert.Equal(i, fromBCD(bcd))

	}

}

func TestToBCD(t *testing.T) {
	assert := assert.New(t)
	for i := 0; i <= 99; i++ {

		bcdVal := toBCD(i)
		assert.Equal(i, fromBCD(bcdVal))
	}

}

func TestDecimalADCWithoutCarry(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			proc := prepareInstructionImmediate(t, instructions, "ADC", toBCD(b))
			proc.registers.A = toBCD(a)
			proc.registers.SR.Clear(SRFlagC)
			proc.registers.SR.Set(SRFlagD)
			require.Nil(proc.Step())

			expectedResult := (a + b) % 100
			expectedResultBCDEncoded := toBCD(expectedResult)

			expectedZFlag := expectedResultBCDEncoded == 0
			expectedNFlag := (expectedResultBCDEncoded & 0x80) > 0
			expectedCFlag := (a + b) > 99

			assert.EqualValues(expectedResultBCDEncoded, proc.registers.A, fmt.Sprintf("%d+%d should = %02x", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedZFlag, proc.registers.SR.IsSet(SRFlagZ), fmt.Sprintf("%d+%d=%02x, expected Z flag incorrect", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedNFlag, proc.registers.SR.IsSet(SRFlagN), fmt.Sprintf("%d+%d=%02x, expected N flag incorrect", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedCFlag, proc.registers.SR.IsSet(SRFlagC), fmt.Sprintf("%d+%d=%02x, expected C flag incorrect", a, b, expectedResultBCDEncoded))

			proc.registers.SR.Clear(SRFlagD) // Decimal mode

		}
	}

}

func TestDecimalADCWithCarry(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			proc := prepareInstructionImmediate(t, instructions, "ADC", toBCD(b))
			proc.registers.A = toBCD(a)
			proc.registers.SR.Set(SRFlagC)
			proc.registers.SR.Set(SRFlagD)
			require.Nil(proc.Step())

			expectedResult := (a + b + 1) % 100
			expectedResultBCDEncoded := toBCD(expectedResult)

			expectedZFlag := expectedResultBCDEncoded == 0
			expectedNFlag := (expectedResultBCDEncoded & 0x80) > 0
			expectedCFlag := (a + b + 1) > 99

			assert.EqualValues(expectedResultBCDEncoded, proc.registers.A, fmt.Sprintf("%d+%d+1 should = %02x", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedZFlag, proc.registers.SR.IsSet(SRFlagZ), fmt.Sprintf("%d+%d+1=%02x, expected Z flag incorrect", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedNFlag, proc.registers.SR.IsSet(SRFlagN), fmt.Sprintf("%d+%d+1=%02x, expected N flag incorrect", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedCFlag, proc.registers.SR.IsSet(SRFlagC), fmt.Sprintf("%d+%d+1=%02x, expected C flag incorrect", a, b, expectedResultBCDEncoded))

			proc.registers.SR.Clear(SRFlagD) // Decimal mode

		}
	}
}
func TestDecimalSBCWithoutCarry(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			proc := prepareInstructionImmediate(t, instructions, "SBC", toBCD(b))
			proc.registers.A = toBCD(a)
			proc.registers.SR.Set(SRFlagC)
			proc.registers.SR.Set(SRFlagD)
			require.Nil(proc.Step())

			expectedCFlag := true
			expectedResult := (a - b)
			if expectedResult < 0 {
				expectedResult += 100
				expectedCFlag = false

			}

			expectedResultBCDEncoded := toBCD(expectedResult)

			expectedZFlag := expectedResultBCDEncoded == 0
			expectedNFlag := (expectedResultBCDEncoded & 0x80) > 0

			assert.EqualValues(expectedResultBCDEncoded, proc.registers.A, fmt.Sprintf("%d-%d should = %02x", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedZFlag, proc.registers.SR.IsSet(SRFlagZ), fmt.Sprintf("%d-%d=%02x, expected Z flag incorrect", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedNFlag, proc.registers.SR.IsSet(SRFlagN), fmt.Sprintf("%d-%d=%02x, expected N flag incorrect", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedCFlag, proc.registers.SR.IsSet(SRFlagC), fmt.Sprintf("%d-%d=%02x, expected C flag incorrect", a, b, expectedResultBCDEncoded))

			proc.registers.SR.Clear(SRFlagD) // Decimal mode

		}
	}

}

func TestDecimalSBCWithCarry(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			proc := prepareInstructionImmediate(t, instructions, "SBC", toBCD(b))
			proc.registers.A = toBCD(a)
			proc.registers.SR.Clear(SRFlagC)
			proc.registers.SR.Set(SRFlagD)
			require.Nil(proc.Step())

			expectedCFlag := true
			expectedResult := (a - b - 1)
			if expectedResult < 0 {
				expectedResult += 100
				expectedCFlag = false

			}
			expectedResultBCDEncoded := toBCD(expectedResult)

			expectedZFlag := expectedResultBCDEncoded == 0
			expectedNFlag := (expectedResultBCDEncoded & 0x80) > 0

			assert.EqualValues(expectedResultBCDEncoded, proc.registers.A, fmt.Sprintf("%d-%d+1 should = %02x", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedZFlag, proc.registers.SR.IsSet(SRFlagZ), fmt.Sprintf("%d-%d-1=%02x, expected Z flag incorrect", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedNFlag, proc.registers.SR.IsSet(SRFlagN), fmt.Sprintf("%d-%d-1=%02x, expected N flag incorrect", a, b, expectedResultBCDEncoded))
			assert.Equal(expectedCFlag, proc.registers.SR.IsSet(SRFlagC), fmt.Sprintf("%d-%d-1=%02x, expected C flag incorrect", a, b, expectedResultBCDEncoded))

			proc.registers.SR.Clear(SRFlagD) // Decimal mode

		}
	}
}

func fromBCD(val uint8) int {
	return int((val>>4)*10 + (val & 0x0f))
}

func toBCD(val int) uint8 {
	return ((uint8(val) / 10) << 4) | (uint8(val) % 10)
}
