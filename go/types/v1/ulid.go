package typesv1

import "encoding/binary"

func ULIDFromBytes(out *ULID, ulid [16]byte) {
	out.High = binary.BigEndian.Uint64(ulid[0:8])
	out.Low = binary.BigEndian.Uint64(ulid[8:16])
}

func ULIDToBytes(b []byte, in *ULID) [16]byte {
	b = binary.BigEndian.AppendUint64(b, in.High)
	b = binary.BigEndian.AppendUint64(b, in.Low)
	return *(*[16]byte)(b[:16])
}
