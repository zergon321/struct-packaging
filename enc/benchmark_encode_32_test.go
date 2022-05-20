package main_enc_test

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"math"
	"math/rand"
	"struct-packaging/fb"
	"struct-packaging/pb"
	"testing"
	"unsafe"

	"github.com/fxamacker/cbor"
	flatbuffers "github.com/google/flatbuffers/go"
	msgpack "github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v2"
)

type Movement32 struct {
	Opcode      int32    `json:"opcode"       yaml:"opcode"       xml:"opcode"       cbor:"opcode"       msgpack:"opcode"       bson:"opcode"      `
	CharacterID [32]byte `json:"character_id" yaml:"character_id" xml:"character_id" cbor:"character_id" msgpack:"character_id" bson:"character_id"`
	X           float64  `json:"x"            yaml:"x"            xml:"x"            cbor:"x"            msgpack:"x"            bson:"x"           `
	Y           float64  `json:"y"            yaml:"y"            xml:"y"            cbor:"y"            msgpack:"y"            bson:"y"           `
	Z           float64  `json:"z"            yaml:"z"            xml:"z"            cbor:"z"            msgpack:"z"            bson:"z"           `
}

type Movement32Alt struct {
	Opcode      int32    `xml:"opcode,attr"      `
	CharacterID [32]byte `xml:"character_id,attr"`
	X           float64  `xml:"x,attr"           `
	Y           float64  `xml:"y,attr"           `
	Z           float64  `xml:"z,attr"           `
}

const (
	movement32Size = int(unsafe.Sizeof(Movement{}))
)

func Benchmark32JSON(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(mv)
	}
}

func Benchmark32YAML(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		yaml.Marshal(mv)
	}
}

func Benchmark32XML(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		xml.Marshal(mv)
	}
}

func Benchmark32XMLAlt(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32Alt{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		xml.Marshal(mv)
	}
}

func Benchmark32Gob(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	buffer := bytes.NewBuffer(make([]byte, 0, movement32Size))
	enc := gob.NewEncoder(buffer)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		enc.Encode(mv)
	}
}

func Benchmark32Msgpack(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		msgpack.Marshal(mv)
	}
}

func Benchmark32BSON(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		bson.MarshalValue(mv)
	}
}

func Benchmark32CBORCanonicalOptions(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		cbor.Marshal(mv,
			cbor.CanonicalEncOptions())
	}
}

func Benchmark32CBORCTAP2Options(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		cbor.Marshal(mv,
			cbor.CTAP2EncOptions())
	}
}

func Benchmark32CBORCoreDetOptions(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		cbor.Marshal(mv,
			cbor.CoreDetEncOptions())
	}
}

func Benchmark32CBORPreferredUnsortedOptions(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		cbor.Marshal(mv,
			cbor.PreferredUnsortedEncOptions())
	}
}

func Benchmark32Binary(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement32Size)
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

func Benchmark32BinaryBigEndian(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement32Size)
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

func Benchmark32BinaryWholeStruct(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement32Size)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		binary.Write(buffer, binary.LittleEndian, mv)
	}
}

func Benchmark32BinaryWholeStructBigEndian(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement32Size)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		binary.Write(buffer, binary.BigEndian, mv)
	}
}

func Benchmark32BinaryNoReflection(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, movement32Size)

	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint32(data, uint32(mv.Opcode))
		copy(data[4:516], mv.CharacterID[:])
		binary.LittleEndian.PutUint64(data[516:], math.Float64bits(mv.X))
		binary.LittleEndian.PutUint64(data[524:], math.Float64bits(mv.Y))
		binary.LittleEndian.PutUint64(data[532:], math.Float64bits(mv.Z))
	}
}

func Benchmark32BinaryBigEndianNoReflection(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, movement32Size)

	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(data, uint32(mv.Opcode))
		copy(data[4:516], mv.CharacterID[:])
		binary.BigEndian.PutUint64(data[516:], math.Float64bits(mv.X))
		binary.BigEndian.PutUint64(data[524:], math.Float64bits(mv.Y))
		binary.BigEndian.PutUint64(data[532:], math.Float64bits(mv.Z))
	}
}

func Benchmark32Protobuf(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
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

func Benchmark32FlatBuffers(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	builder := flatbuffers.NewBuilder(movement32Size)

	for i := 0; i < b.N; i++ {
		builder.Reset()

		fb.CreateMovement(builder, mv.Opcode, mv.CharacterID[:], mv.X, mv.Y, mv.Z)
	}
}

func Benchmark32Unsafe(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		_ = (*[movement32Size]byte)(unsafe.Pointer(&mv))[:]
	}
}
