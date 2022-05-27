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

type Movement128 struct {
	Opcode      int32     `json:"opcode"       yaml:"opcode"       xml:"opcode"       cbor:"opcode"       msgpack:"opcode"       bson:"opcode"      `
	CharacterID [128]byte `json:"character_id" yaml:"character_id" xml:"character_id" cbor:"character_id" msgpack:"character_id" bson:"character_id"`
	X           float64   `json:"x"            yaml:"x"            xml:"x"            cbor:"x"            msgpack:"x"            bson:"x"           `
	Y           float64   `json:"y"            yaml:"y"            xml:"y"            cbor:"y"            msgpack:"y"            bson:"y"           `
	Z           float64   `json:"z"            yaml:"z"            xml:"z"            cbor:"z"            msgpack:"z"            bson:"z"           `
}

type Movement128Alt struct {
	Opcode      int32     `xml:"opcode,attr"      `
	CharacterID [128]byte `xml:"character_id,attr"`
	X           float64   `xml:"x,attr"           `
	Y           float64   `xml:"y,attr"           `
	Z           float64   `xml:"z,attr"           `
}

const (
	movement128Size = int(unsafe.Sizeof(Movement128{}))
)

func Benchmark128JSON(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128YAML(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128XML(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128XMLAlt(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128Alt{
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

func Benchmark128Gob(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	buffer := bytes.NewBuffer(make([]byte, 0, movement128Size))
	enc := gob.NewEncoder(buffer)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		enc.Encode(mv)
	}
}

func Benchmark128Msgpack(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128BSON(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128CBORCanonicalOptions(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128CBORCTAP2Options(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128CBORCoreDetOptions(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128CBORPreferredUnsortedOptions(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
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

func Benchmark128Binary(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement128Size)
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

func Benchmark128BinaryBigEndian(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement128Size)
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

func Benchmark128BinaryWholeStruct(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement128Size)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		binary.Write(buffer, binary.LittleEndian, mv)
	}
}

func Benchmark128BinaryWholeStructBigEndian(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement128Size)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		binary.Write(buffer, binary.BigEndian, mv)
	}
}

func Benchmark128BinaryNoReflection(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, movement128Size)

	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint32(data, uint32(mv.Opcode))
		copy(data[4:132], mv.CharacterID[:])
		binary.LittleEndian.PutUint64(data[132:], math.Float64bits(mv.X))
		binary.LittleEndian.PutUint64(data[140:], math.Float64bits(mv.Y))
		binary.LittleEndian.PutUint64(data[148:], math.Float64bits(mv.Z))
	}
}

func Benchmark128BinaryBigEndianNoReflection(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, movement128Size)

	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(data, uint32(mv.Opcode))
		copy(data[4:132], mv.CharacterID[:])
		binary.BigEndian.PutUint64(data[132:], math.Float64bits(mv.X))
		binary.BigEndian.PutUint64(data[140:], math.Float64bits(mv.Y))
		binary.BigEndian.PutUint64(data[148:], math.Float64bits(mv.Z))
	}
}

func Benchmark128Protobuf(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
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

func Benchmark128FlatBuffers(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	builder := flatbuffers.NewBuilder(movement128Size)

	for i := 0; i < b.N; i++ {
		builder.Reset()

		fb.CreateMovement(builder, mv.Opcode, mv.CharacterID[:], mv.X, mv.Y, mv.Z)
	}
}

func Benchmark128Unsafe(b *testing.B) {
	rand.Seed(128)
	var randomData [128]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement128{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		_ = (*[movement128Size]byte)(unsafe.Pointer(&mv))[:]
	}
}