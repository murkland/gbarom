package lz77

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var ErrInvalid = errors.New("invalid lz77 data")

func Decompress(r io.Reader) ([]byte, error) {
	var header uint32
	if err := binary.Read(r, binary.LittleEndian, &header); err != nil {
		return nil, err
	}

	if header&0xFF != 0x10 {
		return nil, ErrInvalid
	}

	n := int(header >> 8)

	var out bytes.Buffer
	for out.Len() < n {
		var ref uint8
		if err := binary.Read(r, binary.LittleEndian, &ref); err != nil {
			return out.Bytes(), err
		}

		for i := 0; i < 8 && out.Len() < n; i++ {
			if (ref & (0x80 >> i)) == 0 {
				if _, err := io.CopyN(&out, r, 1); err != nil {
					return out.Bytes(), err
				}
				continue
			}

			var info uint16
			if err := binary.Read(r, binary.BigEndian, &info); err != nil {
				return nil, err
			}

			m := int(info >> 12)
			offset := int(info & 0x0FFF)

			for j := 0; j < m+3; j++ {
				out.WriteByte(out.Bytes()[out.Len()-offset-1])
			}
		}
	}

	return out.Bytes()[:n], nil
}
