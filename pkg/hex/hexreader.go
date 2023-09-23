package hex

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

// HexChunk represents a chunk of hex data and the address associated with it
type HexChunk struct {
	// Addr is the address for the data
	Addr uint16

	// Data is the raw binary data
	Data []uint8
}

// HexReader can read intel hex file format, at least as far as needed for 8 bit
type HexReader struct {
	// Chunks are the various chunks of data loaded by the reader
	Chunks []HexChunk
}

// NexHexReader will return a hex reader with the content from the specified
// reader decoded from intel hex format into individual chunks
// Supports only I8HEX hex files (records 0,1)
func NewHexReader(r io.Reader) (*HexReader, error) {
	hr := &HexReader{}
	ln := 0

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		line := sc.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			// Line has no colon, we don't want it
			continue
		}
		ln++
		hexContent := strings.TrimSpace(parts[1])
		data, err := hex.DecodeString(hexContent)
		if err != nil {
			return nil, fmt.Errorf("HEX error, invalid HEX content, line %d", ln)
		}

		if len(data) < 5 {
			return nil, fmt.Errorf("HEX error, line too short, line %d", ln)
		}

		bc := data[0]
		addr := (uint16(data[1]) << 8) | uint16(data[2])
		rt := data[3]
		cs := calculateChecksum(data)
		if data[len(data)-1] != cs {
			return nil, fmt.Errorf("checksum error, line %d", ln)
		}
		switch rt {
		case 0x01:
			// EOF
			return hr, nil
		case 0x00:
			// 8 bit data
			if len(data) != int(bc)+5 {
				return nil, fmt.Errorf("HEX error, line %d, line length incorrect", ln)
			}
			ch := HexChunk{
				Addr: addr,
				Data: data[4 : len(data)-1],
			}
			hr.Chunks = append(hr.Chunks, ch)

		case 0x02:
			return nil, fmt.Errorf("HEX error, record type 2 (extended segment address) unsupported, line %d", ln)
		case 0x03:
			return nil, fmt.Errorf("HEX error, record type 3 (start segment address) unsupported, line %d", ln)
		case 0x04:
			return nil, fmt.Errorf("HEX error, record type 4 (extended linear address) unsupported, line %d", ln)
		case 0x05:
			return nil, fmt.Errorf("HEX error, record type 5 (start linear address) unsupported, line %d", ln)
		default:
			return nil, fmt.Errorf("HEX error, unsupported record type 0x%02x, line %d", rt, ln)
		}
	}

	return nil, errors.New("HEX error, no EOF marker found")

}

// calculateChecksum calculates the checksum for a hex record
func calculateChecksum(data []byte) uint8 {
	var cs uint8
	for _, b := range data[:len(data)-1] {
		cs += uint8(b)
	}
	return (cs ^ 0xff) + 1
}
