package typesv1

func NewScope(schemaURL, name, version string, kvs []*KV) *Scope {
	h64 := fpString(schemaURL)
	h64 ^= fpString(name)
	h64 ^= fpString(version)
	h64 ^= Hash64KeyValues_orderDoesntMatter(kvs)
	return &Scope{
		ScopeHash_64: h64,
		SchemaUrl:    schemaURL,
		Name:         name,
		Version:      version,
		Attributes:   kvs,
	}
}
