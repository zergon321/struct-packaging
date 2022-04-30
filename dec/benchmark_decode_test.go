package main_dec_test

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"math"
	"struct-packaging/pb"
	"testing"
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

func BenchmarkJSON(b *testing.B) {
	characterID := uuid.New()
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data, _ := json.Marshal(mv)
	var newMv Movement

	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &newMv)
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
	enc.Encode(mv)

	var newMv Movement
	dec := gob.NewDecoder(buffer)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		dec.Decode(&newMv)
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

	binary.Write(buffer, binary.LittleEndian, mv.Opcode)
	binary.Write(buffer, binary.LittleEndian, mv.CharacterID[:])
	binary.Write(buffer, binary.LittleEndian, mv.X)
	binary.Write(buffer, binary.LittleEndian, mv.Y)
	binary.Write(buffer, binary.LittleEndian, mv.Z)

	var newMv Movement

	for i := 0; i < b.N; i++ {
		buffer.Reset()

		binary.Read(buffer, binary.LittleEndian, &newMv.Opcode)
		binary.Read(buffer, binary.LittleEndian, newMv.CharacterID[:])
		binary.Read(buffer, binary.LittleEndian, &mv.X)
		binary.Read(buffer, binary.LittleEndian, &mv.Y)
		binary.Read(buffer, binary.LittleEndian, &mv.Z)
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

	binary.Write(buffer, binary.BigEndian, mv.Opcode)
	binary.Write(buffer, binary.BigEndian, mv.CharacterID[:])
	binary.Write(buffer, binary.BigEndian, mv.X)
	binary.Write(buffer, binary.BigEndian, mv.Y)
	binary.Write(buffer, binary.BigEndian, mv.Z)

	var newMv Movement

	for i := 0; i < b.N; i++ {
		buffer.Reset()

		binary.Read(buffer, binary.BigEndian, &newMv.Opcode)
		binary.Read(buffer, binary.BigEndian, newMv.CharacterID[:])
		binary.Read(buffer, binary.BigEndian, &mv.X)
		binary.Read(buffer, binary.BigEndian, &mv.Y)
		binary.Read(buffer, binary.BigEndian, &mv.Z)
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

	binary.LittleEndian.PutUint32(data, uint32(mv.Opcode))

	for i := 4; i < 20; i++ {
		data[i] = mv.CharacterID[i-4]
	}

	binary.LittleEndian.PutUint64(data[20:], math.Float64bits(mv.X))
	binary.LittleEndian.PutUint64(data[28:], math.Float64bits(mv.Y))
	binary.LittleEndian.PutUint64(data[36:], math.Float64bits(mv.Z))

	var newMv Movement

	for i := 0; i < b.N; i++ {
		newMv.Opcode = int32(binary.LittleEndian.Uint32(data))

		for i := 4; i < 20; i++ {
			newMv.CharacterID[i-4] = data[i]
		}

		newMv.X = math.Float64frombits(binary.LittleEndian.Uint64(data[20:]))
		newMv.Y = math.Float64frombits(binary.LittleEndian.Uint64(data[28:]))
		newMv.Z = math.Float64frombits(binary.LittleEndian.Uint64(data[36:]))
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

	binary.BigEndian.PutUint32(data, uint32(mv.Opcode))

	for i := 4; i < 20; i++ {
		data[i] = mv.CharacterID[i-4]
	}

	binary.BigEndian.PutUint64(data[20:], math.Float64bits(mv.X))
	binary.BigEndian.PutUint64(data[28:], math.Float64bits(mv.Y))
	binary.BigEndian.PutUint64(data[36:], math.Float64bits(mv.Z))

	var newMv Movement

	for i := 0; i < b.N; i++ {
		newMv.Opcode = int32(binary.BigEndian.Uint32(data))

		for i := 4; i < 20; i++ {
			newMv.CharacterID[i-4] = data[i]
		}

		newMv.X = math.Float64frombits(binary.BigEndian.Uint64(data[20:]))
		newMv.Y = math.Float64frombits(binary.BigEndian.Uint64(data[28:]))
		newMv.Z = math.Float64frombits(binary.BigEndian.Uint64(data[36:]))
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

	data, _ := proto.Marshal(&mv)

	var newMv pb.Movement

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(data, &newMv)
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

	builder.PlaceInt32(mv.Opcode)
	builder.CreateByteVector(mv.CharacterID[:])
	builder.PlaceFloat64(mv.X)
	builder.PlaceFloat64(mv.Y)
	builder.PlaceFloat64(mv.Z)

	builder.Finish(0)
	builder.FinishedBytes()

	var newMv Movement

	for i := 0; i < b.N; i++ {
		flatbuffers.GetInt32()
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
