package client

import (
	"crypto/tls"
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type ClientOpts struct {
	Host       string
	Token      string
	ApiVersion string
	Tls        *tls.Config
}

type Client struct {
	Handler *sling.Sling
}

type BaseError struct{}

func New(opts *ClientOpts) *Client {
	if opts.ApiVersion == "" {
		opts.ApiVersion = "v1"
	}
	sdk := sling.
		New().
		Base(opts.Host+"/api/"+opts.ApiVersion+"/").
		Set("Authorization", "Bearer "+opts.Token)

	return &Client{sdk}
}

func (c *Client) isApiError(resp *http.Response) error {
	switch resp.StatusCode {
	case 401:
		return NewUnauthorizedError()
	case 403:
		return NewForbiddenError()
	case 404:
		return NewNotFoundError()
	case 500:
		return NewInternalServerError()
	default:
		return nil
	}
}

type Timestamps struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type SoftDeletable struct {
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
