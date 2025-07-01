package typesv1

import (
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"
)

func (x *Log) IsStructured() bool {
	if x.Timestamp != nil || x.SeverityText != "" || x.Body != "" || x.Attributes != nil {
		return true
	}
	return false
}

func (x *Log) FindAttr(k string) *Val {
	if k == string(semconv.LogRecordUIDKey) {
		return ValStr(x.Ulid)
	}
	if k == string(semconv.LogRecordOriginalKey) {
		return ValBlob(x.Raw)
	}
	for _, kv := range x.Attributes {
		if kv.Key == k {
			return kv.Value
		}
	}
	return nil
}
