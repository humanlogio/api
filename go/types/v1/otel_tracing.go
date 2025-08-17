package typesv1

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func TraceIDFromBytesArray(out *TraceID, id [16]byte) *TraceID {
	if out == nil {
		out = new(TraceID)
	}
	out.High = binary.BigEndian.Uint64(id[0:8])
	out.Low = binary.BigEndian.Uint64(id[8:16])
	return out
}

func TraceIDFromBytesSlice(out *TraceID, id []byte) (*TraceID, error) {
	if len(id) < 16 {
		return nil, fmt.Errorf("trace_id has less than 128 bits: %d", len(id)*8)
	}
	if len(id) > 16 {
		return nil, fmt.Errorf("trace_id has more than 128 bits: %d", len(id)*8)
	}
	if out == nil {
		out = new(TraceID)
	}
	out.High = binary.BigEndian.Uint64(id[0:8])
	out.Low = binary.BigEndian.Uint64(id[8:16])
	return out, nil
}

func TraceIDFromHex(out *TraceID, id string) (*TraceID, error) {
	b, err := hex.DecodeString(id)
	if err != nil {
		return nil, err
	}
	return TraceIDFromBytesSlice(out, b)
}

func TraceIDToBytes(b []byte, in *TraceID) [16]byte {
	b = binary.BigEndian.AppendUint64(b, in.High)
	b = binary.BigEndian.AppendUint64(b, in.Low)
	return *(*[16]byte)(b[:16])
}

func SpanIDFromBytesArray(out *SpanID, id [8]byte) *SpanID {
	if out == nil {
		out = new(SpanID)
	}
	out.Id = binary.BigEndian.Uint64(id[0:8])
	return out
}

func SpanIDFromBytesSlice(out *SpanID, id []byte) (*SpanID, error) {
	if len(id) < 8 {
		return nil, fmt.Errorf("trace_id has less than 64 bits: %d", len(id)*8)
	}
	if len(id) > 8 {
		return nil, fmt.Errorf("trace_id has more than 64 bits: %d", len(id)*8)
	}
	if out == nil {
		out = new(SpanID)
	}
	out.Id = binary.BigEndian.Uint64(id[0:8])
	return out, nil
}

func SpanIDFromHex(out *SpanID, id string) (*SpanID, error) {
	b, err := hex.DecodeString(id)
	if err != nil {
		return nil, err
	}
	return SpanIDFromBytesSlice(out, b)
}

func SpanIDToBytes(b []byte, in *SpanID) [8]byte {
	b = binary.BigEndian.AppendUint64(b, in.Id)
	return *(*[8]byte)(b[:8])
}
