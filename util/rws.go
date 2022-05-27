package util

import (
	"fmt"
	"io"
)

// BytesReadWriteSeeker implements Read, Write,
// Seek and ReadByte functions for byte slice.
type BytesReadWriteSeeker struct {
	Data []byte
	Pos  int
}

// Seek moves the pointer to the specified position.
func (brws *BytesReadWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	var start int

	switch whence {
	case io.SeekStart:
		start = 0

	case io.SeekCurrent:
		start = brws.Pos

	case io.SeekEnd:
		start = len(brws.Data)

	default:
		return -1, fmt.Errorf("option not defined")
	}

	newPos := start + int(offset)

	switch {
	case newPos < 0:
		newPos = 0

	case newPos > len(brws.Data):
		newPos = len(brws.Data)
	}

	brws.Pos = newPos

	return int64(brws.Pos), nil
}

// Write writes the data to the inner buffer.
func (brws *BytesReadWriteSeeker) Write(p []byte) (n int, err error) {
	offset := len(p)
	brws.Data = append(brws.Data, p...)
	brws.Pos += offset

	return offset, nil
}

// ReadByte reads exactly 1 byte from the inner buffer.
func (brws *BytesReadWriteSeeker) ReadByte() (byte, error) {
	if brws.Pos == len(brws.Data) {
		return 0, io.EOF
	}

	value := brws.Data[brws.Pos]
	brws.Pos++

	return value, nil
}

// Read reads the bytes from the inner buffer to the provided buffer.
func (brws *BytesReadWriteSeeker) Read(p []byte) (n int, err error) {
	if brws.Pos == len(brws.Data) {
		return -1, io.EOF
	}

	offset := len(p)

	if offset > len(brws.Data)-brws.Pos {
		offset = len(brws.Data) - brws.Pos
	}

	copy(p, brws.Data[brws.Pos:brws.Pos+offset])
	brws.Pos += offset

	return offset, nil
}
