package logql

import (
	"fmt"
	"strconv"
	"time"

	typesv1 "github.com/humanlogio/api/go/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func log(str string, args ...any) {
	fmt.Println(fmt.Sprintf(str, args...))
}

func Parse(input string) (_ *typesv1.LogQuery, err error) {
	p := logQL{Buffer: input}
	if err := p.Init(); err != nil {
		return nil, fmt.Errorf("initializing query parser: %v", err)
	}
	if err := p.Parse(); err != nil {
		return nil, fmt.Errorf("parsing query: %v (%#v)", err, err)
	}
	defer func() {
		e := recover()
		if e != nil {
			switch et := e.(type) {
			case error:
				err = et
			default:
				panic(e)
			}
		}
	}()
	p.Execute()
	if p.err != nil {
		err = p.err
	}
	if p.LogQuery == nil {
		return &typesv1.LogQuery{}, nil
	}
	return p.LogQuery, err
}

func (p *logQL) SetQuery(expr *typesv1.Expr) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.err != nil {
		p.err = fmt.Errorf("invalid `from`: %v", p.err)
	} else {
		p.LogQuery.Query = expr
	}
}

func (p *logQL) SetFrom(t time.Time) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.err != nil {
		p.err = fmt.Errorf("invalid `from`: %v", p.err)
	} else {
		p.LogQuery.From = timestamppb.New(t)
	}
}

func (p *logQL) SetTo(t time.Time) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.err != nil {
		p.err = fmt.Errorf("invalid `to`: %v", p.err)
	} else {
		p.LogQuery.To = timestamppb.New(t)
	}
}

func (p *logQL) SetContextMachine(op typesv1.BinaryOp_Operator, expr *typesv1.Expr) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.LogQuery.Context == nil {
		p.LogQuery.Context = new(typesv1.Context)
	}
	if p.err != nil {
		p.err = fmt.Errorf("invalid `to`: %v", p.err)
	} else {
		p.LogQuery.Context.MachineId = typesv1.ExprBinary(
			typesv1.ExprIdentifier("machine"),
			op,
			expr,
		)
	}
}

func (p *logQL) SetContextSession(op typesv1.BinaryOp_Operator, expr *typesv1.Expr) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.LogQuery.Context == nil {
		p.LogQuery.Context = new(typesv1.Context)
	}
	if p.err != nil {
		p.err = fmt.Errorf("invalid `to`: %v", p.err)
	} else {
		p.LogQuery.Context.SessionId = typesv1.ExprBinary(
			typesv1.ExprIdentifier("session"),
			op,
			expr,
		)
	}
}

func (p *logQL) pushExpr(e *typesv1.Expr) {
	p.Exprs = append(p.Exprs, e)
}

func (p *logQL) popExpr() *typesv1.Expr {
	cur := len(p.Exprs) - 1
	out := p.Exprs[cur]
	p.Exprs = p.Exprs[:cur]
	return out
}

func (p *logQL) parseString(text string) string {
	v, err := strconv.Unquote(text)
	if err != nil {
		p.err = err
		panic(err)
	}
	return v
}

func (p *logQL) parseFloat64(text string) float64 {
	v, err := strconv.ParseFloat(text, 64)
	if err != nil {
		p.err = err
		panic(err)
	}
	return v
}

func (p *logQL) parseInt64(text string) int64 {
	v, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		p.err = err
		panic(err)
	}
	return v
}

func (p *logQL) parseDurationF64(v float64, unit string) time.Duration {
	switch unit {
	case "us", "µs":
		return time.Duration(v * float64(time.Microsecond))
	case "ms":
		return time.Duration(v * float64(time.Millisecond))
	case "s":
		return time.Duration(v * float64(time.Second))
	case "m":
		return time.Duration(v * float64(time.Minute))
	case "h":
		return time.Duration(v * float64(time.Hour))
	default:
		panic(unit)
	}
}

func (p *logQL) parseDurationI64(v int64, unit string) time.Duration {
	switch unit {
	case "ns":
		return time.Duration(v)
	case "us", "µs":
		return time.Duration(v) * time.Microsecond
	case "ms":
		return time.Duration(v) * time.Millisecond
	case "s":
		return time.Duration(v) * time.Second
	case "m":
		return time.Duration(v) * time.Minute
	case "h":
		return time.Duration(v) * time.Hour
	default:
		panic(unit)
	}
}

func (p *logQL) parseTime(layout string, text string) time.Time {
	v, err := time.Parse(layout, text)
	if err != nil {
		p.err = err
		panic(err)
	}
	return v
}

func (p *logQL) pushObj() {
	p.ObjsKVs = append(p.ObjsKVs, []*typesv1.KV{})
}

func (p *logQL) startObjItem(key string) {
	cur := len(p.ObjsKVs) - 1
	p.ObjsKVs[cur] = append(p.ObjsKVs[cur], &typesv1.KV{Key: key})
}

func (p *logQL) closeObjItem(val *typesv1.Val) {
	cur := len(p.ObjsKVs) - 1
	keyi := len(p.ObjsKVs[cur]) - 1
	p.ObjsKVs[cur][keyi].Value = val
}

func (p *logQL) popObj() []*typesv1.KV {
	cur := len(p.ObjsKVs) - 1
	pop := p.ObjsKVs[cur]
	p.ObjsKVs = p.ObjsKVs[:cur]
	return pop
}

func (p *logQL) pushArray() {
	p.Arrs = append(p.Arrs, []*typesv1.Val{})
}

func (p *logQL) addArrItem(v *typesv1.Val) {
	cur := len(p.Arrs) - 1
	p.Arrs[cur] = append(p.Arrs[cur], v)
}

func (p *logQL) popArray() []*typesv1.Val {
	cur := len(p.Arrs) - 1
	pop := p.Arrs[cur]
	p.Arrs = p.Arrs[:cur]
	return pop
}

func (p *logQL) pushFunc() {
	p.FuncCalls = append(p.FuncCalls, &typesv1.FuncCall{})
}

func (p *logQL) setFuncName(name string) {
	cur := len(p.FuncCalls) - 1
	p.FuncCalls[cur].Name = name
}

func (p *logQL) addFuncArg(e *typesv1.Expr) {
	cur := len(p.FuncCalls) - 1
	p.FuncCalls[cur].Args = append(p.FuncCalls[cur].Args, e)
}

func (p *logQL) popFunc() *typesv1.Expr {
	cur := len(p.FuncCalls) - 1
	pop := p.FuncCalls[cur]
	p.FuncCalls = p.FuncCalls[:cur]
	return typesv1.ExprFuncCall(pop.Name, pop.Args...)
}
