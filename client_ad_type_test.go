package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAdTypeRequest(t *testing.T) {
	r, err := NewCreateAdTypeRequest("https://example.org/adzerk/api/", CreateAdTypeJSONRequestBody{Width: 640, Height: 480})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"POST",
		"https://example.org/adzerk/api/v1/adtypes",
		"{\"Height\":480,\"Width\":640}",
	)
}

func TestListAdTypesRequest(t *testing.T) {
	r, err := NewListAdTypesRequest("https://example.org/adzerk/api/", &ListAdTypesParams{})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithoutBody(t, r,
		"GET",
		"https://example.org/adzerk/api/v1/adtypes",
	)
}

func TestDeleteAdTypeRequest(t *testing.T) {
	r, err := NewDeleteAdTypeRequest("https://example.org/adzerk/api/", 123)
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithoutBody(t, r,
		"GET",
		"https://example.org/adzerk/api/v1/adtypes/123/delete",
	)
}
