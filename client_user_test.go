package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUserRequest(t *testing.T) {
	userName := "User Name"
	r, err := NewCreateUserRequest("https://example.org/adzerk/api/", CreateUserJSONRequestBody{
		Email: "example@example.org",
		Name:  &userName,
	})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"POST",
		"https://example.org/adzerk/api/v1/login",
		"{\"Email\":\"example@example.org\",\"Name\":\"User Name\"}",
	)
}

func TestGetUserRequest(t *testing.T) {
	r, err := NewGetUserRequest("https://example.org/adzerk/api/", 123)
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithoutBody(t, r,
		"GET",
		"https://example.org/adzerk/api/v1/login/123",
	)
}

func TestUpdateUserRequest(t *testing.T) {
	r, err := NewUpdateUserRequest("https://example.org/adzerk/api/", 123, UpdateUserJSONRequestBody{
		Id:    123,
		Email: "example@example.org",
		Name:  "User Name",
	})
	require.NoError(t, err)
	require.NotNil(t, r)

	assertRequestWithStringBody(t, r,
		"PUT",
		"https://example.org/adzerk/api/v1/login/123",
		"{\"Email\":\"example@example.org\",\"Name\":\"User Name\",\"id\":123}",
	)
}
