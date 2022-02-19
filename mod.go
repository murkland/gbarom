package gbarom

import (
	"io"
	"os"
	"strings"
)

func ReadROMID(r io.ReadSeeker) (string, error) {
	var romID [4]byte
	if _, err := r.Seek(0x000000AC, os.SEEK_SET); err != nil {
		return "", err
	}

	if _, err := io.ReadFull(r, romID[:]); err != nil {
		return "", err
	}

	return string(romID[:]), nil
}

func ReadROMTitle(r io.ReadSeeker) (string, error) {
	var romTitle [12]byte
	if _, err := r.Seek(0x000000A0, os.SEEK_SET); err != nil {
		return "", err
	}

	if _, err := io.ReadFull(r, romTitle[:]); err != nil {
		return "", err
	}

	return strings.TrimRight(string(romTitle[:]), "\x00"), nil
}
