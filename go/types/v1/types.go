package typesv1

import (
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

func ValArr(v ...*Val) *Val {
	typs := make([]*VarType, 0, len(v))
	for _, val := range v {
		typs = append(typs, val.Type)
	}
	return &Val{Type: TypeArr(typs...), Kind: &Val_Arr{Arr: &Arr{
		Items: v,
	}}}
}

func ValObj(v ...*KV) *Val {
	return &Val{Type: TypeObjFromKVs(v...), Kind: &Val_Obj{Obj: &Obj{Kvs: v}}}
}
