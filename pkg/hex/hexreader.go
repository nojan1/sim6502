package hex

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

type HexChunk struct {
	Addr uint16
	Data []uint8
}

type HexReader struct {
	Chunks []HexChunk
}

func NewHexReader(r io.Reader) (*HexReader, error) {
	hr := &HexReader{}
	ln := 0

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		line := sc.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			// Line has no colon
			continue
		}
		ln++
		hexContent := strings.TrimSpace(parts[1])
		data, err := hex.DecodeString(hexContent)
		if err != nil {
			return nil, fmt.Errorf("HEX invalid HEX content, line %d", ln)
		}

		if len(data) < 5 {
			return nil, fmt.Errorf("HEX error, line %d, line too short", ln)
		}

		bc := data[0]
		addr := (uint16(data[1]) << 8) | uint16(data[2])
		rt := data[3]
		switch rt {
		case 0x01:
			// EOF
			return hr, nil
		case 0x00:
			if len(data) != int(bc)+5 {
				return nil, fmt.Errorf("HEX error, line %d, line length incorrect", ln)
			}
			ch := HexChunk{
				Addr: addr,
				Data: data[4 : len(data)-1],
			}
			hr.Chunks = append(hr.Chunks, ch)
			// TODO: Should probably check the checksum

			// fmt.Printf("Loaded -> 0x%04x: %s\n", addr, hex.EncodeToString(ch.Data))

		case 0x02:
			return nil, errors.New("HEX error, ext seg addr unsupported")
		case 0x03:
			return nil, errors.New("HEX error, start seg addr unsupported")
		case 0x04:
			return nil, errors.New("HEX error, ext linear addr unsupported")
		case 0x05:
			return nil, errors.New("HEX error, start linear addr unsupported")
		default:
			return nil, fmt.Errorf("HEX error, unsupported record type 0x%02x", rt)
		}

	}

	return nil, errors.New("HEX error, no EOF marker")

}
