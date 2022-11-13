package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getVersion(t *testing.T) {
	// Nothing we can test other than that it does not panic without a built executable
	require.NotEmpty(t, getVersion())
	t.Log(getVersion())
}
