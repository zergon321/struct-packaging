package main_dec_test

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"io"
	"math"
	"math/rand"
	"struct-packaging/fb"
	"struct-packaging/pb"
	"testing"
	"unsafe"

	"github.com/fxamacker/cbor"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/vmihailenco/msgpack"
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
	Movement32Size = int(unsafe.Sizeof(Movement32{}))
)

func Benchmark32JSON(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data, _ := json.Marshal(mv)
	var newMv Movement32

	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &newMv)
	}
}

func Benchmark32YAML(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := yaml.Marshal(mv)
	var newMv Movement32

	for i := 0; i < b.N; i++ {
		yaml.Unmarshal(data, &newMv)
	}
}

func Benchmark32YAMLStrict(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := yaml.Marshal(mv)
	var newMv Movement32

	for i := 0; i < b.N; i++ {
		yaml.UnmarshalStrict(data, &newMv)
	}
}

func Benchmark32XML(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := xml.Marshal(mv)
	var newMv Movement32

	for i := 0; i < b.N; i++ {
		xml.Unmarshal(data, &newMv)
	}
}

func Benchmark32XMLAlt(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32Alt{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := xml.Marshal(mv)
	var newMv Movement32Alt

	for i := 0; i < b.N; i++ {
		xml.Unmarshal(data, &newMv)
	}
}

func Benchmark32Gob(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	buffer := bytes.NewBuffer(make([]byte, 0, Movement32Size))
	enc := gob.NewEncoder(buffer)

	var newMv Movement32
	dec := gob.NewDecoder(buffer)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		enc.Encode(mv)
		b.StartTimer()

		dec.Decode(&newMv)
	}
}

func Benchmark32Msgpack(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := msgpack.Marshal(mv)
	var newMv Movement32

	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(data, &newMv)
	}
}

func Benchmark32BSON(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	_, data, _ := bson.MarshalValue(mv)
	var newMv Movement32

	for i := 0; i < b.N; i++ {
		bson.Unmarshal(data, &newMv)
	}
}

func Benchmark32CBORCanonicalOptions(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := cbor.Marshal(mv,
		cbor.CanonicalEncOptions())
	var newMv MovementSlice

	for i := 0; i < b.N; i++ {
		cbor.Unmarshal(data, &newMv)
	}
}

func Benchmark32CBORCTAP2Options(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := cbor.Marshal(mv,
		cbor.CTAP2EncOptions())
	var newMv MovementSlice

	for i := 0; i < b.N; i++ {
		cbor.Unmarshal(data, &newMv)
	}
}

func Benchmark32CBORCoreDetOptions(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := cbor.Marshal(mv,
		cbor.CoreDetEncOptions())
	var newMv MovementSlice

	for i := 0; i < b.N; i++ {
		cbor.Unmarshal(data, &newMv)
	}
}

func Benchmark32CBORPreferredUnsortedOptions(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := cbor.Marshal(mv,
		cbor.PreferredUnsortedEncOptions())
	var newMv MovementSlice

	for i := 0; i < b.N; i++ {
		cbor.Unmarshal(data, &newMv)
	}
}

func Benchmark32Binary(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, Movement32Size)
	buffer := &BytesReadWriteSeeker{
		data: data,
		pos:  0,
	}

	binary.Write(buffer, binary.LittleEndian, mv.Opcode)
	binary.Write(buffer, binary.LittleEndian, mv.CharacterID[:])
	binary.Write(buffer, binary.LittleEndian, mv.X)
	binary.Write(buffer, binary.LittleEndian, mv.Y)
	binary.Write(buffer, binary.LittleEndian, mv.Z)

	var newMv Movement32

	for i := 0; i < b.N; i++ {
		buffer.Seek(0, io.SeekStart)

		binary.Read(buffer, binary.LittleEndian, &newMv.Opcode)
		binary.Read(buffer, binary.LittleEndian, newMv.CharacterID[:])
		binary.Read(buffer, binary.LittleEndian, &mv.X)
		binary.Read(buffer, binary.LittleEndian, &mv.Y)
		binary.Read(buffer, binary.LittleEndian, &mv.Z)
	}
}

func Benchmark32BinaryBigEndian(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, Movement32Size)
	buffer := &BytesReadWriteSeeker{
		data: data,
		pos:  0,
	}

	binary.Write(buffer, binary.BigEndian, mv.Opcode)
	binary.Write(buffer, binary.BigEndian, mv.CharacterID[:])
	binary.Write(buffer, binary.BigEndian, mv.X)
	binary.Write(buffer, binary.BigEndian, mv.Y)
	binary.Write(buffer, binary.BigEndian, mv.Z)

	var newMv Movement32

	for i := 0; i < b.N; i++ {
		buffer.Seek(0, io.SeekStart)

		binary.Read(buffer, binary.BigEndian, &newMv.Opcode)
		binary.Read(buffer, binary.BigEndian, newMv.CharacterID[:])
		binary.Read(buffer, binary.BigEndian, &mv.X)
		binary.Read(buffer, binary.BigEndian, &mv.Y)
		binary.Read(buffer, binary.BigEndian, &mv.Z)
	}
}

func Benchmark32BinaryWholeStruct(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, Movement32Size)
	buffer := &BytesReadWriteSeeker{
		data: data,
		pos:  0,
	}
	binary.Write(buffer, binary.LittleEndian, mv)
	var newMv Movement32

	for i := 0; i < b.N; i++ {
		buffer.Seek(0, io.SeekStart)
		binary.Read(buffer, binary.LittleEndian, &newMv)
	}
}

func Benchmark32BinaryWholeStructBigEndian(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, Movement32Size)
	buffer := &BytesReadWriteSeeker{
		data: data,
		pos:  0,
	}
	binary.Write(buffer, binary.BigEndian, mv)
	var newMv Movement32

	for i := 0; i < b.N; i++ {
		buffer.Seek(0, io.SeekStart)
		binary.Read(buffer, binary.BigEndian, &newMv)
	}
}

func Benchmark32BinaryNoReflection(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, Movement32Size)

	binary.LittleEndian.PutUint32(data, uint32(mv.Opcode))
	copy(data[4:20], mv.CharacterID[:])
	binary.LittleEndian.PutUint64(data[20:], math.Float64bits(mv.X))
	binary.LittleEndian.PutUint64(data[28:], math.Float64bits(mv.Y))
	binary.LittleEndian.PutUint64(data[36:], math.Float64bits(mv.Z))

	var newMv Movement32

	for i := 0; i < b.N; i++ {
		newMv.Opcode = int32(binary.LittleEndian.Uint32(data))
		copy(newMv.CharacterID[:], data[4:20])
		newMv.X = math.Float64frombits(binary.LittleEndian.Uint64(data[20:]))
		newMv.Y = math.Float64frombits(binary.LittleEndian.Uint64(data[28:]))
		newMv.Z = math.Float64frombits(binary.LittleEndian.Uint64(data[36:]))
	}
}

func Benchmark32BinaryBigEndianNoReflection(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, Movement32Size)

	binary.BigEndian.PutUint32(data, uint32(mv.Opcode))
	copy(data[4:20], mv.CharacterID[:])
	binary.BigEndian.PutUint64(data[20:], math.Float64bits(mv.X))
	binary.BigEndian.PutUint64(data[28:], math.Float64bits(mv.Y))
	binary.BigEndian.PutUint64(data[36:], math.Float64bits(mv.Z))

	var newMv Movement32

	for i := 0; i < b.N; i++ {
		newMv.Opcode = int32(binary.BigEndian.Uint32(data))
		copy(newMv.CharacterID[:], data[4:20])
		newMv.X = math.Float64frombits(binary.BigEndian.Uint64(data[20:]))
		newMv.Y = math.Float64frombits(binary.BigEndian.Uint64(data[28:]))
		newMv.Z = math.Float64frombits(binary.BigEndian.Uint64(data[36:]))
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

	data, _ := proto.Marshal(&mv)

	var newMv pb.Movement

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(data, &newMv)
	}
}

func Benchmark32FlatBuffers(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	builder := flatbuffers.NewBuilder(Movement32Size)
	data := fb.CreateMovement(builder, mv.Opcode, mv.CharacterID[:], mv.X, mv.Y, mv.Z)

	for i := 0; i < b.N; i++ {
		fb.ReadMovement(data)
	}
}

func Benchmark32Unsafe(b *testing.B) {
	rand.Seed(32)
	var randomData [32]byte
	rand.Read(randomData[:])
	characterID := randomData
	mv := Movement32{
		Opcode:      32,
		CharacterID: [32]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := (*[Movement32Size]byte)(unsafe.Pointer(&mv))[:]

	for i := 0; i < b.N; i++ {
		_ = *(*Movement32)(unsafe.Pointer(&data[0]))
	}
}
