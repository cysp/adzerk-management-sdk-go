package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	s := NewHttpTestServer()
	c, err := NewClient(s.URL, NewApiKeySecuritySource(""))
	require.NoError(t, err)

	assert.NotNil(t, c)
}
