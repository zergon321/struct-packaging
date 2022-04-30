package fb

import flatbuffers "github.com/google/flatbuffers/go"

func CreateMovement(builder *flatbuffers.Builder, opcode int32, characterID []byte, x, y, z float64) []byte {
	builder.Reset()

	offset := builder.CreateByteVector(characterID)

	MovementStart(builder)
	MovementAddOpcode(builder, opcode)
	MovementAddCharacterId(builder, offset)
	MovementAddX(builder, x)
	MovementAddY(builder, y)
	MovementAddZ(builder, z)

	movPos := MovementEnd(builder)
	builder.Finish(movPos)

	return builder.Bytes[builder.Head():]
}

func ReadMovement(buf []byte) Movement {
	return *GetRootAsMovement(buf, 0)
}
