package main_enc_test

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"math"
	"struct-packaging/fb"
	"struct-packaging/pb"
	"testing"
	"unsafe"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v2"
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

func BenchmarkJSON(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(mv)
	}
}

func BenchmarkYAML(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		yaml.Marshal(mv)
	}
}

func BenchmarkGob(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	buffer := bytes.NewBuffer(make([]byte, movementSize))
	enc := gob.NewEncoder(buffer)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		enc.Encode(mv)
	}
}

func BenchmarkBinary(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movementSize)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < b.N; i++ {
		buffer.Reset()

		binary.Write(buffer, binary.LittleEndian, mv.Opcode)
		binary.Write(buffer, binary.LittleEndian, mv.CharacterID[:])
		binary.Write(buffer, binary.LittleEndian, mv.X)
		binary.Write(buffer, binary.LittleEndian, mv.Y)
		binary.Write(buffer, binary.LittleEndian, mv.Z)
	}
}

func BenchmarkBinaryBigEndian(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movementSize)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < b.N; i++ {
		buffer.Reset()

		binary.Write(buffer, binary.BigEndian, mv.Opcode)
		binary.Write(buffer, binary.BigEndian, mv.CharacterID[:])
		binary.Write(buffer, binary.BigEndian, mv.X)
		binary.Write(buffer, binary.BigEndian, mv.Y)
		binary.Write(buffer, binary.BigEndian, mv.Z)
	}
}

func BenchmarkBinaryNoReflection(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, movementSize)

	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint32(data, uint32(mv.Opcode))

		for i := 4; i < 20; i++ {
			data[i] = mv.CharacterID[i-4]
		}

		binary.LittleEndian.PutUint64(data[20:], math.Float64bits(mv.X))
		binary.LittleEndian.PutUint64(data[28:], math.Float64bits(mv.Y))
		binary.LittleEndian.PutUint64(data[36:], math.Float64bits(mv.Z))
	}
}

func BenchmarkBinaryBigEndianNoReflection(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, movementSize)

	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(data, uint32(mv.Opcode))

		for i := 4; i < 20; i++ {
			data[i] = mv.CharacterID[i-4]
		}

		binary.BigEndian.PutUint64(data[20:], math.Float64bits(mv.X))
		binary.BigEndian.PutUint64(data[28:], math.Float64bits(mv.Y))
		binary.BigEndian.PutUint64(data[36:], math.Float64bits(mv.Z))
	}
}

func BenchmarkProtobuf(b *testing.B) {
	characterID := uuid.New()
	mv := pb.Movement{
		Opcode:      32,
		CharacterID: characterID[:],
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		proto.Marshal(&mv)
	}
}

func BenchmarkFlatBuffers(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	builder := flatbuffers.NewBuilder(movementSize)

	for i := 0; i < b.N; i++ {
		builder.Reset()

		fb.CreateMovement(builder, mv.Opcode, mv.CharacterID[:], mv.X, mv.Y, mv.Z)
	}
}

func BenchmarkUnsafe(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		_ = (*[movementSize]byte)(unsafe.Pointer(&mv))[:]
	}
}
