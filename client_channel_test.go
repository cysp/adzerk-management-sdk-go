package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateChannelRequest(t *testing.T) {
	r, err := NewCreateChannelRequest("https://example.org/adzerk/api/", CreateChannelJSONRequestBody{Title: "Title", AdTypes: []int32{}})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"POST",
		"https://example.org/adzerk/api/v1/channel",
		"{\"AdTypes\":[],\"Engine\":0,\"Title\":\"Title\"}",
	)
}

func TestGetChannelRequest(t *testing.T) {
	r, err := NewGetChannelRequest("https://example.org/adzerk/api/", 123)
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithoutBody(t, r,
		"GET",
		"https://example.org/adzerk/api/v1/channel/123",
	)
}

func TestUpdateChannelRequest(t *testing.T) {
	r, err := NewUpdateChannelRequest("https://example.org/adzerk/api/", 123, UpdateChannelJSONRequestBody{Id: 123, Title: "Title", AdTypes: []int32{}})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"PUT",
		"https://example.org/adzerk/api/v1/channel/123",
		"{\"AdTypes\":[],\"Engine\":0,\"Id\":123,\"Title\":\"Title\"}",
	)
}

func TestDeleteChannelRequest(t *testing.T) {
	r, err := NewDeleteChannelRequest("https://example.org/adzerk/api/", 123)
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithoutBody(t, r,
		"GET",
		"https://example.org/adzerk/api/v1/channel/123/delete",
	)
}
