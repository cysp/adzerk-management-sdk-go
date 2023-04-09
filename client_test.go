package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type httpRequestDoer struct {
	do func(req *http.Request) (*http.Response, error)
}

func (h *httpRequestDoer) Do(req *http.Request) (*http.Response, error) {
	return h.do(req)
}

func TestNewClient(t *testing.T) {
	c, err := NewClientWithResponses("https://example.org/adzerk/api/", WithHTTPClient(&httpRequestDoer{do: func(req *http.Request) (*http.Response, error) {
		return nil, nil
	}}))
	require.NoError(t, err)

	assert.NotNil(t, c)
}
