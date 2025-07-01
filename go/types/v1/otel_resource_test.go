package typesv1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceH64(t *testing.T) {
	tests := []struct {
		name string
		kvs  []*KV
		want uint64
	}{
		{
			name: "simple order forward",
			kvs: []*KV{
				{Key: "hello", Value: ValStr("world")},
				{Key: "bonjour", Value: ValStr("le monde")},
				{Key: "salut", Value: ValNull()},
				{Key: "bye", Value: ValBool(false)},
			},
			want: 0x409610f4f1bf80f3,
		},
		{
			name: "simple order backward",
			kvs: []*KV{
				{Key: "bye", Value: ValBool(false)},
				{Key: "salut", Value: ValNull()},
				{Key: "bonjour", Value: ValStr("le monde")},
				{Key: "hello", Value: ValStr("world")},
			},
			want: 0x409610f4f1bf80f3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hash64KeyValues_orderDoesntMatter(tt.kvs)
			require.Equal(t, tt.want, got)
		})
	}
}
