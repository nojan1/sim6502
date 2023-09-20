package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusRegisters(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	type TestSpec struct {
		flag     SRFlag
		flagID   string
		setValue uint8
	}

	tests := []TestSpec{
		{SRFlagN, "N", 0x80},
		{SRFlagV, "V", 0x40},
		{SRFlagB, "B", 0x10},
		{SRFlagD, "D", 0x08},
		{SRFlagI, "I", 0x04},
		{SRFlagZ, "Z", 0x02},
		{SRFlagC, "C", 0x01},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test_Flag_%s", test.flagID), func(t *testing.T) {
			sr := StatusRegister{}
			require.Zero(sr.Value())
			assert.False(sr.IsSet(test.flag))
			sr.Set(test.flag)
			assert.EqualValues(test.setValue, sr.Value())
			assert.True(sr.IsSet(test.flag))
			sr.Clear(test.flag)
			assert.EqualValues(0x00, sr.Value())
			assert.False(sr.IsSet(test.flag))
		})
	}
}
