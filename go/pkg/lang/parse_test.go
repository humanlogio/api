package lang

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/humanlogio/api/go/types/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func jsontime(t time.Time) string {
	v, err := t.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(v)
}

func jsonval(a any) string {
	v, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(v)
}

func TestParseJSONLogQuery(t *testing.T) {
	now := time.Date(2024, 12, 9, 16, 02, 41, 0, time.UTC)

	tests := []struct {
		name string
		in   string
		want *v1.LogQuery
	}{
		{
			name: "basic",
			in: `{
				"from": ` + jsontime(now.Add(-time.Hour)) + `,
				"context":{
					"machine_id":{
						"binary":{
							"lhs":{
								"identifier":{
									"name":"machine_id"
								}
							},
							"op": ` + jsonval(v1.BinaryOp_SET_IN) + `,
							"rhs":{
								"literal":{
									"arr": {"items":[{"i64":1},{"i64":2},{"i64":3}]}
								}
							}
						}
					}
				},
				"query": {
					"pipe": {}
				}
			}`,
			want: &v1.LogQuery{
				From: timestamppb.New(now.Add(-time.Hour)),
				Context: &v1.Context{
					MachineId: v1.ExprBinary(
						v1.ExprIdentifier("machine_id"),
						v1.BinaryOp_SET_IN,
						v1.ExprLiteral(v1.ValArr(
							v1.ValI64(1),
							v1.ValI64(2),
							v1.ValI64(3),
						)),
					),
				},
				Query: v1.ExprPipe(nil, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSONLogQuery(tt.in)
			require.NoError(t, err)
			require.Empty(t, cmp.Diff(tt.want, got, protocmp.Transform()))
		})
	}
}
