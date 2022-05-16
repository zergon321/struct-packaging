package main_enc_test

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"math"
	"struct-packaging/fb"
	"struct-packaging/pb"
	"testing"
	"unsafe"

	"github.com/fxamacker/cbor"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
	msgpack "github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v2"
)

type Movement struct {
	Opcode      int32    `json:"opcode"       yaml:"opcode"       xml:"opcode"       cbor:"opcode"       msgpack:"opcode"       bson:"opcode"      `
	CharacterID [16]byte `json:"character_id" yaml:"character_id" xml:"character_id" cbor:"character_id" msgpack:"character_id" bson:"character_id"`
	X           float64  `json:"x"            yaml:"x"            xml:"x"            cbor:"x"            msgpack:"x"            bson:"x"           `
	Y           float64  `json:"y"            yaml:"y"            xml:"y"            cbor:"y"            msgpack:"y"            bson:"y"           `
	Z           float64  `json:"z"            yaml:"z"            xml:"z"            cbor:"z"            msgpack:"z"            bson:"z"           `
}

type MovementAlt struct {
	Opcode      int32    `xml:"opcode,attr"      `
	CharacterID [16]byte `xml:"character_id,attr"`
	X           float64  `xml:"x,attr"           `
	Y           float64  `xml:"y,attr"           `
	Z           float64  `xml:"z,attr"           `
}

const (
	movementSize = int(unsafe.Sizeof(Movement{}))
)

func BenchmarkJSON(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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

func BenchmarkXML(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		xml.Marshal(mv)
	}
}

func BenchmarkXMLAlt(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := MovementAlt{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		xml.Marshal(mv)
	}
}

func BenchmarkGob(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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

func BenchmarkMsgpack(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		msgpack.Marshal(mv)
	}
}

func BenchmarkBSON(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		bson.MarshalValue(mv)
	}
}

func BenchmarkCBORCanonicalOptions(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		cbor.Marshal(mv,
			cbor.CanonicalEncOptions())
	}
}

func BenchmarkCBORCTAP2Options(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		cbor.Marshal(mv,
			cbor.CTAP2EncOptions())
	}
}

func BenchmarkCBORCoreDetOptions(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		cbor.Marshal(mv,
			cbor.CoreDetEncOptions())
	}
}

func BenchmarkCBORPreferredUnsortedOptions(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	for i := 0; i < b.N; i++ {
		cbor.Marshal(mv,
			cbor.PreferredUnsortedEncOptions())
	}
}

func BenchmarkBinary(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
