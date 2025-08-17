package typesv1

func NewScope(schemaURL, name, version string, kvs []*KV) *Scope {
	h64 := Hash64String(schemaURL)
	h64 ^= Hash64String(name)
	h64 ^= Hash64String(version)
	h64 ^= Hash64KeyValues_orderDoesntMatter(kvs)
	return &Scope{
		ScopeHash_64: h64,
		SchemaUrl:    schemaURL,
		Name:         name,
		Version:      version,
		Attributes:   kvs,
	}
}
