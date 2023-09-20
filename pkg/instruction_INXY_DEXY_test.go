package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests INX, INY, DEX, DEY operations
type InxyDexyTestScenario struct {
	name          string
	op            string
	set           func(*Processor, uint8)
	get           func(*Processor) uint8
	initialValue  uint8
	expectedValue uint8
}

func getX(proc *Processor) uint8 {
	return proc.registers.X
}
func getY(proc *Processor) uint8 {
	return proc.registers.Y
}

func setX(proc *Processor, val uint8) {
	proc.registers.X = val
}
func setY(proc *Processor, val uint8) {
	proc.registers.Y = val
}

var InxyDexyTests = []InxyDexyTestScenario{
	{"INX 0 1", "INX", setX, getX, 0, 1},
	{"INX 254 255", "INX", setX, getX, 0xFE, 0xFF},
	{"INX 255 0", "INX", setX, getX, 0xFF, 0x00},
	{"INY 0 1", "INY", setY, getY, 0, 1},
	{"INY 254 255", "INY", setY, getY, 0xFE, 0xFF},
	{"INY 255 0", "INY", setY, getY, 0xFF, 0x00},
	{"DEX 1 0", "DEX", setX, getX, 1, 0},
	{"DEX 0 255", "DEX", setX, getX, 0, 0xFF},
	{"DEX 255 254", "DEX", setX, getX, 0xFF, 0xFE},
}

func TestInxyDexy(t *testing.T) {
	for _, test := range InxyDexyTests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			proc := prepareInstructionImplied(t, instructions, test.op)
			test.set(proc, test.initialValue)
			err, _ := proc.Step()
			require.Nil(err, fmt.Sprintf("error on step: %v", err))

			val := test.get(proc)

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
