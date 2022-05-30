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

type Movement65536 struct {
	Opcode      int32     `json:"opcode"       yaml:"opcode"       xml:"opcode"       cbor:"opcode"       msgpack:"opcode"       bson:"opcode"      `
	CharacterID [65536]byte `json:"character_id" yaml:"character_id" xml:"character_id" cbor:"character_id" msgpack:"character_id" bson:"character_id"`
	X           float64   `json:"x"            yaml:"x"            xml:"x"            cbor:"x"            msgpack:"x"            bson:"x"           `
	Y           float64   `json:"y"            yaml:"y"            xml:"y"            cbor:"y"            msgpack:"y"            bson:"y"           `
	Z           float64   `json:"z"            yaml:"z"            xml:"z"            cbor:"z"            msgpack:"z"            bson:"z"           `
}

type Movement65536Alt struct {
	Opcode      int32     `xml:"opcode,attr"      `
	CharacterID [65536]byte `xml:"character_id,attr"`
	X           float64   `xml:"x,attr"           `
	Y           float64   `xml:"y,attr"           `
	Z           float64   `xml:"z,attr"           `
}

const (
	movement65536Size = int(unsafe.Sizeof(Movement65536{}))
)

func Benchmark65536JSON(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536YAML(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536XML(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536XMLAlt(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536Alt{
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

func Benchmark65536Gob(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	buffer := bytes.NewBuffer(make([]byte, 0, movement65536Size))
	enc := gob.NewEncoder(buffer)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		enc.Encode(mv)
	}
}

func Benchmark65536Msgpack(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536BSON(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536CBORCanonicalOptions(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536CBORCTAP2Options(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536CBORCoreDetOptions(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536CBORPreferredUnsortedOptions(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
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

func Benchmark65536Binary(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement65536Size)
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

func Benchmark65536BinaryBigEndian(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement65536Size)
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

func Benchmark65536BinaryWholeStruct(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement65536Size)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		binary.Write(buffer, binary.LittleEndian, mv)
	}
}

func Benchmark65536BinaryWholeStructBigEndian(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movement65536Size)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		binary.Write(buffer, binary.BigEndian, mv)
	}
}

func Benchmark65536BinaryNoReflection(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, movement65536Size)

	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint32(data, uint32(mv.Opcode))
		copy(data[4:65540], mv.CharacterID[:])
		binary.LittleEndian.PutUint64(data[65540:], math.Float64bits(mv.X))
		binary.LittleEndian.PutUint64(data[65548:], math.Float64bits(mv.Y))
		binary.LittleEndian.PutUint64(data[65556:], math.Float64bits(mv.Z))
	}
}

func Benchmark65536BinaryBigEndianNoReflection(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, movement65536Size)

	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(data, uint32(mv.Opcode))
		copy(data[4:65540], mv.CharacterID[:])
		binary.BigEndian.PutUint64(data[65540:], math.Float64bits(mv.X))
		binary.BigEndian.PutUint64(data[65548:], math.Float64bits(mv.Y))
		binary.BigEndian.PutUint64(data[65556:], math.Float64bits(mv.Z))
	}
}

func Benchmark65536Protobuf(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
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

func Benchmark65536FlatBuffers(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	builder := flatbuffers.NewBuilder(movement65536Size)

	for i := 0; i < b.N; i++ {
		builder.Reset()

		fb.CreateMovement(builder, mv.Opcode, mv.CharacterID[:], mv.X, mv.Y, mv.Z)
	}
}

func Benchmark65536Unsafe(b *testing.B) {
	rand.Seed(65536)
	var randomData [65536]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement65536{
		Opcode:      32,
		CharacterID: characterID,
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		_ = (*[movement65536Size]byte)(unsafe.Pointer(&mv))[:]
	}
}