package main

import (
	"fmt"
	"net"
	"unsafe"
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

//go:noinline
func down(data []byte) Movement {
	return *(*Movement)(unsafe.Pointer(&data[0]))
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9828")
	handleError(err)
	conn, err := listener.Accept()
	handleError(err)

	for {
		var buffer [movSize]byte
		total := 0

		for total < int(movSize) {
			read, err := conn.Read(buffer[total:])
			handleError(err)
			total += read
		}

		mov := down(buffer[:])
		fmt.Println(mov)
		fmt.Println(buffer)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
