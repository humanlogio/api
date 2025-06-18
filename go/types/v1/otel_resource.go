package typesv1

func NewResource(schemaURL string, kvs []*KV) *Resource {
	h64 := fpString(schemaURL)
	h64 ^= Hash64KeyValues_orderDoesntMatter(kvs)
	return &Resource{
		ResourceHash_64: h64,
		SchemaUrl:       schemaURL,
		Attributes:      kvs,
	}
}
