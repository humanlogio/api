package logql

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/humanlogio/api/go/types/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in   string
		want *v1.LogQuery
		err  error
	}{
		{`{session=1734666428038101000} msg="AuthenticateUser"`, q(nil, nil, qctx(nil, eq(id("session"), i64(1734666428038101000))), eq(id("msg"), str("AuthenticateUser"))), nil},
		{`{from=2006-01-02T15:04:05.999999999+07:00}`, q(ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil, nil), nil},
		{`{from=2006-01-02T15:04:05+07:00}`, q(ts("2006-01-02T15:04:05+07:00"), nil, nil, nil), nil},
		{`{to=2006-01-02T15:04:05.999999999+07:00}`, q(nil, ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil), nil},
		{`{to=2006-01-02T15:04:05+07:00}`, q(nil, ts("2006-01-02T15:04:05+07:00"), nil, nil), nil},
		{`{from=2006-01-02T15:04:07+07:00 to=2006-01-02T15:04:05+07:00}`, q(ts("2006-01-02T15:04:07+07:00"), ts("2006-01-02T15:04:05+07:00"), nil, nil), nil},
		{`{	from	=	2006-01-02T15:04:05.999999999+07:00}`, q(ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil, nil), nil},
		{`{	from	=	2006-01-02T15:04:05.999999999+07:00		}`, q(ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil, nil), nil},
		{`{	from	=	2006-01-02T15:04:05.999999999+07:00		to	=	2006-01-02T15:04:05.999999999+07:00		}`, q(ts("2006-01-02T15:04:05.999999999+07:00"), ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil), nil},
		{`{session=1}`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		{`{session = 1}`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		{`{session =1}`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		{`{session= 1}`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		{`{ session  = 1 }`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		{`{machine=1}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		{`{machine = 1}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		{`{machine =1}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		{`{machine= 1}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		{`{ machine  = 1 }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		{`{machine=1 session=2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{machine =1 session=2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{machine=1 session =2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{machine=1 session= 2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{ machine=1 session=2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{ machine =1 session=2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{ machine=1 session =2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{ machine=1 session= 2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{ machine=1 session=2   }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{ machine =1 session=2   }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{ machine=1 session =2   }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{ machine=1 session= 2   }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{machine=1	session=2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{machine	=1	session=2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{machine=1	session	=2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`{machine=1	session=	2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		{`lvl`, q(nil, nil, nil, id("lvl")), nil},
		{`msg`, q(nil, nil, nil, id("msg")), nil},
		{`kv`, q(nil, nil, nil, id("kv")), nil},
		{`"hello world"`, q(nil, nil, nil, str("hello world")), nil},
		{`true`, q(nil, nil, nil, boo(true)), nil},
		{`false`, q(nil, nil, nil, boo(false)), nil},
		{`.110`, q(nil, nil, nil, f64(.11)), nil},
		{`1111.0`, q(nil, nil, nil, f64(1111)), nil},
		{`10.3e9`, q(nil, nil, nil, f64(10.3e9)), nil},
		{`-0`, q(nil, nil, nil, neg(i64(0))), nil},
		{`-10000`, q(nil, nil, nil, neg(i64(10000))), nil},
		{`0`, q(nil, nil, nil, i64(0)), nil},
		{`1`, q(nil, nil, nil, i64(1)), nil},
		{`-1`, q(nil, nil, nil, neg(i64(1))), nil},
		{`""`, q(nil, nil, nil, str("")), nil},
		{`"hello world"`, q(nil, nil, nil, str("hello world")), nil},
		{`"hello world\n\t"`, q(nil, nil, nil, str("hello world\n\t")), nil},
		{`2006-01-02T15:04:05.999999999+07:00`, q(nil, nil, nil, tse("2006-01-02T15:04:05.999999999+07:00")), nil},
		{`ts:2006-01-02T15:04:05.999999999+07:00`, q(nil, nil, nil, tse("2006-01-02T15:04:05.999999999+07:00")), nil},
		{`{"hello":"world", "hello2":1}`, q(nil, nil, nil, obj(kv("hello", v1.ValStr("world")), kv("hello2", v1.ValI64(1)))), nil},
		{`{"hello":"world"}`, q(nil, nil, nil, obj(kv("hello", v1.ValStr("world")))), nil},
		{`{"hello":1}`, q(nil, nil, nil, obj(kv("hello", v1.ValI64(1)))), nil},
		{`{"hello":ts:2006-01-02T15:04:05.999999999+07:00}`, q(nil, nil, nil, obj(kv("hello", v1.ValTimestamp(ts("2006-01-02T15:04:05.999999999+07:00"))))), nil},
		{`{"hello":2006-01-02T15:04:05.999999999+07:00}`, q(nil, nil, nil, obj(kv("hello", v1.ValTimestamp(ts("2006-01-02T15:04:05.999999999+07:00"))))), nil},
		{`{}`, q(nil, nil, nil, obj()), nil},
		{`[]`, q(nil, nil, nil, arr()), nil},
		{`[1, "deux", 3.0]`, q(nil, nil, nil, arr(v1.ValI64(1), v1.ValStr("deux"), v1.ValF64(3.0))), nil},
		{`[[]]`, q(nil, nil, nil, arr(v1.ValArr())), nil},
		{`[[],[]]`, q(nil, nil, nil, arr(v1.ValArr(), v1.ValArr())), nil},
		{`[[1]]`, q(nil, nil, nil, arr(v1.ValArr(v1.ValI64(1)))), nil},
		{`[[],[1],[2,3]]`, q(nil, nil, nil, arr(v1.ValArr(), v1.ValArr(v1.ValI64(1)), v1.ValArr(v1.ValI64(2), v1.ValI64(3)))), nil},
		{`[{}]`, q(nil, nil, nil, arr(v1.ValObj())), nil},
		{`[{"hello":[]}]`, q(nil, nil, nil, arr(v1.ValObj(kv("hello", v1.ValArr())))), nil},
		{`[{"hello":[{},{}]}]`, q(nil, nil, nil, arr(v1.ValObj(kv("hello", v1.ValArr(v1.ValObj(), v1.ValObj()))))), nil},
		{`1000s`, q(nil, nil, nil, dure(1000*time.Second)), nil},
		{`1.5us`, q(nil, nil, nil, dure(1500*time.Nanosecond)), nil},
		{`1.5µs`, q(nil, nil, nil, dure(1500*time.Nanosecond)), nil},
		{`1.5ms`, q(nil, nil, nil, dure(1500*time.Microsecond)), nil},
		{`15ns`, q(nil, nil, nil, dure(15*time.Nanosecond)), nil},
		{`1.5s`, q(nil, nil, nil, dure(1500*time.Millisecond)), nil},
		{`-1.5s`, q(nil, nil, nil, neg(dure(1500*time.Millisecond))), nil},
		{`-1.5m`, q(nil, nil, nil, neg(dure(90*time.Second))), nil},
		{`-1.5h`, q(nil, nil, nil, neg(dure(90*time.Minute))), nil},
		{`now()`, q(nil, nil, nil, fn("now")), nil},
		{`now(1h)`, q(nil, nil, nil, fn("now", dure(1*time.Hour))), nil},
		{`fn(fn1(),fn2(fn3()))`, q(nil, nil, nil, fn("fn", fn("fn1"), fn("fn2", fn("fn3")))), nil},
		{`len(msg)`, q(nil, nil, nil, fn("len", id("msg"))), nil},
		{`!true`, q(nil, nil, nil, not(boo(true))), nil},
		{`! true`, q(nil, nil, nil, not(boo(true))), nil},
		{`!(eval())`, q(nil, nil, nil, not(fn("eval"))), nil},
		{`!( eval())`, q(nil, nil, nil, not(fn("eval"))), nil},
		{`1 + 2`, q(nil, nil, nil, add(i64(1), i64(2))), nil},
		{`1 + 2 + 3`, q(nil, nil, nil, add(add(i64(1), i64(2)), i64(3))), nil},
		{`1 - 2`, q(nil, nil, nil, sub(i64(1), i64(2))), nil},
		{`1 - 2 - 3`, q(nil, nil, nil, sub(sub(i64(1), i64(2)), i64(3))), nil},
		{`1 * 2`, q(nil, nil, nil, mul(i64(1), i64(2))), nil},
		{`-1 * 2`, q(nil, nil, nil, mul(neg(i64(1)), i64(2))), nil},
		{`1 / 2`, q(nil, nil, nil, div(i64(1), i64(2))), nil},
		{`1 + 2 * 3`, q(nil, nil, nil, add(i64(1), mul(i64(2), i64(3)))), nil},
		{`1 - 2 * 3`, q(nil, nil, nil, sub(i64(1), mul(i64(2), i64(3)))), nil},
		{`1 + 2 / 3`, q(nil, nil, nil, add(i64(1), div(i64(2), i64(3)))), nil},
		{`1 - 2 / 3`, q(nil, nil, nil, sub(i64(1), div(i64(2), i64(3)))), nil},
		{`1 = 2`, q(nil, nil, nil, eq(i64(1), i64(2))), nil},
		{`1 != 2`, q(nil, nil, nil, noteq(i64(1), i64(2))), nil},
		{`1 > 2`, q(nil, nil, nil, gt(i64(1), i64(2))), nil},
		{`1 >= 2`, q(nil, nil, nil, gte(i64(1), i64(2))), nil},
		{`1 < 2`, q(nil, nil, nil, lt(i64(1), i64(2))), nil},
		{`1 <= 2`, q(nil, nil, nil, lte(i64(1), i64(2))), nil},
		{`true && false`, q(nil, nil, nil, and(boo(true), boo(false))), nil},
		{`true || false`, q(nil, nil, nil, or(boo(true), boo(false))), nil},
		{`true && 1 > 2`, q(nil, nil, nil, and(boo(true), gt(i64(1), i64(2)))), nil},
		{`1 > 2 && true`, q(nil, nil, nil, and(gt(i64(1), i64(2)), boo(true))), nil},
		{`!true || false`, q(nil, nil, nil, or(not(boo(true)), boo(false))), nil},
		{`!(true || false)`, q(nil, nil, nil, not(or(boo(true), boo(false)))), nil},
		{`true && !false`, q(nil, nil, nil, and(boo(true), not(boo(false)))), nil},
		{`!(true && false)`, q(nil, nil, nil, not(and(boo(true), boo(false)))), nil},
		{`1 in [1,2,3]`, q(nil, nil, nil, setin(i64(1), arr(v1.ValI64(1), v1.ValI64(2), v1.ValI64(3)))), nil},
		{`1 not in [1,2,3]`, q(nil, nil, nil, setnotin(i64(1), arr(v1.ValI64(1), v1.ValI64(2), v1.ValI64(3)))), nil},
		{`true || false || 1`, q(nil, nil, nil, or(or(boo(true), boo(false)), i64(1))), nil},
		{`len(msg) > 0 | true`, q(nil, nil, nil, pipe(gt(fn("len", id("msg")), i64(0)), boo(true))), nil},
		{`len(msg) > 0 | true | lvl in [1,2,3]`, q(nil, nil, nil, pipe(pipe(gt(fn("len", id("msg")), i64(0)), boo(true)), setin(id("lvl"), arr(v1.ValI64(1), v1.ValI64(2), v1.ValI64(3))))), nil},
		{`kv["hello"]["world"][0]`, q(nil, nil, nil, idx(idx(idx(id("kv"), str("hello")), str("world")), i64(0))), nil},
		{` (  kv ["hello"] ["world"] [0]  ) `, q(nil, nil, nil, idx(idx(idx(id("kv"), str("hello")), str("world")), i64(0))), nil},
		{`msg="hello world" | len(lvl) > 0`, q(nil, nil, nil, pipe(eq(id("msg"), str("hello world")), gt(fn("len", id("lvl")), i64(0)))), nil},
		{`kv.hello.world`, q(nil, nil, nil, selector(selector(id("kv"), "hello"), "world")), nil},
		{`lvl.k.v | len(lvl)`, q(nil, nil, nil, pipe(selector(selector(id("lvl"), "k"), "v"), fn("len", id("lvl")))), nil},
		// not supported yet (not hard, just a lot of noisy work)
		// {`{"hello":now()}`, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Parse(tt.in)
			t.Logf("err=%v", err)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, pjson(tt.want), pjson(got))
				diff := cmp.Diff(pjson(tt.want), pjson(got), protocmp.Transform())
				if len(diff) > 0 {
					require.Empty(t, diff)
				}
			}
		})
	}
}

func pjson(v proto.Message) string {
	r, err := protojson.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(r)
}

func ts(str string) *timestamppb.Timestamp {
	t, err := time.Parse(time.RFC3339Nano, str)
	if err != nil {
		panic(err)
	}
	return timestamppb.New(t)
}

func tse(str string) *v1.Expr {
	return v1.ExprLiteral(v1.ValTimestamp(ts(str)))
}

func dure(v time.Duration) *v1.Expr {
	return v1.ExprLiteral(v1.ValDuration(v))
}

func q(f, t *timestamppb.Timestamp, ctx *v1.Context, e *v1.Expr) *v1.LogQuery {
	return &v1.LogQuery{From: f, To: t, Context: ctx, Query: e}
}

func qctx(m, s *v1.Expr) *v1.Context {
	return &v1.Context{MachineId: m, SessionId: s}
}

func neg(e *v1.Expr) *v1.Expr {
	return v1.ExprUnary(v1.UnaryOp_NEG, e)
}

func not(e *v1.Expr) *v1.Expr {
	return v1.ExprUnary(v1.UnaryOp_NOT, e)
}

func add(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_NUM_ADD, rhs)
}

func sub(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_NUM_SUB, rhs)
}

func mul(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_NUM_MUL, rhs)
}

func div(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_NUM_DIV, rhs)
}

func eq(l, r *v1.Expr) *v1.Expr {
	return v1.ExprBinary(l, v1.BinaryOp_CMP_EQ, r)
}

func noteq(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_CMP_NOTEQ, rhs)
}

func gt(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_CMP_GT, rhs)
}

func gte(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_CMP_GTE, rhs)
}

func lt(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_CMP_LT, rhs)
}

func lte(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_CMP_LTE, rhs)
}

func and(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_LOG_AND, rhs)
}

func or(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_LOG_OR, rhs)
}

func setin(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_SET_IN, rhs)
}

func setnotin(lhs, rhs *v1.Expr) *v1.Expr {
	return v1.ExprBinary(lhs, v1.BinaryOp_SET_NOTIN, rhs)
}

func pipe(head, tail *v1.Expr) *v1.Expr {
	return v1.ExprPipe(head, tail)
}

func selector(x *v1.Expr, id string) *v1.Expr {
	return v1.ExprSelector(x, id)
}

func id(str string) *v1.Expr {
	return v1.ExprIdentifier(str)
}

func str(v string) *v1.Expr {
	return v1.ExprLiteral(v1.ValStr(v))
}

func i64(i int64) *v1.Expr {
	return v1.ExprLiteral(v1.ValI64(i))
}

func f64(f float64) *v1.Expr {
	return v1.ExprLiteral(v1.ValF64(f))
}

func boo(v bool) *v1.Expr {
	return v1.ExprLiteral(v1.ValBool(v))
}

func kv(k string, v *v1.Val) *v1.KV {
	return &v1.KV{Key: k, Value: v}
}

func obj(v ...*v1.KV) *v1.Expr {
	return v1.ExprLiteral(v1.ValObj(v...))
}

func arr(v ...*v1.Val) *v1.Expr {
	return v1.ExprLiteral(v1.ValArr(v...))
}

func fn(name string, e ...*v1.Expr) *v1.Expr {
	return v1.ExprFuncCall(name, e...)
}

func idx(x *v1.Expr, index *v1.Expr) *v1.Expr {
	return v1.ExprIndexor(x, index)
}
