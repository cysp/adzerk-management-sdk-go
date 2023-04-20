package client

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertRequest(t *testing.T, r *http.Request, method string, url string) {
	t.Helper()

	assert.EqualValues(t, method, r.Method)
	assert.EqualValues(t, url, r.URL.String())
}

func assertRequestWithoutBody(t *testing.T, r *http.Request, method string, url string) {
	t.Helper()

	assertRequest(t, r, method, url)

	assert.Nil(t, r.Body)
}

func assertRequestWithStringBody(t *testing.T, r *http.Request, method string, url string, body string) {
	t.Helper()

	assertRequest(t, r, method, url)

	requestBody, err := io.ReadAll(r.Body)
	require.NoError(t, err)
	assert.EqualValues(t, body, string(requestBody))
}
