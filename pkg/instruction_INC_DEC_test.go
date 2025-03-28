package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests INC/DEC mem operations

type IncDecTestScenario struct {
	name          string
	op            string
	initialValue  uint8
	expectedValue uint8
}

var IncDecTests = []IncDecTestScenario{
	{"INC 0 1", "INC", 0, 1},
	{"INC 254 255", "INC", 0xFE, 0xFF},
	{"INC 255 0", "INC", 0xFF, 0x00},
	{"DEC 1 0", "DEC", 1, 0},
	{"DEC 0 255", "DEC", 0, 0xFF},
}

func TestIncDec(t *testing.T) {
	for _, test := range IncDecTests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			proc := prepareInstructionAbsolute(t, instructions, test.op, 0x1234)
			proc.memory.Write(0x1234, test.initialValue)
			err, _ := proc.Step()
			require.Nil(err, fmt.Sprintf("error on step: %v", err))

			val := proc.memory.Read(0x1234, true)

			assert.Equal(test.expectedValue, val)

			if val&0x80 > 0 {
				assert.True(proc.registers.SR.IsSet(SRFlagN), "Negative flag should have been set")
			} else {
				assert.False(proc.registers.SR.IsSet(SRFlagN), "Negative flag should not have been set")
			}
			if val == 0 {
				assert.True(proc.registers.SR.IsSet(SRFlagZ), "Zero flag should have been set")
			} else {
				assert.False(proc.registers.SR.IsSet(SRFlagZ), "Zero flag should not have been set")
			}
		})
	}
}
