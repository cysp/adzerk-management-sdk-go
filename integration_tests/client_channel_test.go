package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/cysp/adzerk-management-sdk-go"
	. "github.com/cysp/adzerk-management-sdk-go/testserver"
)

func TestChannelIntegration(t *testing.T) {
	s := NewHttpTestServer()
	defer s.Close()

	t.Run("Create", func(t *testing.T) {
		testCreateChannel(t, s.URL)
	})
	t.Run("Get", func(t *testing.T) {
		testGetChannel(t, s.URL)
	})
	t.Run("Update", func(t *testing.T) {
		testUpdateChannel(t, s.URL)
	})
	t.Run("Delete", func(t *testing.T) {
		testDeleteChannel(t, s.URL)
	})
}

func testCreateChannel(t *testing.T, server string) {
	c, err := NewClientWithResponses(server)
	require.NoError(t, err)

	r, err := c.CreateChannelWithResponse(context.Background(), CreateChannelJSONRequestBody{Title: "Title", AdTypes: []int32{}})
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, 200, r.StatusCode())
	if assert.NotNil(t, r.JSON200) {
		assert.EqualValues(t, 100001, r.JSON200.Id)
		assert.EqualValues(t, "Title", r.JSON200.Title)
	}
}

func testGetChannel(t *testing.T, server string) {
	c, err := NewClientWithResponses(server)
	require.NoError(t, err)

	r, err := c.GetChannelWithResponse(context.Background(), 100001)
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, 200, r.StatusCode())
	if assert.NotNil(t, r.JSON200) {
		assert.EqualValues(t, 100001, r.JSON200.Id)
	}
}

func testUpdateChannel(t *testing.T, server string) {
	c, err := NewClientWithResponses(server)
	require.NoError(t, err)

	r, err := c.UpdateChannelWithResponse(context.Background(), 100001, UpdateChannelJSONRequestBody{Id: 100001, Title: "Title", Engine: 0})
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, 200, r.StatusCode())
	if assert.NotNil(t, r.JSON200) {
		assert.EqualValues(t, 100001, r.JSON200.Id)
		assert.EqualValues(t, "Title", r.JSON200.Title)
	}
}

func testDeleteChannel(t *testing.T, server string) {
	c, err := NewClientWithResponses(server)
	require.NoError(t, err)

	r, err := c.DeleteChannelWithResponse(context.Background(), 100001)
	require.NoError(t, err)
	require.NotNil(t, r)

	assert.EqualValues(t, 200, r.StatusCode())
}
