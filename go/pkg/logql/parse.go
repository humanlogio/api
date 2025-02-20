package logql

import (
	"fmt"
	"strconv"
	"time"

	typesv1 "github.com/humanlogio/api/go/types/v1"
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
		if perr, ok := err.(*parseError); ok {
			return nil, perr
		}
		return nil, fmt.Errorf("parsing query: %v (%#v)", err, err)
	}
	defer func() {
		// e := recover()
		// if e != nil {
		// 	switch et := e.(type) {
		// 	case error:
		// 		err = et
		// 	default:
		// 		panic(e)
		// 	}
		// }
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

func (p *logQL) SetQuery(stmts []*typesv1.Statement) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.err != nil {
		p.err = fmt.Errorf("invalid query: %v", p.err)
	} else {
		p.LogQuery.Query = &typesv1.Statements{Statements: stmts}
	}
}

func (p *logQL) SetFrom(expr *typesv1.Expr) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.LogQuery.Timerange == nil {
		p.LogQuery.Timerange = new(typesv1.Timerange)
	}
	if p.err != nil {
		p.err = fmt.Errorf("invalid `from`: %v", p.err)
	} else {
		p.LogQuery.Timerange.From = expr
	}
}

func (p *logQL) SetTo(expr *typesv1.Expr) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.LogQuery.Timerange == nil {
		p.LogQuery.Timerange = new(typesv1.Timerange)
	}
	if p.err != nil {
		p.err = fmt.Errorf("invalid `to`: %v", p.err)
	} else {
		p.LogQuery.Timerange.To = expr
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

func (p *logQL) addFilterStatement(op *typesv1.FilterOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Filter{Filter: op},
	})
}

func (p *logQL) addSummarizeStatement(op *typesv1.SummarizeOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Summarize{Summarize: op},
	})
}

func (p *logQL) addProjectStatement(op *typesv1.ProjectOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Project{Project: op},
	})
}

func (p *logQL) addProjectAwayStatement(op *typesv1.ProjectAwayOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_ProjectAway{ProjectAway: op},
	})
}

func (p *logQL) addProjectKeepStatement(op *typesv1.ProjectKeepOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_ProjectKeep{ProjectKeep: op},
	})
}

func (p *logQL) addExtendStatement(op *typesv1.ExtendOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Extend{Extend: op},
	})
}

func (p *logQL) addCountStatement(op *typesv1.CountOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Count{Count: op},
	})
}

func (p *logQL) addDistinctStatement(op *typesv1.DistinctOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Distinct{Distinct: op},
	})
}

func (p *logQL) addSampleStatement(op *typesv1.SampleOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Sample{Sample: op},
	})
}

func (p *logQL) addSearchStatement(op *typesv1.SearchOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Search{Search: op},
	})
}

func (p *logQL) addSortStatement(op *typesv1.SortOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Sort{Sort: op},
	})
}

func (p *logQL) addTakeStatement(op *typesv1.TakeOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Take{Take: op},
	})
}

func (p *logQL) addTopStatement(op *typesv1.TopOperator) {
	p.Stmts = append(p.Stmts, &typesv1.Statement{
		Stmt: &typesv1.Statement_Top{Top: op},
	})
}

func (p *logQL) setRenderSplitByStatement(op *typesv1.SplitOperator) {
	if p.LogQuery == nil {
		p.LogQuery = new(typesv1.LogQuery)
	}
	if p.LogQuery.Query == nil {
		p.LogQuery.Query = new(typesv1.Statements)
	}
	p.LogQuery.Query.Render = &typesv1.RenderStatement{
		Stmt: &typesv1.RenderStatement_Split{Split: op},
	}
}

func (p *logQL) setFilterOp(e *typesv1.Expr) {
	p.FilterOp = &typesv1.FilterOperator{Expr: e}
}

func (p *logQL) startSummarizeOp() {
	p.SummarizeOp = &typesv1.SummarizeOperator{
		Parameters: &typesv1.SummarizeOperator_Parameters{},
	}
}

func (p *logQL) addSummarizeParameterUnnamedFunc(fn *typesv1.FuncCall) {
	p.SummarizeOp.Parameters.Parameters = append(p.SummarizeOp.Parameters.Parameters,
		&typesv1.SummarizeOperator_Parameter{AggregateFunction: fn},
	)
}

func (p *logQL) startSummarizeParameterNamedFunc(text string) {
	p.SummarizeOp.Parameters.Parameters = append(p.SummarizeOp.Parameters.Parameters,
		&typesv1.SummarizeOperator_Parameter{
			Column: &typesv1.Identifier{Name: text},
		},
	)
}

func (p *logQL) endSummarizeParameterNamedFunc(fn *typesv1.FuncCall) {
	top := len(p.SummarizeOp.Parameters.Parameters) - 1
	p.SummarizeOp.Parameters.Parameters[top].AggregateFunction = fn
}

func (p *logQL) addSummarizeByUnnamedGroupExpression(expr *typesv1.Expr) {
	if p.SummarizeOp.ByGroupExpressions == nil {
		p.SummarizeOp.ByGroupExpressions = new(typesv1.SummarizeOperator_ByGroupExpressions)
	}
	p.SummarizeOp.ByGroupExpressions.Groups = append(p.SummarizeOp.ByGroupExpressions.Groups, &typesv1.SummarizeOperator_ByGroupExpression{
		Scalar: expr,
	})
}

func (p *logQL) startSummarizeByUnnamedGroupExpression(text string) {
	if p.SummarizeOp.ByGroupExpressions == nil {
		p.SummarizeOp.ByGroupExpressions = new(typesv1.SummarizeOperator_ByGroupExpressions)
	}
	p.SummarizeOp.ByGroupExpressions.Groups = append(p.SummarizeOp.ByGroupExpressions.Groups,
		&typesv1.SummarizeOperator_ByGroupExpression{
			Column: &typesv1.Identifier{Name: text},
		},
	)
}

func (p *logQL) endSummarizeByUnnamedGroupExpression(expr *typesv1.Expr) {
	top := len(p.SummarizeOp.ByGroupExpressions.Groups) - 1
	p.SummarizeOp.ByGroupExpressions.Groups[top].Scalar = expr
}

func (p *logQL) startProjectOp() {
	p.ProjectOp = &typesv1.ProjectOperator{}
}

func (p *logQL) startProjectOpArg(id string) {
	p.ProjectOp.Projections = append(p.ProjectOp.Projections, &typesv1.ProjectOperator_Projection{
		Column: &typesv1.Identifier{Name: id},
	})
}

func (p *logQL) setProjectOpArgValue(e *typesv1.Expr) {
	p.ProjectOp.Projections[len(p.ProjectOp.Projections)-1].Value = e
}

func (p *logQL) startProjectAwayOp() {
	p.ProjectAwayOp = &typesv1.ProjectAwayOperator{}
}

func (p *logQL) addProjectAwayOpArg(id string) {
	p.ProjectAwayOp.Projections = append(p.ProjectAwayOp.Projections, &typesv1.ProjectAwayOperator_Projection{
		Column: &typesv1.Identifier{Name: id},
	})
}

func (p *logQL) startProjectKeepOp() {
	p.ProjectKeepOp = &typesv1.ProjectKeepOperator{}
}

func (p *logQL) addProjectKeepOpArg(id string) {
	p.ProjectKeepOp.Projections = append(p.ProjectKeepOp.Projections, &typesv1.ProjectKeepOperator_Projection{
		Column: &typesv1.Identifier{Name: id},
	})
}

func (p *logQL) startExtendOp() {
	p.ExtendOp = &typesv1.ExtendOperator{}
}

func (p *logQL) setExtendOpArgColumnName(name string) {
	p.ExtendOp.Projections = append(p.ExtendOp.Projections, &typesv1.ExtendOperator_Projection{
		Column: &typesv1.Identifier{Name: name},
	})
}

func (p *logQL) setExtendOpArgValue(e *typesv1.Expr) {
	p.ExtendOp.Projections[len(p.ExtendOp.Projections)-1].Value = e
}

func (p *logQL) startCountOp() {
	p.CountOp = &typesv1.CountOperator{}
}

func (p *logQL) startDistinctOp() {
	p.DistinctOp = &typesv1.DistinctOperator{}
}

func (p *logQL) startSampleOp() {
	p.SampleOp = &typesv1.SampleOperator{}
}

func (p *logQL) setSampleOpCount(v int64) {
	p.SampleOp.Count = v
}

func (p *logQL) startSearchOp() {
	p.SearchOp = &typesv1.SearchOperator{}
}

func (p *logQL) setSearchOpKindDefault() {
	v := typesv1.SearchOperator_Default
	p.SearchOp.Kind = &v
}

func (p *logQL) setSearchOpKindCaseInsensitive() {
	v := typesv1.SearchOperator_CaseInsensitive
	p.SearchOp.Kind = &v
}

func (p *logQL) setSearchOpKindCaseSensitive() {
	v := typesv1.SearchOperator_CaseSensitive
	p.SearchOp.Kind = &v
}

func (p *logQL) setSearchOpPredicateLiteral(literal string) {
	p.SearchOp.Predicate = &typesv1.SearchOperator_Literal_{
		Literal: literal,
	}
}

func (p *logQL) setSearchOpPredicateFieldSearch(column, literal string) {
	p.SearchOp.Predicate = &typesv1.SearchOperator_Field{
		Field: &typesv1.SearchOperator_FieldSearch{
			Column:  column,
			Literal: literal,
		},
	}
}

func (p *logQL) setSearchOpPredicateExactSearch(column, literal string) {
	p.SearchOp.Predicate = &typesv1.SearchOperator_Exact{
		Exact: &typesv1.SearchOperator_ExactSearch{
			Column:  column,
			Literal: literal,
		},
	}
}

func (p *logQL) setSearchOpPredicateRegexSearch(column, regex string) {
	p.SearchOp.Predicate = &typesv1.SearchOperator_Regex{
		Regex: &typesv1.SearchOperator_RegexSearch{
			Column: column,
			Regex:  regex,
		},
	}
}

func (p *logQL) startSortOp() {
	p.SortOp = &typesv1.SortOperator{}
}

func (p *logQL) startSortOpArg(column string) {
	p.SortOp.ByColumns = append(p.SortOp.ByColumns, &typesv1.SortOperator_ByColumn{
		Column: &typesv1.Identifier{Name: column},
	})
}

func (p *logQL) setSortOpArgOrderAsc() {
	v := typesv1.SortOperator_Asc
	top := len(p.SortOp.ByColumns) - 1
	p.SortOp.ByColumns[top].Order = &v
}

func (p *logQL) setSortOpArgOrderDesc() {
	v := typesv1.SortOperator_Desc
	top := len(p.SortOp.ByColumns) - 1
	p.SortOp.ByColumns[top].Order = &v
}

func (p *logQL) startTakeOp() {
	p.TakeOp = &typesv1.TakeOperator{}
}

func (p *logQL) setTakeOpCount(v int64) {
	p.TakeOp.Count = v
}

func (p *logQL) startTopOp() {
	p.TopOp = &typesv1.TopOperator{}
}

func (p *logQL) setTopOpCount(v int64) {
	p.TopOp.Count = v
}

func (p *logQL) setTopOpByColumnScalar(e *typesv1.Expr) {
	p.TopOp.ByColumn = &typesv1.TopOperator_ByColumn{
		Scalar: e,
	}
}

func (p *logQL) setTopOpByColumnOrderAsc() {
	v := typesv1.TopOperator_Asc
	p.TopOp.ByColumn.Order = &v
}

func (p *logQL) setTopOpByColumnOrderDesc() {
	v := typesv1.TopOperator_Desc
	p.TopOp.ByColumn.Order = &v
}

func (p *logQL) startRenderSplitOp() {
	p.SplitByOp = &typesv1.SplitOperator{By: &typesv1.SplitOperator_ByOperator{}}
}

func (p *logQL) addRenderSplitByOp(e *typesv1.Expr) {
	p.SplitByOp.By.Scalars = append(p.SplitByOp.By.Scalars, e)
}

func (p *logQL) parseDoubleQuoteString(text string) string {
	v, err := strconv.Unquote(text)
	if err != nil {
		p.err = err
		panic(err)
	}
	return v
}

func (p *logQL) parseSingleQuoteString(text string) string {
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
	case "d":
		return 24 * time.Duration(v*float64(time.Hour))
	case "w":
		return 7 * 24 * time.Duration(v*float64(time.Hour))
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
	case "d":
		return time.Duration(v) * time.Hour * 24
	case "w":
		return time.Duration(v) * time.Hour * 24 * 7
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

func (p *logQL) popFunc() *typesv1.FuncCall {
	cur := len(p.FuncCalls) - 1
	pop := p.FuncCalls[cur]
	p.FuncCalls = p.FuncCalls[:cur]
	return pop
}
