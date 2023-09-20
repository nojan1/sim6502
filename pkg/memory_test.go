package sim6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleMemory(t *testing.T) {
	assert := assert.New(t)
	mem := MappableMemory{}
	mem.Clear()

	for i := 0; i < 65536; i++ {
		mem.Write(uint16(i), uint8(i&0xff))
		assert.Equal(uint8(i&0xff), mem.Read(uint16(i)), fmt.Sprintf("at mem location 0x%04x", i))
	}
}

type MemoryMapTester struct {
	val                       uint8
	calledWrites, calledReads int
}

func (m *MemoryMapTester) AddressRange() []MappedMemoryAddressRange {
	return []MappedMemoryAddressRange{{0x5000, 0x5000}}
}

func (m *MemoryMapTester) Write(addr uint16, val uint8) {
	m.val = val
	m.calledWrites++
}

func (m *MemoryMapTester) Read(addr uint16) uint8 {
	m.calledReads++
	return m.val
}

func TestMappedMemory(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	mem := MappableMemory{}
	mem.Clear()
	mapper := &MemoryMapTester{}
	mem.Map(mapper)

	require.EqualValues(mem.Read(0x5000), 0)

	mem.Write(0x5000, 55)
	assert.EqualValues(55, mem.Read(0x5000))

	mem.Read(0x5001)
	mem.Write(0x5001, 99)

	assert.Equal(2, mapper.calledReads)
	assert.Equal(1, mapper.calledWrites)
}
