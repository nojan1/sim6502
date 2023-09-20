package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwosComplement(t *testing.T) {
	for i := 0; i <= 127; i++ {
		t.Run(fmt.Sprintf("2C_%d", i), func(t *testing.T) {
			assert := assert.New(t)
			assert.EqualValues(-i, comp2(uint8(i)))
			assert.EqualValues(i, comp2(uint8(-i)))
		})
	}
}

func TestSamePage(t *testing.T) {
	assert := assert.New(t)
	assert.True(samePage(0x8050, 0x809f))
	assert.False(samePage(0x8050, 0x8100))
}
