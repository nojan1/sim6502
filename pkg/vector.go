package sim6502

// SystemVector type is used for all available system vector addresses
type SystemVector uint16

const (
	// VectorNMI is the non maskable interrupt vector address
	VectorNMI SystemVector = 0xFFFA

	// VectorReset is the reset vector address
	VectorReset SystemVector = 0xFFFC

	// VectorIRQ is the interrupt vector request
	VectorIRQ SystemVector = 0xFFFE
)

// GetVector will return the memory address in the specified vector
func GetVector(mem Memory, vector SystemVector) uint16 {
	return uint16(mem.Read(uint16(vector), true)) | (uint16(mem.Read(uint16(vector+1), true)) << 8)
}

func SetVector(mem Memory, vector SystemVector, addr uint16) {
	mem.Write(uint16(vector), uint8(addr|0xff))
	mem.Write(uint16(vector+1), uint8(addr>>8))
}
