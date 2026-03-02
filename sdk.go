package aerostack

import (
	"context"
	"net/http"
)

// SDK provides a clean, ergonomic facade over the generated Aerostack API client.
// It supports both Public Key (client-side) and Secret Key (server-side) operation.
type SDK struct {
	Database       *DatabaseFacade
	Authentication *AuthenticationAPIService
	Cache          *CacheAPIService
	Queue          *QueueAPIService
	Storage        *StorageAPIService
	AI             *AIAPIService
	Services       *ServicesAPIService
	Gateway        *GatewayAPIService
	Realtime       *Realtime

	cfg    *Configuration
	apiKey string
}

// DatabaseFacade provides ergonomic access to database operations.
type DatabaseFacade struct {
	api *DatabaseAPIService
}

// DbQuery is an ergonomic wrapper for the generated DbQuery operation.
func (f *DatabaseFacade) DbQuery(ctx context.Context, req DbQueryRequest) (*DbQueryResult, *http.Response, error) {
	return f.api.DbQuery(ctx).DbQueryRequest(req).Execute()
}

// SDKOption is a function that configures an SDK instance.
type SDKOption func(*Configuration)

// WithServerURL sets the base URL for the API.
func WithServerURL(url string) SDKOption {
	return func(cfg *Configuration) {
		cfg.Servers[0].URL = url
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) SDKOption {
	return func(cfg *Configuration) {
		cfg.HTTPClient = client
	}
}

// New creates a new SDK instance.
// Initialize with a Public Key for client-side environments or a Secret Key for server-side.
func New(apiKey string, opts ...SDKOption) *SDK {
	cfg := NewConfiguration()
	for _, opt := range opts {
		opt(cfg)
	}

	client := NewAPIClient(cfg)
	sdk := &SDK{
		Database:       &DatabaseFacade{api: client.DatabaseAPI},
		Authentication: client.AuthenticationAPI,
		Cache:          client.CacheAPI,
		Queue:          client.QueueAPI,
		Storage:        client.StorageAPI,
		AI:             client.AIAPI,
		Services:       client.ServicesAPI,
		Gateway:        client.GatewayAPI,
		Realtime:       newRealtime(cfg),
		cfg:            cfg,
		apiKey:         apiKey,
	}

	return sdk
}

// Context returns a new context with the SDK's default API key attached (or a provided one).
func (s *SDK) Context(ctx context.Context, optionalApiKey ...string) context.Context {
	key := s.apiKey
	if len(optionalApiKey) > 0 {
		key = optionalApiKey[0]
	}
	return context.WithValue(ctx, ContextAPIKeys, map[string]APIKey{
		"ApiKeyAuth": {Key: key},
	})
}
