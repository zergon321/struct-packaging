package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"struct-packaging/fb"
	"struct-packaging/pb"
	"unsafe"

	"github.com/fxamacker/cbor"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
)

type Movement struct {
	Opcode      int32    `json:"opcode"       yaml:"opcode"       xml:"opcode"       cbor:"opcode"       msgpack:"opcode"       bson:"opcode"      `
	CharacterID [16]byte `json:"character_id" yaml:"character_id" xml:"character_id" cbor:"character_id" msgpack:"character_id" bson:"character_id"`
	X           float64  `json:"x"            yaml:"x"            xml:"x"            cbor:"x"            msgpack:"x"            bson:"x"           `
	Y           float64  `json:"y"            yaml:"y"            xml:"y"            cbor:"y"            msgpack:"y"            bson:"y"           `
	Z           float64  `json:"z"            yaml:"z"            xml:"z"            cbor:"z"            msgpack:"z"            bson:"z"           `
}

type MovementByte struct {
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

func main() {
	characterID, _ := uuid.Parse("1d9ce1d6-7ec5-48d3-be1d-ffaa0056921c")
	mv := Movement{
		Opcode:      32,
		CharacterID: [16]byte(characterID),
		X:           13.34,
		Y:           20.36,
		Z:           45.13,
	}
	mvAlt := MovementAlt{
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

	yamlData, err := json.Marshal(mv)
	handleError(err)

	xmlData, err := xml.Marshal(mv)
	handleError(err)

	xmlAltData, err := xml.Marshal(mvAlt)
	handleError(err)

	buffer := bytes.NewBuffer(make([]byte, 0, movementSize))
	buffer.Reset()
	enc := gob.NewEncoder(buffer)
	err = enc.Encode(mv)
	handleError(err)
	gobData := buffer.Bytes()

	msgpackData, err := msgpack.Marshal(mv)
	handleError(err)

	_, bsonData, err := bson.MarshalValue(mv)
	handleError(err)

	cborCanonicalData, err := cbor.Marshal(
		mv, cbor.CanonicalEncOptions())
	handleError(err)
	cborCTAP2Data, err := cbor.Marshal(
		mv, cbor.CTAP2EncOptions())
	handleError(err)
	cborCoreDetData, err := cbor.Marshal(
		mv, cbor.CoreDetEncOptions())
	handleError(err)
	cborPreferredUnsortedData, err := cbor.Marshal(
		mv, cbor.PreferredUnsortedEncOptions())
	handleError(err)

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
	fmt.Println("YAML data bytes length:", len(yamlData))
	fmt.Println("XML data bytes length:", len(xmlData))
	fmt.Println("XML tags only data bytes length:", len(xmlAltData))
	fmt.Println("gob data bytes length:", len(gobData))
	fmt.Println("Msgpack data bytes length:", len(msgpackData))
	fmt.Println("BSON data bytes length:", len(bsonData))
	fmt.Println("CBOR canonical options data bytes length:", len(cborCanonicalData))
	fmt.Println("CBOR CTAP2 options data bytes length:", len(cborCTAP2Data))
	fmt.Println("CBOR core det options data bytes length:", len(cborCoreDetData))
	fmt.Println("CBOR preferred unsorted options data bytes length:", len(cborPreferredUnsortedData))
	fmt.Println("binary data bytes length:", len(binData))
	fmt.Println("Protobuf data bytes length:", len(pbData))
	fmt.Println("flat buffers data bytes length:", len(fbData))
	fmt.Println("unsafe cast data bytes length:", len(unsafeData))

	var newMv MovementByte

	err = cbor.Unmarshal(cborCanonicalData, &newMv)
	handleError(err)
	fmt.Println(newMv)

	err = cbor.Unmarshal(cborCTAP2Data, &newMv)
	handleError(err)
	fmt.Println(newMv)

	err = cbor.Unmarshal(cborCoreDetData, &newMv)
	handleError(err)
	fmt.Println(newMv)

	err = cbor.Unmarshal(cborPreferredUnsortedData, &newMv)
	handleError(err)
	fmt.Println(newMv)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
