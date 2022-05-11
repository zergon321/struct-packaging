package main

import (
	"net"
	"unsafe"

	"github.com/google/uuid"
)

type Movement struct {
	Opcode      int32
	CharacterID [16]byte
	X           float64
	Y           float64
	Z           float64
}

const (
	movSize = uint64(unsafe.Sizeof(Movement{}))
)

func main() {
	characterID := uuid.New()
	conn, err := net.Dial("tcp", "127.0.0.1:9828")
	handleError(err)

	for {
		mov := Movement{
			Opcode:      16,
			CharacterID: characterID,
			X:           16.36,
			Y:           19.45,
			Z:           87.51,
		}

		data := (*[movSize]byte)(unsafe.Pointer(&mov))[:]
		_, err = conn.Write(data)
		handleError(err)

		//fmt.Println(data)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
