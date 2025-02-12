package logql

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	typesv1 "github.com/humanlogio/api/go/types/v1"
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
		{`{ from == now()-1s }`, q(tr(sub(fn("now"), dure(time.Second)), nil), nil, nil), nil},
		{
			`{session==1734666428038101000} filter msg=="AuthenticateUser"`,
			q(
				nil,
				qctx(nil, eq(id("session"), i64(1734666428038101000))),
				stmts(
					filter(
						eq(
							id("msg"),
							str("AuthenticateUser"),
						),
					),
				),
			),
			nil,
		},
		{
			`{session==1734666428038101000} filter msg=="AuthenticateUser" | filter msg=="AuthenticateUser"`,
			q(
				nil,
				qctx(nil, eq(id("session"), i64(1734666428038101000))),
				stmts(
					filter(
						eq(
							id("msg"),
							str("AuthenticateUser"),
						),
					),
					filter(
						eq(
							id("msg"),
							str("AuthenticateUser"),
						),
					),
				),
			),
			nil,
		},
		{
			`{session==1734666428038101000} where msg=="AuthenticateUser"`,
			q(
				nil,
				qctx(nil, eq(id("session"), i64(1734666428038101000))),
				stmts(
					filter(
						eq(
							id("msg"),
							str("AuthenticateUser"),
						),
					),
				),
			),
			nil,
		},
		{
			`where msg=="AuthenticateUser"`,
			q(
				nil,
				nil,
				stmts(
					filter(
						eq(
							id("msg"),
							str("AuthenticateUser"),
						),
					),
				),
			),
			nil,
		},
		{
			`summarize count()`,
			q(
				nil,
				nil,
				stmts(
					summarize(
						summarizeParams(
							summarizeUnnamedParam(fnc("count")),
						),
						nil,
					),
				),
			),
			nil,
		},
		{
			`summarize count() by msg`,
			q(
				nil,
				nil,
				stmts(
					summarize(
						summarizeParams(
							summarizeUnnamedParam(fnc("count")),
						),
						summarizeGroupExpressions(
							summarizeUnnamedGroup(id("msg")),
						),
					),
				),
			),
			nil,
		},
		{
			`summarize count(), avg() by msg`,
			q(
				nil,
				nil,
				stmts(
					summarize(
						summarizeParams(
							summarizeUnnamedParam(fnc("count")),
							summarizeUnnamedParam(fnc("avg")),
						),
						summarizeGroupExpressions(
							summarizeUnnamedGroup(id("msg")),
						),
					),
				),
			),
			nil,
		},
		{
			`summarize count() by msg,lvl`,
			q(
				nil,
				nil,
				stmts(
					summarize(
						summarizeParams(
							summarizeUnnamedParam(fnc("count")),
						),
						summarizeGroupExpressions(
							summarizeUnnamedGroup(id("msg")),
							summarizeUnnamedGroup(id("lvl")),
						),
					),
				),
			),
			nil,
		},
		{
			`summarize msg_count = count() by msg,lvl`,
			q(
				nil,
				nil,
				stmts(
					summarize(
						summarizeParams(
							summarizeNamedParam("msg_count", fnc("count")),
						),
						summarizeGroupExpressions(
							summarizeUnnamedGroup(id("msg")),
							summarizeUnnamedGroup(id("lvl")),
						),
					),
				),
			),
			nil,
		},
		{
			`summarize msg_count = count(), msg_avg = avg() by msg,lvl`,
			q(
				nil,
				nil,
				stmts(
					summarize(
						summarizeParams(
							summarizeNamedParam("msg_count", fnc("count")),
							summarizeNamedParam("msg_avg", fnc("avg")),
						),
						summarizeGroupExpressions(
							summarizeUnnamedGroup(id("msg")),
							summarizeUnnamedGroup(id("lvl")),
						),
					),
				),
			),
			nil,
		},
		{
			`summarize count(), msg_avg = avg() by msg, level=lvl`,
			q(
				nil,
				nil,
				stmts(
					summarize(
						summarizeParams(
							summarizeUnnamedParam(fnc("count")),
							summarizeNamedParam("msg_avg", fnc("avg")),
						),
						summarizeGroupExpressions(
							summarizeUnnamedGroup(id("msg")),
							summarizeNamedGroup("level", id("lvl")),
						),
					),
				),
			),
			nil,
		},
		{
			`where msg=="request start" | summarize avg(kv["req_ms"]) by msg`,
			q(
				nil,
				nil,
				stmts(
					filter(
						eq(
							id("msg"),
							str("request start"),
						),
					),
					summarize(
						summarizeParams(
							summarizeUnnamedParam(
								fnc("avg", idx(id("kv"), str("req_ms"))),
							),
						),
						summarizeGroupExpressions(
							summarizeUnnamedGroup(id("msg")),
						),
					),
				),
			),
			nil,
		},
		{
			`project msg_len = len(msg)`,
			q(
				nil,
				nil,
				stmts(
					project(
						projection(
							"msg_len",
							fn("len", id("msg")),
						),
					),
				),
			),
			nil,
		},
		{
			`project msg`,
			q(
				nil,
				nil,
				stmts(
					project(
						projection(
							"msg",
							nil,
						),
					),
				),
			),
			nil,
		},
		{
			`project_away msg`,
			q(
				nil,
				nil,
				stmts(
					projectAway(
						projectionAway(
							"msg",
						),
					),
				),
			),
			nil,
		},
		{
			`project_keep msg, lvl`,
			q(
				nil,
				nil,
				stmts(
					projectKeep(
						projectionKeep("msg"),
						projectionKeep("lvl"),
					),
				),
			),
			nil,
		},
		{
			`count`,
			q(
				nil,
				nil,
				stmts(
					countOp(),
				),
			),
			nil,
		},
		{
			`distinct`,
			q(
				nil,
				nil,
				stmts(
					distinctOp(),
				),
			),
			nil,
		},
		{
			`sample 10`,
			q(
				nil,
				nil,
				stmts(
					sampleOp(10),
				),
			),
			nil,
		},
		{
			`search msg:"hello world"`,
			q(
				nil,
				nil,
				stmts(
					searchOp_Field("msg", "hello world", nil),
				),
			),
			nil,
		},
		{
			`search msg=="hello world"`,
			q(
				nil,
				nil,
				stmts(
					searchOp_Exact("msg", "hello world", nil),
				),
			),
			nil,
		},
		{
			`search msg matches regex "hello world"`,
			q(
				nil,
				nil,
				stmts(
					searchOp_Regex("msg", "hello world", nil),
				),
			),
			nil,
		},
		{
			`search kind=default msg=="hello world"`,
			q(
				nil,
				nil,
				stmts(
					searchOp_Exact("msg", "hello world", searchKindDefault()),
				),
			),
			nil,
		},
		{
			`search kind=case_insensitive msg=="hello world"`,
			q(
				nil,
				nil,
				stmts(
					searchOp_Exact("msg", "hello world", searchKindCaseInsensitive()),
				),
			),
			nil,
		},
		{
			`search kind=case_sensitive msg=="hello world"`,
			q(
				nil,
				nil,
				stmts(
					searchOp_Exact("msg", "hello world", searchKindCaseSensitive()),
				),
			),
			nil,
		},
		{
			`sort by msg`,
			q(
				nil,
				nil,
				stmts(
					sortOp(
						sortBy("msg", nil),
					),
				),
			),
			nil,
		},
		{
			`sort by msg, lvl asc`,
			q(
				nil,
				nil,
				stmts(
					sortOp(
						sortBy("msg", nil),
						sortBy("lvl", sortAsc()),
					),
				),
			),
			nil,
		},
		{
			`take 10`,
			q(
				nil,
				nil,
				stmts(
					takeOp(10),
				),
			),
			nil,
		},
		{
			`top 10 by msg`,
			q(
				nil,
				nil,
				stmts(
					topOp(10, id("msg"), nil),
				),
			),
			nil,
		},
		{
			`top 10 by msg asc`,
			q(
				nil,
				nil,
				stmts(
					topOp(10, id("msg"), topAsc()),
				),
			),
			nil,
		},
		{
			`top 10 by msg desc`,
			q(
				nil,
				nil,
				stmts(
					topOp(10, id("msg"), topDesc()),
				),
			),
			nil,
		},
		{
			`top 10 by kv["hello"]`,
			q(
				nil,
				nil,
				stmts(
					topOp(
						10,
						idx(
							id("kv"),
							str("hello"),
						),
						nil,
					),
				),
			),
			nil,
		},
		{
			`| render split by msg`,
			q(
				nil,
				nil,
				render(
					nil,
					splitby(
						id("msg"),
					),
				),
			),
			nil,
		},
		{
			`filter true | render split by msg`,
			q(
				nil, nil,
				render(
					stmts(
						filter(boo(true)),
					),
					splitby(
						id("msg"),
					),
				),
			),
			nil,
		},
		{
			`{from==2006-01-02T15:04:05.999999999+07:00}`,
			q(
				tr(tse("2006-01-02T15:04:05.999999999+07:00"), nil),
				nil,
				nil,
			),
			nil,
		},
		{
			`{from==2006-01-02T15:04:06.00000000Z}`,
			q(
				tr(tse("2006-01-02T15:04:06.00000000Z"), nil),
				nil,
				nil,
			),
			nil,
		},
		// {`{from==2006-01-02T15:04:05+07:00}`, q(ts("2006-01-02T15:04:05+07:00"), nil, nil, nil), nil},
		// {`{to==2006-01-02T15:04:05.999999999+07:00}`, q(nil, ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil), nil},
		// {`{to==2006-01-02T15:04:05+07:00}`, q(nil, ts("2006-01-02T15:04:05+07:00"), nil, nil), nil},
		// {`{from==2006-01-02T15:04:07+07:00 to==2006-01-02T15:04:05+07:00}`, q(ts("2006-01-02T15:04:07+07:00"), ts("2006-01-02T15:04:05+07:00"), nil, nil), nil},
		// {`{	from	==	2006-01-02T15:04:05.999999999+07:00}`, q(ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil, nil), nil},
		// {`{	from	==	2006-01-02T15:04:05.999999999+07:00		}`, q(ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil, nil), nil},
		// {`{	from	==	2006-01-02T15:04:05.999999999+07:00		to	==	2006-01-02T15:04:05.999999999+07:00		}`, q(ts("2006-01-02T15:04:05.999999999+07:00"), ts("2006-01-02T15:04:05.999999999+07:00"), nil, nil), nil},
		// {`{session==1}`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		// {`{session == 1}`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		// {`{session ==1}`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		// {`{session== 1}`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		// {`{ session  == 1 }`, q(nil, nil, qctx(nil, eq(id("session"), i64(1))), nil), nil},
		// {`{machine==1}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		// {`{machine == 1}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		// {`{machine ==1}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		// {`{machine== 1}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		// {`{ machine  == 1 }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), nil), nil), nil},
		// {`{machine==1 session==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{machine ==1 session==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{machine==1 session ==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{machine==1 session== 2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{ machine==1 session==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{ machine ==1 session==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{ machine==1 session ==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{ machine==1 session== 2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{ machine==1 session==2   }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{ machine ==1 session==2   }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{ machine==1 session ==2   }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{ machine==1 session== 2   }`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{machine==1	session==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{machine	==1	session==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{machine==1	session	==2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`{machine==1	session==	2}`, q(nil, nil, qctx(eq(id("machine"), i64(1)), eq(id("session"), i64(2))), nil), nil},
		// {`lvl`, q(nil, nil, nil, id("lvl")), nil},
		// {`msg`, q(nil, nil, nil, id("msg")), nil},
		// {`kv`, q(nil, nil, nil, id("kv")), nil},
		// {`"hello world"`, q(nil, nil, nil, str("hello world")), nil},
		// {`true`, q(nil, nil, nil, boo(true)), nil},
		// {`false`, q(nil, nil, nil, boo(false)), nil},
		// {`.110`, q(nil, nil, nil, f64(.11)), nil},
		// {`1111.0`, q(nil, nil, nil, f64(1111)), nil},
		// {`10.3e9`, q(nil, nil, nil, f64(10.3e9)), nil},
		// {`-0`, q(nil, nil, nil, neg(i64(0))), nil},
		// {`-10000`, q(nil, nil, nil, neg(i64(10000))), nil},
		// {`0`, q(nil, nil, nil, i64(0)), nil},
		// {`1`, q(nil, nil, nil, i64(1)), nil},
		// {`-1`, q(nil, nil, nil, neg(i64(1))), nil},
		// {`""`, q(nil, nil, nil, str("")), nil},
		// {`"hello world"`, q(nil, nil, nil, str("hello world")), nil},
		// {`"hello world\n\t"`, q(nil, nil, nil, str("hello world\n\t")), nil},
		// {`2006-01-02T15:04:05.999999999+07:00`, q(nil, nil, nil, tse("2006-01-02T15:04:05.999999999+07:00")), nil},
		// {`ts:2006-01-02T15:04:05.999999999+07:00`, q(nil, nil, nil, tse("2006-01-02T15:04:05.999999999+07:00")), nil},
		// {`{"hello":"world", "hello2":1}`, q(nil, nil, nil, obj(kv("hello", v1.ValStr("world")), kv("hello2", v1.ValI64(1)))), nil},
		// {`{"hello":"world"}`, q(nil, nil, nil, obj(kv("hello", v1.ValStr("world")))), nil},
		// {`{"hello":1}`, q(nil, nil, nil, obj(kv("hello", v1.ValI64(1)))), nil},
		// {`{"hello":ts:2006-01-02T15:04:05.999999999+07:00}`, q(nil, nil, nil, obj(kv("hello", v1.ValTimestamp(ts("2006-01-02T15:04:05.999999999+07:00"))))), nil},
		// {`{"hello":2006-01-02T15:04:05.999999999+07:00}`, q(nil, nil, nil, obj(kv("hello", v1.ValTimestamp(ts("2006-01-02T15:04:05.999999999+07:00"))))), nil},
		// {`{}`, q(nil, nil, nil, obj()), nil},
		// {`[]`, q(nil, nil, nil, arr()), nil},
		// {`[1, "deux", 3.0]`, q(nil, nil, nil, arr(v1.ValI64(1), v1.ValStr("deux"), v1.ValF64(3.0))), nil},
		// {`[[]]`, q(nil, nil, nil, arr(v1.ValArr())), nil},
		// {`[[],[]]`, q(nil, nil, nil, arr(v1.ValArr(), v1.ValArr())), nil},
		// {`[[1]]`, q(nil, nil, nil, arr(v1.ValArr(v1.ValI64(1)))), nil},
		// {`[[],[1],[2,3]]`, q(nil, nil, nil, arr(v1.ValArr(), v1.ValArr(v1.ValI64(1)), v1.ValArr(v1.ValI64(2), v1.ValI64(3)))), nil},
		// {`[{}]`, q(nil, nil, nil, arr(v1.ValObj())), nil},
		// {`[{"hello":[]}]`, q(nil, nil, nil, arr(v1.ValObj(kv("hello", v1.ValArr())))), nil},
		// {`[{"hello":[{},{}]}]`, q(nil, nil, nil, arr(v1.ValObj(kv("hello", v1.ValArr(v1.ValObj(), v1.ValObj()))))), nil},
		// {`1000s`, q(nil, nil, nil, dure(1000*time.Second)), nil},
		// {`1.5us`, q(nil, nil, nil, dure(1500*time.Nanosecond)), nil},
		// {`1.5µs`, q(nil, nil, nil, dure(1500*time.Nanosecond)), nil},
		// {`1.5ms`, q(nil, nil, nil, dure(1500*time.Microsecond)), nil},
		// {`15ns`, q(nil, nil, nil, dure(15*time.Nanosecond)), nil},
		// {`1.5s`, q(nil, nil, nil, dure(1500*time.Millisecond)), nil},
		// {`-1.5s`, q(nil, nil, nil, neg(dure(1500*time.Millisecond))), nil},
		// {`-1.5m`, q(nil, nil, nil, neg(dure(90*time.Second))), nil},
		// {`-1.5h`, q(nil, nil, nil, neg(dure(90*time.Minute))), nil},
		// {`now()`, q(nil, nil, nil, fn("now")), nil},
		// {`now(1h)`, q(nil, nil, nil, fn("now", dure(1*time.Hour))), nil},
		// {`fn(fn1(),fn2(fn3()))`, q(nil, nil, nil, fn("fn", fn("fn1"), fn("fn2", fn("fn3")))), nil},
		// {`len(msg)`, q(nil, nil, nil, fn("len", id("msg"))), nil},
		// {`!true`, q(nil, nil, nil, not(boo(true))), nil},
		// {`! true`, q(nil, nil, nil, not(boo(true))), nil},
		// {`!(eval())`, q(nil, nil, nil, not(fn("eval"))), nil},
		// {`!( eval())`, q(nil, nil, nil, not(fn("eval"))), nil},
		// {`1 + 2`, q(nil, nil, nil, add(i64(1), i64(2))), nil},
		// {`1 + 2 + 3`, q(nil, nil, nil, add(add(i64(1), i64(2)), i64(3))), nil},
		// {`1 - 2`, q(nil, nil, nil, sub(i64(1), i64(2))), nil},
		// {`1 - 2 - 3`, q(nil, nil, nil, sub(sub(i64(1), i64(2)), i64(3))), nil},
		// {`1 * 2`, q(nil, nil, nil, mul(i64(1), i64(2))), nil},
		// {`-1 * 2`, q(nil, nil, nil, mul(neg(i64(1)), i64(2))), nil},
		// {`1 / 2`, q(nil, nil, nil, div(i64(1), i64(2))), nil},
		// {`1 + 2 * 3`, q(nil, nil, nil, add(i64(1), mul(i64(2), i64(3)))), nil},
		// {`1 - 2 * 3`, q(nil, nil, nil, sub(i64(1), mul(i64(2), i64(3)))), nil},
		// {`1 + 2 / 3`, q(nil, nil, nil, add(i64(1), div(i64(2), i64(3)))), nil},
		// {`1 - 2 / 3`, q(nil, nil, nil, sub(i64(1), div(i64(2), i64(3)))), nil},
		// {`1 == 2`, q(nil, nil, nil, eq(i64(1), i64(2))), nil},
		// {`1 != 2`, q(nil, nil, nil, noteq(i64(1), i64(2))), nil},
		// {`1 > 2`, q(nil, nil, nil, gt(i64(1), i64(2))), nil},
		// {`1 >= 2`, q(nil, nil, nil, gte(i64(1), i64(2))), nil},
		// {`1 < 2`, q(nil, nil, nil, lt(i64(1), i64(2))), nil},
		// {`1 <= 2`, q(nil, nil, nil, lte(i64(1), i64(2))), nil},
		// {`true && false`, q(nil, nil, nil, and(boo(true), boo(false))), nil},
		// {`true || false`, q(nil, nil, nil, or(boo(true), boo(false))), nil},
		// {`true && 1 > 2`, q(nil, nil, nil, and(boo(true), gt(i64(1), i64(2)))), nil},
		// {`1 > 2 && true`, q(nil, nil, nil, and(gt(i64(1), i64(2)), boo(true))), nil},
		// {`!true || false`, q(nil, nil, nil, or(not(boo(true)), boo(false))), nil},
		// {`!(true || false)`, q(nil, nil, nil, not(or(boo(true), boo(false)))), nil},
		// {`true && !false`, q(nil, nil, nil, and(boo(true), not(boo(false)))), nil},
		// {`!(true && false)`, q(nil, nil, nil, not(and(boo(true), boo(false)))), nil},
		// {`1 in [1,2,3]`, q(nil, nil, nil, setin(i64(1), arr(v1.ValI64(1), v1.ValI64(2), v1.ValI64(3)))), nil},
		// {`1 not in [1,2,3]`, q(nil, nil, nil, setnotin(i64(1), arr(v1.ValI64(1), v1.ValI64(2), v1.ValI64(3)))), nil},
		// {`true || false || 1`, q(nil, nil, nil, or(or(boo(true), boo(false)), i64(1))), nil},
		// {`len(msg) > 0 | true`, q(nil, nil, nil, pipe(gt(fn("len", id("msg")), i64(0)), boo(true))), nil},
		// {`len(msg) > 0 | true | lvl in [1,2,3]`, q(nil, nil, nil, pipe(pipe(gt(fn("len", id("msg")), i64(0)), boo(true)), setin(id("lvl"), arr(v1.ValI64(1), v1.ValI64(2), v1.ValI64(3))))), nil},
		// {`kv["hello"]["world"][0]`, q(nil, nil, nil, idx(idx(idx(id("kv"), str("hello")), str("world")), i64(0))), nil},
		// {` (  kv ["hello"] ["world"] [0]  ) `, q(nil, nil, nil, idx(idx(idx(id("kv"), str("hello")), str("world")), i64(0))), nil},
		// {`msg=="hello world" | len(lvl) > 0`, q(nil, nil, nil, pipe(eq(id("msg"), str("hello world")), gt(fn("len", id("lvl")), i64(0)))), nil},
		// {`kv.hello.world`, q(nil, nil, nil, selector(selector(id("kv"), "hello"), "world")), nil},
		// {`lvl.k.v | len(lvl)`, q(nil, nil, nil, pipe(selector(selector(id("lvl"), "k"), "v"), fn("len", id("lvl")))), nil},
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

func q(tr *v1.Timerange, ctx *v1.Context, s *v1.Statements) *v1.LogQuery {
	return &v1.LogQuery{Timerange: tr, Context: ctx, Query: s}
}

func stmts(stmts ...*v1.Statement) *v1.Statements {
	return &v1.Statements{Statements: stmts}
}

func splitby(scalars ...*v1.Expr) *v1.RenderStatement {
	return &v1.RenderStatement{Stmt: &v1.RenderStatement_Split{
		Split: &v1.SplitOperator{By: &v1.SplitOperator_ByOperator{
			Scalars: scalars,
		}},
	}}
}

func render(stmts *v1.Statements, render *v1.RenderStatement) *v1.Statements {
	if stmts == nil {
		stmts = new(v1.Statements)
	}
	stmts.Render = render
	return stmts
}

func filter(e *v1.Expr) *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_Filter{
		Filter: &v1.FilterOperator{Expr: e},
	}}
}

func projection(column string, value *v1.Expr) *v1.ProjectOperator_Projection {
	return &v1.ProjectOperator_Projection{
		Column: &v1.Identifier{Name: column},
		Value:  value,
	}
}

func project(pjs ...*v1.ProjectOperator_Projection) *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_Project{
		Project: &v1.ProjectOperator{Projections: pjs},
	}}
}

func projectionAway(column string) *v1.ProjectAwayOperator_Projection {
	return &v1.ProjectAwayOperator_Projection{
		Column: &v1.Identifier{Name: column},
	}
}

func projectAway(pjs ...*v1.ProjectAwayOperator_Projection) *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_ProjectAway{
		ProjectAway: &v1.ProjectAwayOperator{Projections: pjs},
	}}
}

func projectionKeep(column string) *v1.ProjectKeepOperator_Projection {
	return &v1.ProjectKeepOperator_Projection{
		Column: &v1.Identifier{Name: column},
	}
}

func projectKeep(pjs ...*v1.ProjectKeepOperator_Projection) *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_ProjectKeep{
		ProjectKeep: &v1.ProjectKeepOperator{Projections: pjs},
	}}
}

func countOp() *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_Count{
		Count: &v1.CountOperator{},
	}}
}

func distinctOp() *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_Distinct{
		Distinct: &v1.DistinctOperator{},
	}}
}

func sampleOp(v int64) *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_Sample{
		Sample: &v1.SampleOperator{Count: v},
	}}
}

func searchKindDefault() *typesv1.SearchOperator_Kind {
	v := typesv1.SearchOperator_Default
	return &v
}

func searchKindCaseInsensitive() *typesv1.SearchOperator_Kind {
	v := typesv1.SearchOperator_CaseInsensitive
	return &v
}

func searchKindCaseSensitive() *typesv1.SearchOperator_Kind {
	v := typesv1.SearchOperator_CaseSensitive
	return &v
}

func searchOp_Literal(literal string, kind *typesv1.SearchOperator_Kind) *typesv1.Statement {
	return &v1.Statement{Stmt: &typesv1.Statement_Search{
		Search: &v1.SearchOperator{
			Predicate: &v1.SearchOperator_Literal_{Literal: literal},
			Kind:      kind,
		},
	}}
}

func searchOp_Field(column, literal string, kind *typesv1.SearchOperator_Kind) *typesv1.Statement {
	return &v1.Statement{Stmt: &typesv1.Statement_Search{
		Search: &v1.SearchOperator{
			Predicate: &v1.SearchOperator_Field{Field: &v1.SearchOperator_FieldSearch{
				Column:  column,
				Literal: literal,
			}},
			Kind: kind,
		},
	}}
}

func searchOp_Exact(column, literal string, kind *typesv1.SearchOperator_Kind) *typesv1.Statement {
	return &v1.Statement{Stmt: &typesv1.Statement_Search{
		Search: &v1.SearchOperator{
			Predicate: &v1.SearchOperator_Exact{Exact: &v1.SearchOperator_ExactSearch{
				Column:  column,
				Literal: literal,
			}},
			Kind: kind,
		},
	}}
}

func searchOp_Regex(column, regex string, kind *typesv1.SearchOperator_Kind) *typesv1.Statement {
	return &v1.Statement{Stmt: &typesv1.Statement_Search{
		Search: &v1.SearchOperator{
			Predicate: &v1.SearchOperator_Regex{Regex: &v1.SearchOperator_RegexSearch{
				Column: column,
				Regex:  regex,
			}},
			Kind: kind,
		},
	}}
}

func sortAsc() *typesv1.SortOperator_Order {
	v := typesv1.SortOperator_Asc
	return &v
}

func sortDesc() *typesv1.SortOperator_Order {
	v := typesv1.SortOperator_Desc
	return &v
}

func sortBy(colName string, order *v1.SortOperator_Order) *v1.SortOperator_ByColumn {
	return &v1.SortOperator_ByColumn{
		Column: &v1.Identifier{Name: colName},
		Order:  order,
	}
}

func sortOp(col *v1.SortOperator_ByColumn, cols ...*v1.SortOperator_ByColumn) *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_Sort{
		Sort: &v1.SortOperator{
			ByColumns: append([]*v1.SortOperator_ByColumn{
				col,
			}, cols...),
		},
	}}
}

func takeOp(v int64) *v1.Statement {
	return &v1.Statement{Stmt: &v1.Statement_Take{
		Take: &v1.TakeOperator{Count: v},
	}}
}

func topAsc() *typesv1.TopOperator_Order {
	v := typesv1.TopOperator_Asc
	return &v
}

func topDesc() *typesv1.TopOperator_Order {
	v := typesv1.TopOperator_Desc
	return &v
}

func topOp(v int64, by *typesv1.Expr, order *typesv1.TopOperator_Order) *v1.Statement {
	return &v1.Statement{Stmt: &typesv1.Statement_Top{
		Top: &typesv1.TopOperator{
			Count: v,
			ByColumn: &v1.TopOperator_ByColumn{
				Scalar: by,
				Order:  order,
			},
		},
	}}
}

func summarizeNamedParam(name string, aggregate *v1.FuncCall) *v1.SummarizeOperator_Parameter {
	return &v1.SummarizeOperator_Parameter{
		Column:            &v1.Identifier{Name: name},
		AggregateFunction: aggregate,
	}
}

func summarizeUnnamedParam(aggregate *v1.FuncCall) *v1.SummarizeOperator_Parameter {
	return &v1.SummarizeOperator_Parameter{
		AggregateFunction: aggregate,
	}
}

func summarizeNamedGroup(name string, expr *v1.Expr) *v1.SummarizeOperator_ByGroupExpression {
	return &v1.SummarizeOperator_ByGroupExpression{
		Column: &v1.Identifier{Name: name},
		Scalar: expr,
	}
}

func summarizeUnnamedGroup(expr *v1.Expr) *v1.SummarizeOperator_ByGroupExpression {
	return &v1.SummarizeOperator_ByGroupExpression{
		Scalar: expr,
	}
}

func summarizeParams(params ...*v1.SummarizeOperator_Parameter) *v1.SummarizeOperator_Parameters {
	return &v1.SummarizeOperator_Parameters{Parameters: params}
}

func summarizeGroupExpressions(ges ...*v1.SummarizeOperator_ByGroupExpression) *v1.SummarizeOperator_ByGroupExpressions {
	return &v1.SummarizeOperator_ByGroupExpressions{Groups: ges}
}

func summarize(params *v1.SummarizeOperator_Parameters, ges *v1.SummarizeOperator_ByGroupExpressions) *v1.Statement {
	op := &v1.SummarizeOperator{
		Parameters:         params,
		ByGroupExpressions: ges,
	}
	return &v1.Statement{Stmt: &v1.Statement_Summarize{
		Summarize: op,
	}}
}

func tr(from, to *v1.Expr) *v1.Timerange {
	return &typesv1.Timerange{From: from, To: to}
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

func fnc(name string, e ...*v1.Expr) *v1.FuncCall {
	return &v1.FuncCall{Name: name, Args: e}
}

func idx(x *v1.Expr, index *v1.Expr) *v1.Expr {
	return v1.ExprIndexor(x, index)
}
