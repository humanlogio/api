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
		return fpString(vv.Str)
	case *Val_Blob:
		return fpBlob(vv.Blob)
	case *Val_Null:
		return fpNull()
	case *Val_Bool:
		return fpBool(vv.Bool)
	case *Val_I64:
		return fpI64(vv.I64)
	case *Val_F64:
		return fpF64(vv.F64)
	case *Val_Ts:
		return fpTs(vv.Ts.Seconds, vv.Ts.Nanos)
	case *Val_Dur:
		return fpDur(vv.Dur.Seconds, vv.Dur.Nanos)
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

func fpString(v string) uint64 {
	fp := slices.Concat([]byte{tagStr}, []byte(v))
	return xxhash.Sum64(fp)
}

func fpBlob(v []byte) uint64 {
	fp := slices.Concat([]byte{tagBlob}, v)
	return xxhash.Sum64(fp)
}

func fpNull() uint64 {
	return nullh
}

func fpBool(v bool) uint64 {
	if v {
		return trueh
	}
	return falseh
}

func fpI64(v int64) uint64 {
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

func fpF64(v float64) uint64 {
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

func fpTs(seconds int64, nanos int32) uint64 {
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

func fpDur(seconds int64, nanos int32) uint64 {
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
