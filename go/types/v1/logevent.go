package typesv1

import (
	"time"

	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func KeyVal(k string, v *Val) *KV {
	return &KV{Key: k, Value: v}
}

func ValStr(v string) *Val {
	return &Val{Kind: &Val_Str{Str: v}}
}

func ValF64(v float64) *Val {
	return &Val{Kind: &Val_F64{F64: v}}
}

func ValI64(v int64) *Val {
	return &Val{Kind: &Val_I64{I64: v}}
}

func ValBool(v bool) *Val {
	return &Val{Kind: &Val_Bool{Bool: v}}
}

func ValArr(v []*Val) *Val {
	return &Val{Kind: &Val_Arr{Arr: &Arr{
		Items: v,
	}}}
}

func ValObj(v ...*KV) *Val {
	return &Val{Kind: &Val_Obj{Obj: &Obj{Kvs: v}}}
}

func ValTime(v time.Time) *Val {
	return &Val{Kind: &Val_Ts{Ts: timestamppb.New(v)}}
}

func ValDuration(v time.Duration) *Val {
	return &Val{Kind: &Val_Dur{Dur: durationpb.New(v)}}
}
