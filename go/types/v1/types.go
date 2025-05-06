package typesv1

import (
	"bytes"
	"time"

	"google.golang.org/protobuf/proto"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func KeyVal(k string, v *Val) *KV {
	return &KV{Key: k, Value: v}
}

func EqualTypes(a, b *VarType) bool {
	return proto.Equal(a, b)
}

func EqualVal(a, b *Val) bool {
	switch ak := a.Kind.(type) {
	default:
		return false
	case *Val_Str:
		return ak.Str == b.GetStr()
	case *Val_F64:
		return ak.F64 == b.GetF64()
	case *Val_I64:
		return ak.I64 == b.GetI64()
	case *Val_Bool:
		return ak.Bool == b.GetBool()
	case *Val_Ts:
		return ak.Ts.AsTime().Equal(b.GetTs().AsTime())
	case *Val_Dur:
		return ak.Dur == b.GetDur()
	case *Val_Blob:
		return bytes.Equal(ak.Blob, b.GetBlob())
	case *Val_Arr:
		if len(ak.Arr.Items) != len(b.GetArr().Items) {
			return false
		}
		for i, ael := range ak.Arr.Items {
			bel := b.GetArr().Items[i]
			if !EqualVal(ael, bel) {
				return false
			}
		}
		return true
	case *Val_Obj:
		if len(ak.Obj.Kvs) != len(b.GetObj().Kvs) {
			return false
		}
		// for simplicity, order matters
		// TODO: make it so the order doesn't matter
		for i, akv := range ak.Obj.Kvs {
			bkv := b.GetObj().Kvs[i]
			if akv.Key != bkv.Key {
				return false
			}
			if !EqualVal(akv.Value, bkv.Value) {
				return false
			}
		}
		return true
	case *Val_Map:
		if len(ak.Map.Entries) != len(b.GetMap().Entries) {
			return false
		}
		// for simplicity, order matters
		// TODO: make it so the order doesn't matter
		for i, akv := range ak.Map.Entries {
			bkv := b.GetMap().Entries[i]
			if !EqualVal(akv.Key, bkv.Key) {
				return false
			}
			if !EqualVal(akv.Value, bkv.Value) {
				return false
			}
		}
		return true
	case *Val_Null:
		_, ok := b.Kind.(*Val_Null)
		return ok
	}
}

func TypeUnknown() *VarType {
	return &VarType{Type: &VarType_Scalar{Scalar: ScalarType_unknown}}
}

func TypeStr() *VarType {
	return &VarType{Type: &VarType_Scalar{Scalar: ScalarType_str}}
}

func TypeF64() *VarType {
	return &VarType{Type: &VarType_Scalar{Scalar: ScalarType_f64}}
}

func TypeI64() *VarType {
	return &VarType{Type: &VarType_Scalar{Scalar: ScalarType_i64}}
}

func TypeBool() *VarType {
	return &VarType{Type: &VarType_Scalar{Scalar: ScalarType_bool}}
}

func TypeTimestamp() *VarType {
	return &VarType{Type: &VarType_Scalar{Scalar: ScalarType_ts}}
}

func TypeDuration() *VarType {
	return &VarType{Type: &VarType_Scalar{Scalar: ScalarType_dur}}
}

func TypeBlob() *VarType {
	return &VarType{Type: &VarType_Scalar{Scalar: ScalarType_blob}}
}

func TypeArr(v ...*VarType) *VarType {
	atyp := &VarType_ArrayType{Items: &VarType{Type: &VarType_Scalar{Scalar: ScalarType_unknown}}}
	for i, item := range v {
		if i == 0 {
			atyp.Items.Type = item.Type
		} else if item.Type != atyp.Items.Type {
			atyp.Items.Type = &VarType_Scalar{Scalar: ScalarType_unknown}
			break
		}
	}
	return &VarType{Type: &VarType_Array{Array: atyp}}
}

func TypeObj(types map[string]*VarType) *VarType {
	otyp := &VarType_ObjectType{Kvs: types}
	return &VarType{Type: &VarType_Object{Object: otyp}}
}

func TypeObjFromKVs(v ...*KV) *VarType {
	types := make(map[string]*VarType)
	for _, kv := range v {
		types[kv.Key] = kv.Value.Type
	}
	return TypeObj(types)
}

func TypeMap(k, v *VarType) *VarType {
	return &VarType{Type: &VarType_Map{Map: &VarType_MapType{Key: k, Value: v}}}
}

func TypeNull() *VarType {
	return &VarType{Type: &VarType_Null_{}}
}

func ValStr(v string) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_str}}
	return &Val{Type: typ, Kind: &Val_Str{Str: v}}
}

func ValF64(v float64) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_f64}}
	return &Val{Type: typ, Kind: &Val_F64{F64: v}}
}

func ValI64(v int64) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_i64}}
	return &Val{Type: typ, Kind: &Val_I64{I64: v}}
}

func ValBool(v bool) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_bool}}
	return &Val{Type: typ, Kind: &Val_Bool{Bool: v}}
}

func ValTime(v time.Time) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_ts}}
	return &Val{Type: typ, Kind: &Val_Ts{Ts: timestamppb.New(v)}}
}

func ValTimestamp(v *timestamppb.Timestamp) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_ts}}
	return &Val{Type: typ, Kind: &Val_Ts{Ts: v}}
}

func ValDuration(v time.Duration) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_dur}}
	return &Val{Type: typ, Kind: &Val_Dur{Dur: durationpb.New(v)}}
}

func ValDurationPB(v *durationpb.Duration) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_dur}}
	return &Val{Type: typ, Kind: &Val_Dur{Dur: v}}
}

func ValBlob(blob []byte) *Val {
	typ := &VarType{Type: &VarType_Scalar{Scalar: ScalarType_blob}}
	return &Val{Type: typ, Kind: &Val_Blob{Blob: blob}}
}

func ValArr(v ...*Val) *Val {
	typs := make([]*VarType, 0, len(v))
	for _, val := range v {
		typs = append(typs, val.Type)
	}
	return &Val{Type: TypeArr(typs...), Kind: &Val_Arr{Arr: &Arr{
		Items: v,
	}}}
}

func ValMap(kt, vt *VarType, v ...*Map_Entry) *Val {
	return &Val{Type: TypeMap(kt, vt), Kind: &Val_Map{Map: &Map{
		Entries: v,
	}}}
}

func ValObj(v ...*KV) *Val {
	return &Val{Type: TypeObjFromKVs(v...), Kind: &Val_Obj{Obj: &Obj{Kvs: v}}}
}

func ValNull() *Val {
	return &Val{Type: TypeNull(), Kind: &Val_Null{}}
}
