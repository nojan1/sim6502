package sim6502

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CompTestSpec struct {
	op      string
	name    string
	reg     uint8
	comp    uint8
	setter  func(*Processor, uint8)
	expectC bool
	expectN bool
	expectZ bool
}

func setA(p *Processor, val uint8) {
	p.registers.A = val
}

var CompTests = []CompTestSpec{
	{name: "A(50) CMP 50", op: "CMP", reg: 50, comp: 50, setter: setA, expectC: true, expectN: false, expectZ: true},
	{name: "A(51) CMP 50", op: "CMP", reg: 51, comp: 50, setter: setA, expectC: true, expectN: false, expectZ: false},
	{name: "A(49) CMP 50", op: "CMP", reg: 49, comp: 50, setter: setA, expectC: false, expectN: true, expectZ: false},
	{name: "X(50) CPX 50", op: "CPX", reg: 50, comp: 50, setter: setX, expectC: true, expectN: false, expectZ: true},
	{name: "X(51) CPX 50", op: "CPX", reg: 51, comp: 50, setter: setX, expectC: true, expectN: false, expectZ: false},
	{name: "X(49) CPX 50", op: "CPX", reg: 49, comp: 50, setter: setX, expectC: false, expectN: true, expectZ: false},
	{name: "Y(50) CPY 50", op: "CPY", reg: 50, comp: 50, setter: setY, expectC: true, expectN: false, expectZ: true},
	{name: "Y(51) CPY 50", op: "CPY", reg: 51, comp: 50, setter: setY, expectC: true, expectN: false, expectZ: false},
	{name: "Y(49) CPY 50", op: "CPY", reg: 49, comp: 50, setter: setY, expectC: false, expectN: true, expectZ: false},
}

func TestCompare(t *testing.T) {

	for _, test := range CompTests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)
			proc := prepareInstructionImmediate(t, instructions, test.op, test.comp)
			test.setter(proc, test.reg)
			err, _ := proc.Step()
			require.Nil(err)
			assert.EqualValues(test.expectC, proc.registers.SR.IsSet(SRFlagC), "C Flag error")
			assert.EqualValues(test.expectZ, proc.registers.SR.IsSet(SRFlagZ), "Z Flag error")
			assert.EqualValues(test.expectN, proc.registers.SR.IsSet(SRFlagN), "N Flag error")
		})
	}

}
