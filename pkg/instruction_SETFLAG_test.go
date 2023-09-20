package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests all set flag instructions

type SetTestScenario struct {
	name string
	op   string
	flag SRFlag
	set  bool
}

var SetTests = []SetTestScenario{
	{name: "SEC Set", op: "SEC", flag: SRFlagC, set: true},
	{name: "SEC Not Set", op: "SEC", flag: SRFlagC, set: false},
	{name: "SEI Set", op: "SEI", flag: SRFlagI, set: true},
	{name: "SEI Not Set", op: "SEI", flag: SRFlagI, set: false},
	{name: "SED Set", op: "SED", flag: SRFlagD, set: true},
	{name: "SED Not Set", op: "SED", flag: SRFlagD, set: false},
}

func TestSetting(t *testing.T) {

	for _, test := range SetTests {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)
			proc := prepareInstructionImplied(t, instructions, test.op)
			proc.registers.SR.SetTo(test.flag, test.set)
			require.Equal(test.set, proc.registers.SR.IsSet(test.flag), "register was not correctly set")
			err, _ := proc.Step()
			require.Nil(err, fmt.Sprintf("error on step: %v", err))
			assert.True(proc.registers.SR.IsSet(test.flag), "register was not set")
		})
	}
}
