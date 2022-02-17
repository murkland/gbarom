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
		var reference uint8
		if err := binary.Read(r, binary.BigEndian, &reference); err != nil {
			return nil, err
		}

		for i := 0; i < 8; i++ {
			if reference&0x80 != 0 {
				var info uint16
				if err := binary.Read(r, binary.BigEndian, &info); err != nil {
					return nil, err
				}

				m := 3 + int(info>>12)
				offset := out.Len() - 1 - int(info&0x0FFF)

				if _, err := out.Write(out.Bytes()[offset : offset+m]); err != nil {
					return nil, err
				}

				if out.Len() >= n {
					// TODO: Maybe check if the length is too long and return an error.
					break
				}
				offset += m
			} else {
				if _, err := io.CopyN(&out, r, 1); err != nil {
					return nil, err
				}
			}
			reference <<= 1
			if out.Len() >= n {
				// TODO: Maybe check if the length is too long and return an error.
				break
			}
		}
	}

	return out.Bytes(), nil
}
