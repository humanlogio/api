package typesv1

import (
	"fmt"
	"slices"

	"github.com/cespare/xxhash"
)

const (
	tagNull       = uint8(0x0)
	tagBool       = uint8(0x1)
	tagI64        = uint8(0x2)
	tagF64        = uint8(0x3)
	tagTs         = uint8(0x4)
	tagDur        = uint8(0x5)
	tagStr        = uint8(0x6)
	tagBlob       = uint8(0x7)
	tagArray      = uint8(0x8)
	tagMapEntries = uint8(0x9)
	tagKVs        = uint8(0xA)
	tagHash64     = uint8(0xB)
	tagTraceID    = uint8(0xC)
	tagSpanID     = uint8(0xD)
	tagULID       = uint8(0xE)
)

var (
	nullh          = xxhash.Sum64([]byte{tagNull, 0x0})
	trueh          = xxhash.Sum64([]byte{tagBool, 0x1})
	falseh         = xxhash.Sum64([]byte{tagBool, 0x0})
	tagArrayh      = xxhash.Sum64([]byte{tagArray})
	tagMapEntriesh = xxhash.Sum64([]byte{tagMapEntries})
	tagKVsh        = xxhash.Sum64([]byte{tagKVs})
)

func Hash64Value(val *Val) uint64 {
	switch vv := val.Kind.(type) {
	case *Val_Str:
		return Hash64String(vv.Str)
	case *Val_TraceId:
		return Hash64TraceID(vv.TraceId)
	case *Val_SpanId:
		return Hash64SpanID(vv.SpanId)
	case *Val_Ulid:
		return Hash64ULID(vv.Ulid)
	case *Val_Blob:
		return Hash64Blob(vv.Blob)
	case *Val_Null:
		return Hash64Null()
	case *Val_Bool:
		return Hash64Bool(vv.Bool)
	case *Val_I64:
		return Hash64I64(vv.I64)
	case *Val_F64:
		return Hash64F64(vv.F64)
	case *Val_Hash64:
		return Hash64Hash64(vv.Hash64)
	case *Val_Ts:
		return Hash64Ts(vv.Ts.Seconds, vv.Ts.Nanos)
	case *Val_Dur:
		return Hash64Dur(vv.Dur.Seconds, vv.Dur.Nanos)
	case *Val_Arr:
		return Hash64Values_orderDoesntMatter(vv.Arr.Items)
	case *Val_Obj:
		return Hash64KeyValues_orderDoesntMatter(vv.Obj.Kvs)
	case *Val_Map:
		return Hash64MapEntries_orderDoesntMatter(vv.Map.Entries)
	default:
		panic(fmt.Sprintf("missing case: %T (%#v)", vv, vv))
	}
}

func Hash64KeyValues_orderDoesntMatter(kvs []*KV) uint64 {
	xored := tagKVsh
	for _, kv := range kvs {
		h := xxhash.Sum64String(kv.Key) ^ Hash64Value(kv.Value)
		xored ^= h
	}
	return xored
}

func Hash64MapEntries_orderDoesntMatter(kvs []*Map_Entry) uint64 {
	xored := tagMapEntriesh
	for _, kv := range kvs {
		h := Hash64Value(kv.Key) ^ Hash64Value(kv.Value)
		xored ^= h
	}
	return xored
}

func Hash64Values_orderDoesntMatter(arr []*Val) uint64 {
	xored := tagArrayh
	for _, el := range arr {
		h := Hash64Value(el)
		xored ^= h
	}
	return xored
}

func Hash64String(v string) uint64 {
	fp := slices.Concat([]byte{tagStr}, []byte(v))
	return xxhash.Sum64(fp)
}

func Hash64TraceID(v *TraceID) uint64 {
	fp := []byte{
		tagTraceID,
		byte(v.High >> 52),
		byte(v.High >> 48),
		byte(v.High >> 32),
		byte(v.High >> 24),
		byte(v.High >> 16),
		byte(v.High >> 8),
		byte(v.High >> 0),
		byte(v.Low >> 52),
		byte(v.Low >> 48),
		byte(v.Low >> 32),
		byte(v.Low >> 24),
		byte(v.Low >> 16),
		byte(v.Low >> 8),
		byte(v.Low >> 0),
	}
	return xxhash.Sum64(fp)
}

func Hash64SpanID(v *SpanID) uint64 {
	fp := []byte{
		tagSpanID,
		byte(v.Id >> 52),
		byte(v.Id >> 48),
		byte(v.Id >> 32),
		byte(v.Id >> 24),
		byte(v.Id >> 16),
		byte(v.Id >> 8),
		byte(v.Id >> 0),
	}
	return xxhash.Sum64(fp)
}

func Hash64ULID(v *ULID) uint64 {
	fp := []byte{
		tagULID,
		byte(v.High >> 52),
		byte(v.High >> 48),
		byte(v.High >> 32),
		byte(v.High >> 24),
		byte(v.High >> 16),
		byte(v.High >> 8),
		byte(v.High >> 0),
		byte(v.Low >> 52),
		byte(v.Low >> 48),
		byte(v.Low >> 32),
		byte(v.Low >> 24),
		byte(v.Low >> 16),
		byte(v.Low >> 8),
		byte(v.Low >> 0),
	}
	return xxhash.Sum64(fp)
}

func Hash64Blob(v []byte) uint64 {
	fp := slices.Concat([]byte{tagBlob}, v)
	return xxhash.Sum64(fp)
}

func Hash64Null() uint64 {
	return nullh
}

func Hash64Bool(v bool) uint64 {
	if v {
		return trueh
	}
	return falseh
}

func Hash64I64(v int64) uint64 {
	fp := []byte{
		tagI64,
		byte(v >> 52),
		byte(v >> 48),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v >> 0),
	}
	return xxhash.Sum64(fp)
}

func Hash64F64(v float64) uint64 {
	fp := []byte{
		tagF64,
		byte(int64(v) >> 52),
		byte(int64(v) >> 48),
		byte(int64(v) >> 32),
		byte(int64(v) >> 24),
		byte(int64(v) >> 16),
		byte(int64(v) >> 8),
		byte(int64(v) >> 0),
	}
	return xxhash.Sum64(fp)
}

func Hash64Hash64(v uint64) uint64 {
	fp := []byte{
		tagHash64,
		byte(v >> 52),
		byte(v >> 48),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v >> 0),
	}
	return xxhash.Sum64(fp)
}

func Hash64Ts(seconds int64, nanos int32) uint64 {
	fp := []byte{
		tagTs,
		byte(nanos >> 24),
		byte(nanos >> 16),
		byte(nanos >> 8),
		byte(seconds >> 52),
		byte(seconds >> 48),
		byte(seconds >> 32),
		byte(seconds >> 24),
		byte(seconds >> 16),
		byte(seconds >> 8),
		byte(seconds >> 0),
	}
	return xxhash.Sum64(fp)
}

func Hash64Dur(seconds int64, nanos int32) uint64 {
	fp := []byte{
		tagDur,
		byte(nanos >> 24),
		byte(nanos >> 16),
		byte(nanos >> 8),
		byte(seconds >> 52),
		byte(seconds >> 48),
		byte(seconds >> 32),
		byte(seconds >> 24),
		byte(seconds >> 16),
		byte(seconds >> 8),
		byte(seconds >> 0),
	}
	return xxhash.Sum64(fp)
}
