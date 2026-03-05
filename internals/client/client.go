package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// through an interface we define what the client can do
type APIClient interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetPosts(ctx context.Context) ([]models.Post, error)
}

type HTTPClient struct {
	baseUrl string
	httpClient *htttp.Client
	tocken string

}

func NewHTTPClient(baseUrl string, token string) *HTTPClient {
	return &HTTPClient{
		base.Url: baseUrl,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		tocken: token,
	}
}

func (c *HTTPClient) GetUsers(ctx context.Context) ([]models.User, error){
	var users []models.User
	err := c.doRequest(ctx, "/users", &users)
	if err != nil {
		return nil. ftm.Errorf("get users: %w", err)
	}

	return users, nil
}