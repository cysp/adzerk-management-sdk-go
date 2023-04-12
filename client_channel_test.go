package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateChannel(t *testing.T) {
	s := NewHttpTestServer()

	c, err := NewClient(s.URL, NewApiKeySecuritySource("api-key"))
	require.NoError(t, err)

	r, err := c.CreateChannel(context.Background(), NewOptCreateChannelReq(CreateChannelReq{
		Title:   "Title",
		AdTypes: []int32{},
	}))
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.ID.Value, 123)
	assert.EqualValues(t, r.Title.Value, "Title")
}

func TestGetChannel(t *testing.T) {
	s := NewHttpTestServer()

	c, err := NewClient(s.URL, NewApiKeySecuritySource("api-key"))
	require.NoError(t, err)

	r, err := c.GetChannel(context.Background(), GetChannelParams{ID: 123})
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.ID.Value, 123)
}

func TestUpdateChannel(t *testing.T) {
	s := NewHttpTestServer()

	c, err := NewClient(s.URL, NewApiKeySecuritySource("api-key"))
	require.NoError(t, err)

	r, err := c.UpdateChannel(context.Background(), NewOptUpdateChannelReq(UpdateChannelReq{
		ID:      123,
		Title:   "Title",
		AdTypes: []int32{},
		Engine:  0,
	}), UpdateChannelParams{ID: 123})
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, r.ID.Value, 123)
	assert.EqualValues(t, r.Title.Value, "Title")
}
