// Package clientcredentials implements the API client credentials
// access method.
//
// API client credentials are an identifier and secret issued to a
// single account to access the API, commonly used for authenticating
// automated systems.
package clientcredentials

import (
	"context"
	"net/http"

	"github.com/brightbox/gobrightbox/v2/endpoint"
	"golang.org/x/oauth2"
	oauth2cc "golang.org/x/oauth2/clientcredentials"
)

// Config includes the items necessary to authenticate using client
// credentials
type Config struct {
	ID     string
	Secret string
	endpoint.Config
}

// Client creates an oauth2 clientcredential client from the config.
func (c *Config) Client(ctx context.Context) (*http.Client, oauth2.TokenSource, error) {
	tokenURL, err := c.TokenURL()
	if err != nil {
		return nil, nil, err
	}

	conf := oauth2cc.Config{
		ClientID:     c.ID,
		ClientSecret: c.Secret,
		Scopes:       c.Scopes,
		TokenURL:     tokenURL,
	}
	ts := conf.TokenSource(ctx)
	return oauth2.NewClient(ctx, ts), ts, nil
}
