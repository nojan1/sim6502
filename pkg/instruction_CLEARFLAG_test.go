package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests all clear flag instructions

type ClearTestScneario struct {
	name string
	op   string
	flag SRFlag
	set  bool
}

var ClearTests = []ClearTestScneario{
	{name: "CLC Set", op: "CLC", flag: SRFlagC, set: true},
	{name: "CLC Not Set", op: "CLC", flag: SRFlagC, set: false},
	{name: "CLD Set", op: "CLD", flag: SRFlagD, set: true},
	{name: "CLD Not Set", op: "CLD", flag: SRFlagD, set: false},
	{name: "CLI Set", op: "CLI", flag: SRFlagI, set: true},
	{name: "CLI Not Set", op: "CLI", flag: SRFlagI, set: false},
	{name: "CLV Set", op: "CLV", flag: SRFlagV, set: true},
	{name: "CLV Not Set", op: "CLV", flag: SRFlagV, set: false},
}

func TestClearing(t *testing.T) {

	for _, test := range ClearTests {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)
			proc := prepareInstructionImplied(t, instructions, test.op)
			proc.registers.SR.SetTo(test.flag, test.set)
			require.Equal(test.set, proc.registers.SR.IsSet(test.flag), "register was not correctly set")
			err, _ := proc.Step()
			require.Nil(err, fmt.Sprintf("error on step: %v", err))
			assert.False(proc.registers.SR.IsSet(test.flag), "register was not cleared")
		})
	}
}
