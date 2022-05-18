package main_dec_test

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"struct-packaging/fb"
	"struct-packaging/pb"
	"testing"
	"unsafe"

	"github.com/fxamacker/cbor"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v2"
)

type BytesReadSeeker struct {
	data []byte
	pos  int
}

func (brs *BytesReadSeeker) Seek(offset int64, whence int) (int64, error) {
	var start int

	switch whence {
	case io.SeekStart:
		start = 0

	case io.SeekCurrent:
		start = brs.pos

	case io.SeekEnd:
		start = len(brs.data)

	default:
		return -1, fmt.Errorf("option not defined")
	}

	newPos := start + int(offset)

	switch {
	case newPos < 0:
		newPos = 0

	case newPos > len(brs.data):
		newPos = len(brs.data)
	}

	brs.pos = newPos

	return int64(brs.pos), nil
}

func (brw *BytesReadSeeker) Write(p []byte) (n int, err error) {
	offset := len(p)
	brw.data = append(brw.data, p...)
	brw.pos += offset

	return offset, nil
}

func (brs *BytesReadSeeker) Read(p []byte) (n int, err error) {
	if brs.pos == len(brs.data) {
		return -1, io.EOF
	}

	offset := len(p)

	if offset > len(brs.data)-brs.pos {
		offset = len(brs.data) - brs.pos
	}

	copy(p, brs.data[brs.pos:brs.pos+offset])
	brs.pos += offset

	return offset, nil
}

type Movement struct {
	Opcode      int32    `json:"opcode"       yaml:"opcode"       xml:"opcode"       cbor:"opcode"       msgpack:"opcode"       bson:"opcode"      `
	CharacterID [16]byte `json:"character_id" yaml:"character_id" xml:"character_id" cbor:"character_id" msgpack:"character_id" bson:"character_id"`
	X           float64  `json:"x"            yaml:"x"            xml:"x"            cbor:"x"            msgpack:"x"            bson:"x"           `
	Y           float64  `json:"y"            yaml:"y"            xml:"y"            cbor:"y"            msgpack:"y"            bson:"y"           `
	Z           float64  `json:"z"            yaml:"z"            xml:"z"            cbor:"z"            msgpack:"z"            bson:"z"           `
}

type MovementSlice struct {
	Opcode      int32   `json:"opcode"       yaml:"opcode"       xml:"opcode"       cbor:"opcode"       msgpack:"opcode"       bson:"opcode"      `
	CharacterID []byte  `json:"character_id" yaml:"character_id" xml:"character_id" cbor:"character_id" msgpack:"character_id" bson:"character_id"`
	X           float64 `json:"x"            yaml:"x"            xml:"x"            cbor:"x"            msgpack:"x"            bson:"x"           `
	Y           float64 `json:"y"            yaml:"y"            xml:"y"            cbor:"y"            msgpack:"y"            bson:"y"           `
	Z           float64 `json:"z"            yaml:"z"            xml:"z"            cbor:"z"            msgpack:"z"            bson:"z"           `
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

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func BenchmarkJSON(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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

func BenchmarkYAML(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := yaml.Marshal(mv)
	var newMv Movement

	for i := 0; i < b.N; i++ {
		yaml.Unmarshal(data, &newMv)
	}
}

func BenchmarkYAMLStrict(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	data, _ := yaml.Marshal(mv)
	var newMv Movement

	for i := 0; i < b.N; i++ {
		yaml.UnmarshalStrict(data, &newMv)
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
	data, _ := xml.Marshal(mv)
	var newMv Movement

	for i := 0; i < b.N; i++ {
		xml.Unmarshal(data, &newMv)
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
	data, _ := xml.Marshal(mv)
	var newMv MovementAlt

	for i := 0; i < b.N; i++ {
		xml.Unmarshal(data, &newMv)
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

	buffer := bytes.NewBuffer(make([]byte, 0, movementSize))
	enc := gob.NewEncoder(buffer)

	var newMv Movement
	dec := gob.NewDecoder(buffer)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		enc.Encode(mv)
		b.StartTimer()

		dec.Decode(&newMv)
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
	data, _ := msgpack.Marshal(mv)
	var newMv Movement

	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(data, &newMv)
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
	_, data, _ := bson.MarshalValue(mv)
	var newMv Movement

	for i := 0; i < b.N; i++ {
		bson.Unmarshal(data, &newMv)
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
	data, _ := cbor.Marshal(mv,
		cbor.CanonicalEncOptions())
	var newMv MovementSlice

	for i := 0; i < b.N; i++ {
		cbor.Unmarshal(data, &newMv)
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
	data, _ := cbor.Marshal(mv,
		cbor.CTAP2EncOptions())
	var newMv MovementSlice

	for i := 0; i < b.N; i++ {
		cbor.Unmarshal(data, &newMv)
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
	data, _ := cbor.Marshal(mv,
		cbor.CoreDetEncOptions())
	var newMv MovementSlice

	for i := 0; i < b.N; i++ {
		cbor.Unmarshal(data, &newMv)
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
	data, _ := cbor.Marshal(mv,
		cbor.PreferredUnsortedEncOptions())
	var newMv MovementSlice

	for i := 0; i < b.N; i++ {
		cbor.Unmarshal(data, &newMv)
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
	buffer := &BytesReadSeeker{
		data: data,
		pos:  0,
	}

	binary.Write(buffer, binary.LittleEndian, mv.Opcode)
	binary.Write(buffer, binary.LittleEndian, mv.CharacterID[:])
	binary.Write(buffer, binary.LittleEndian, mv.X)
	binary.Write(buffer, binary.LittleEndian, mv.Y)
	binary.Write(buffer, binary.LittleEndian, mv.Z)

	var newMv Movement

	for i := 0; i < b.N; i++ {
		buffer.Seek(0, io.SeekStart)

		binary.Read(buffer, binary.LittleEndian, &newMv.Opcode)
		binary.Read(buffer, binary.LittleEndian, newMv.CharacterID[:])
		binary.Read(buffer, binary.LittleEndian, &mv.X)
		binary.Read(buffer, binary.LittleEndian, &mv.Y)
		binary.Read(buffer, binary.LittleEndian, &mv.Z)
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
	buffer := &BytesReadSeeker{
		data: data,
		pos:  0,
	}

	binary.Write(buffer, binary.BigEndian, mv.Opcode)
	binary.Write(buffer, binary.BigEndian, mv.CharacterID[:])
	binary.Write(buffer, binary.BigEndian, mv.X)
	binary.Write(buffer, binary.BigEndian, mv.Y)
	binary.Write(buffer, binary.BigEndian, mv.Z)

	var newMv Movement

	for i := 0; i < b.N; i++ {
		buffer.Seek(0, io.SeekStart)

		binary.Read(buffer, binary.BigEndian, &newMv.Opcode)
		binary.Read(buffer, binary.BigEndian, newMv.CharacterID[:])
		binary.Read(buffer, binary.BigEndian, &mv.X)
		binary.Read(buffer, binary.BigEndian, &mv.Y)
		binary.Read(buffer, binary.BigEndian, &mv.Z)
	}
}

func BenchmarkBinaryWholeStruct(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movementSize)
	buffer := &BytesReadSeeker{
		data: data,
		pos:  0,
	}
	binary.Write(buffer, binary.LittleEndian, mv)
	var newMv Movement

	for i := 0; i < b.N; i++ {
		buffer.Seek(0, io.SeekStart)
		binary.Read(buffer, binary.LittleEndian, &newMv)
	}
}

func BenchmarkBinaryWholeStructBigEndian(b *testing.B) {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	data := make([]byte, 0, movementSize)
	buffer := &BytesReadSeeker{
		data: data,
		pos:  0,
	}
	binary.Write(buffer, binary.BigEndian, mv)
	var newMv Movement

	for i := 0; i < b.N; i++ {
		buffer.Seek(0, io.SeekStart)
		binary.Read(buffer, binary.BigEndian, &newMv)
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
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
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}

	builder := flatbuffers.NewBuilder(movementSize)
	data := fb.CreateMovement(builder, mv.Opcode, mv.CharacterID[:], mv.X, mv.Y, mv.Z)

	for i := 0; i < b.N; i++ {
		fb.ReadMovement(data)
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

	data := (*[movementSize]byte)(unsafe.Pointer(&mv))[:]

	for i := 0; i < b.N; i++ {
		_ = *(*Movement)(unsafe.Pointer(&data[0]))
	}
}
