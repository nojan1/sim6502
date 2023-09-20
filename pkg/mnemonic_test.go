package sim6502

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests that all instructions are assigned a mnemonic that matches their instruction struct type
// .. ensures against common typos

func TestMnemonics(t *testing.T) {
	instrs := make([]*instruction, len(instructions))
	copy(instrs, instructions)

	loadExtendedInstructions(instrs)

	for _, i := range instrs {

		if i == nil {
			continue
		}
		instructionName := strings.ToUpper(reflect.TypeOf(i.Impl).Elem().Name())
		t.Run(fmt.Sprintf("Check_Mnemonic_%s", instructionName), func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(instructionName, i.Impl.Mnemonic(), fmt.Sprintf("Mnemonic of instruction %s is incorrectly defined", instructionName))
		})
	}
}
