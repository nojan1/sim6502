package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests all relative braching instructions

type BranchTestScenario struct {
	name         string
	op           string
	flag         SRFlag
	set          bool
	shouldBranch bool
}

var BranchTests = []BranchTestScenario{
	{name: "BEQ Branch", op: "BEQ", flag: SRFlagZ, set: true, shouldBranch: true},
	{name: "BEQ No Branch", op: "BEQ", flag: SRFlagZ, set: false, shouldBranch: false},
	{name: "BCC Branch", op: "BCC", flag: SRFlagC, set: false, shouldBranch: true},
	{name: "BCC No Branch", op: "BCC", flag: SRFlagC, set: true, shouldBranch: false},
	{name: "BCS Branch", op: "BCS", flag: SRFlagC, set: true, shouldBranch: true},
	{name: "BCS No Branch", op: "BCS", flag: SRFlagC, set: false, shouldBranch: false},
	{name: "BMI Branch", op: "BMI", flag: SRFlagN, set: true, shouldBranch: true},
	{name: "BMI No Branch", op: "BMI", flag: SRFlagN, set: false, shouldBranch: false},
	{name: "BNE Branch", op: "BNE", flag: SRFlagZ, set: false, shouldBranch: true},
	{name: "BNE No Branch", op: "BNE", flag: SRFlagZ, set: true, shouldBranch: false},
	{name: "BVC Branch", op: "BVC", flag: SRFlagV, set: false, shouldBranch: true},
	{name: "BVC No Branch", op: "BVC", flag: SRFlagV, set: true, shouldBranch: false},
	{name: "BPL Branch", op: "BPL", flag: SRFlagN, set: false, shouldBranch: true},
	{name: "BPL No Branch", op: "BPL", flag: SRFlagN, set: true, shouldBranch: false},
	{name: "BVS Branch", op: "BVS", flag: SRFlagV, set: true, shouldBranch: true},
	{name: "BVS No Branch", op: "BVS", flag: SRFlagV, set: false, shouldBranch: false},
}

func TestBranching(t *testing.T) {

	for _, test := range BranchTests {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)
			proc := prepareInstructionRelative(t, instructions, test.op, 0x05)
			proc.registers.SR.SetTo(test.flag, test.set)
			initPC := proc.registers.PC.Current()
			err, _ := proc.Step()
			require.Nil(err, fmt.Sprintf("error on step: %v", err))

			if test.shouldBranch {
				assert.EqualValues(initPC+7, proc.registers.PC.Current(), "PC did not increment by 7 (2 for the op + 5 offset), indicating no branch took place")
			} else {
				assert.EqualValues(initPC+2, proc.registers.PC.Current(), "PC did not increment by 2 (2 for the op + no branch), indicating a branch took place")
			}
		})
	}
}
