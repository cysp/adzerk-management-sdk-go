package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateChannelSiteMapRequest(t *testing.T) {
	r, err := NewCreateChannelSiteMapRequest("https://example.org/adzerk/api/", CreateChannelSiteMapJSONRequestBody{
		ChannelId: 123,
		SiteId:    234,
		Priority:  10,
	})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"POST",
		"https://example.org/adzerk/api/v1/channelSite",
		"{\"channelId\":123,\"priority\":10,\"siteId\":234}",
	)
}

func TestGetChannelSiteMapRequest(t *testing.T) {
	r, err := NewGetChannelSiteMapRequest("https://example.org/adzerk/api/", 123, 234)
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithoutBody(t, r,
		"GET",
		"https://example.org/adzerk/api/v1/channel/123/site/234",
	)
}

func TestUpdateChannelSiteMapRequest(t *testing.T) {
	r, err := NewUpdateChannelSiteMapRequest("https://example.org/adzerk/api/", UpdateChannelSiteMapJSONRequestBody{
		ChannelId: 123,
		SiteId:    234,
		Priority:  10,
	})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"PUT",
		"https://example.org/adzerk/api/v1/channelSite",
		"{\"channelId\":123,\"priority\":10,\"siteId\":234}",
	)
}

func TestDeleteChannelSiteMapRequest(t *testing.T) {
	r, err := NewDeleteChannelSiteMapRequest("https://example.org/adzerk/api/", 123, 234)
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithoutBody(t, r,
		"GET",
		"https://example.org/adzerk/api/v1/channel/123/site/234/delete",
	)
}
