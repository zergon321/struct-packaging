package main

import (
	"fmt"
	"unsafe"
)

type Movement struct {
	Opcode int
	X      float64
	Y      float64
	Z      float64
	K      int32
}

const (
	movementSize = int(unsafe.Sizeof(Movement{}))
)

func up(mv Movement) []byte {
	// leaking parameter (closure)
	return (*[movementSize]byte)(unsafe.Pointer(&mv))[:]
}

func main() {
	mv := Movement{
		Opcode: 32,
		X:      13.34,
		Y:      20.36,
		Z:      45.13,
		K:      16,
	}
	packet := (*[movementSize]byte)(unsafe.Pointer(&mv))[:]

	fmt.Println(mv)
	fmt.Println(len(packet))
	fmt.Println(packet)

	newMv := *(*Movement)(unsafe.Pointer(&packet[0]))

	fmt.Println(newMv)

	data := up(mv)

	fmt.Println(data)
}
