package client

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateChannelRequest(t *testing.T) {
	r, err := NewCreateChannelRequest("https://example.org/adzerk/api/", map[string]interface{}{"Title": "Title"})
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.Method, "POST")
	assert.EqualValues(t, r.URL.String(), "https://example.org/adzerk/api/v1/channel")
	body, err := io.ReadAll(r.Body)
	require.NoError(t, err)
	assert.EqualValues(t, body, "{\"Title\":\"Title\"}")
}

func TestCreateChannel(t *testing.T) {
	c, err := NewClientWithResponses("https://example.org/adzerk/api/", WithHTTPClient(&httpRequestDoer{do: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader("{\"Id\":123,\"Title\":\"Title\"}")),
		}, nil
	}}))
	require.NoError(t, err)

	r, err := c.CreateChannelWithResponse(context.Background(), map[string]interface{}{"Title": "Title"})
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.StatusCode(), 200)
	if assert.NotNil(t, r.JSON200) {
		assert.EqualValues(t, *r.JSON200.Id, 123)
		assert.EqualValues(t, *r.JSON200.Title, "Title")
	}
}

func TestGetChannelRequest(t *testing.T) {
	r, err := NewGetChannelRequest("https://example.org/adzerk/api/", 123)
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.Method, "GET")
	assert.EqualValues(t, r.URL.String(), "https://example.org/adzerk/api/v1/channel/123")
}

func TestGetChannel(t *testing.T) {
	c, err := NewClientWithResponses("https://example.org/adzerk/api/", WithHTTPClient(&httpRequestDoer{do: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader("{\"Id\":123}")),
		}, nil
	}}))
	require.NoError(t, err)

	r, err := c.GetChannelWithResponse(context.Background(), 123)
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.StatusCode(), 200)
	if assert.NotNil(t, r.JSON200) {
		assert.EqualValues(t, *r.JSON200.Id, 123)
	}
}

func TestUpdateChannelRequest(t *testing.T) {
	r, err := NewUpdateChannelRequest("https://example.org/adzerk/api/", 123, map[string]interface{}{"Id": 123, "Title": "Title"})
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.Method, "PUT")
	assert.EqualValues(t, r.URL.String(), "https://example.org/adzerk/api/v1/channel/123")
	body, err := io.ReadAll(r.Body)
	require.NoError(t, err)
	assert.EqualValues(t, body, "{\"Id\":123,\"Title\":\"Title\"}")
}

func TestUpdateChannel(t *testing.T) {
	c, err := NewClientWithResponses("https://example.org/adzerk/api/", WithHTTPClient(&httpRequestDoer{do: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       req.Body,
		}, nil
	}}))
	require.NoError(t, err)

	r, err := c.UpdateChannelWithResponse(context.Background(), 123, map[string]interface{}{"Id": 123, "Title": "Title"})
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.StatusCode(), 200)
	if assert.NotNil(t, r.JSON200) {
		assert.EqualValues(t, *r.JSON200.Id, 123)
		assert.EqualValues(t, *r.JSON200.Title, "Title")
	}
}
