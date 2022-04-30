// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Movement struct {
	_tab flatbuffers.Table
}

func GetRootAsMovement(buf []byte, offset flatbuffers.UOffsetT) *Movement {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Movement{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMovement(buf []byte, offset flatbuffers.UOffsetT) *Movement {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Movement{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Movement) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Movement) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Movement) Opcode() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Movement) MutateOpcode(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *Movement) CharacterId(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Movement) CharacterIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Movement) CharacterIdBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Movement) MutateCharacterId(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Movement) X() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Movement) MutateX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *Movement) Y() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Movement) MutateY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func (rcv *Movement) Z() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Movement) MutateZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func MovementStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func MovementAddOpcode(builder *flatbuffers.Builder, opcode int32) {
	builder.PrependInt32Slot(0, opcode, 0)
}
func MovementAddCharacterId(builder *flatbuffers.Builder, characterId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(characterId), 0)
}
func MovementStartCharacterIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MovementAddX(builder *flatbuffers.Builder, x float64) {
	builder.PrependFloat64Slot(2, x, 0.0)
}
func MovementAddY(builder *flatbuffers.Builder, y float64) {
	builder.PrependFloat64Slot(3, y, 0.0)
}
func MovementAddZ(builder *flatbuffers.Builder, z float64) {
	builder.PrependFloat64Slot(4, z, 0.0)
}
func MovementEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
