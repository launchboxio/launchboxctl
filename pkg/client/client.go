package client

import (
	"crypto/tls"
	"github.com/dghubble/sling"
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

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type SoftDeletable struct {
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
