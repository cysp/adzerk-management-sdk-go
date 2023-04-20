package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {

	c, err := NewClientWithResponses("https://example.org/adzerk/api/")
	require.NoError(t, err)

	assert.NotNil(t, c)
}
