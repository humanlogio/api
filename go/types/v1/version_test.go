package typesv1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionAsSemver(t *testing.T) {
	v1 := Version{Major: 1, Minor: 2}
	v2 := Version{Major: 1, Minor: 3}

	v1sem, err := v1.AsSemver()
	require.NoError(t, err)
	v2sem, err := v2.AsSemver()
	require.NoError(t, err)
	require.Equal(t, -1, v1sem.Compare(v2sem))
}
