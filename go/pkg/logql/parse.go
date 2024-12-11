package logql

import (
	typesv1 "github.com/humanlogio/api/go/types/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func ParseLogQuery(q string) (*typesv1.LogQuery, error) {
	return &typesv1.LogQuery{}, nil
}

func ParseJSONLogQuery(q string) (*typesv1.LogQuery, error) {
	out := &typesv1.LogQuery{}
	return out, protojson.Unmarshal([]byte(q), out)
}
