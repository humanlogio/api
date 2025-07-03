package typesv1

import semconv "go.opentelemetry.io/otel/semconv/v1.32.0"

func NewResource(schemaURL string, kvs []*KV) *Resource {
	h64 := fpString(schemaURL)
	h64 ^= Hash64KeyValues_orderDoesntMatter(kvs)
	return &Resource{
		ResourceHash_64: h64,
		SchemaUrl:       schemaURL,
		Attributes:      kvs,
	}
}

const svcNameKey = string(semconv.ServiceNameKey)

func (v *Resource) LookupServiceName() string {
	for _, kv := range v.Attributes {
		if svcNameKey == kv.Key {
			return kv.Value.GetStr()
		}
	}
	return ""
}
