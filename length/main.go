package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"struct-packaging/fb"
	"struct-packaging/pb"
	"unsafe"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type Movement struct {
	Opcode      int32
	CharacterID [16]byte
	X           float64
	Y           float64
	Z           float64
}

const (
	movementSize = int(unsafe.Sizeof(Movement{}))
)

func main() {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	pbmv := pb.Movement{
		Opcode:      32,
		CharacterID: characterID[:],
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	jsonData, err := json.Marshal(mv)
	handleError(err)

	buffer := bytes.NewBuffer(make([]byte, 0, movementSize))
	buffer.Reset()
	enc := gob.NewEncoder(buffer)
	err = enc.Encode(mv)
	handleError(err)
	gobData := buffer.Bytes()

	buffer = bytes.NewBuffer(make([]byte, movementSize))
	buffer.Reset()
	err = binary.Write(buffer, binary.LittleEndian, mv.Opcode)
	handleError(err)
	err = binary.Write(buffer, binary.LittleEndian, mv.CharacterID[:])
	handleError(err)
	err = binary.Write(buffer, binary.LittleEndian, mv.X)
	handleError(err)
	err = binary.Write(buffer, binary.LittleEndian, mv.Y)
	handleError(err)
	err = binary.Write(buffer, binary.LittleEndian, mv.Z)
	handleError(err)
	binData := buffer.Bytes()

	pbData, err := proto.Marshal(&pbmv)
	handleError(err)

	builder := flatbuffers.NewBuilder(movementSize)
	fbData := fb.CreateMovement(builder, mv.Opcode, mv.CharacterID[:], mv.X, mv.Y, mv.Z)

	unsafeData := (*[movementSize]byte)(unsafe.Pointer(&mv))[:]

	fmt.Println("JSON data bytes length:", len(jsonData))
	fmt.Println("gob data bytes length:", len(gobData))
	fmt.Println("binary data bytes length:", len(binData))
	fmt.Println("Protobuf data bytes length:", len(pbData))
	fmt.Println("flat buffers data bytes length:", len(fbData))
	fmt.Println("unsafe cast data bytes length:", len(unsafeData))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
