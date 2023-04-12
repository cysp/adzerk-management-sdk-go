package client

import (
	"context"
)

type apiKeySecuritySource struct {
	apiKey string
}

func NewApiKeySecuritySource(apiKey string) SecuritySource {
	return apiKeySecuritySource{
		apiKey: apiKey,
	}
}

func (s apiKeySecuritySource) ApiKeyAuth(ctx context.Context, operationName string) (ApiKeyAuth, error) {
	return ApiKeyAuth{
		APIKey: s.apiKey,
	}, nil
}
