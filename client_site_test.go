package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSiteRequest(t *testing.T) {
	r, err := NewCreateSiteRequest("https://example.org/adzerk/api/", CreateSiteJSONRequestBody{Title: "Title", URL: "https://example.org"})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"POST",
		"https://example.org/adzerk/api/v1/site",
		"{\"Title\":\"Title\",\"URL\":\"https://example.org\"}",
	)
}

func TestGetSiteRequest(t *testing.T) {
	r, err := NewGetSiteRequest("https://example.org/adzerk/api/", 123)
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithoutBody(t, r,
		"GET",
		"https://example.org/adzerk/api/v1/site/123",
	)
}

func TestUpdateSiteRequest(t *testing.T) {
	r, err := NewUpdateSiteRequest("https://example.org/adzerk/api/", 123, UpdateSiteJSONRequestBody{Id: 123, Title: "Title", URL: "https://example.org"})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"PUT",
		"https://example.org/adzerk/api/v1/site/123",
		"{\"Id\":123,\"Title\":\"Title\",\"URL\":\"https://example.org\"}",
	)
}

func TestDeleteSiteRequest(t *testing.T) {
	isDeleted := true
	r, err := NewUpdateSiteRequest("https://example.org/adzerk/api/", 123, UpdateSiteJSONRequestBody{Id: 123, Title: "Title", URL: "https://example.org", IsDeleted: &isDeleted})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"PUT",
		"https://example.org/adzerk/api/v1/site/123",
		"{\"Id\":123,\"IsDeleted\":true,\"Title\":\"Title\",\"URL\":\"https://example.org\"}",
	)
}
