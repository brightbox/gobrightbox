// Package passwordcredentials implements the API Resource Owner Password
// Credentials access method.
//
// This access methods uses a user ID and password that can access several accounts
package passwordcredentials

import (
	"context"
	"net/http"

	"github.com/brightbox/gobrightbox/v2/endpoint"
	"golang.org/x/oauth2"
)

// Config includes the items necessary to authenticate using password credentials
// Set the Account entry to determine which account you wish to access
type Config struct {
	UserName string
	Password string
	ID       string
	Secret   string
	endpoint.Config
}

// Client creates an oauth2 password credential client from the config
func (c *Config) Client(ctx context.Context) (*http.Client, oauth2.TokenSource, error) {
	endpoint, err := c.Endpoint()
	if err != nil {
		return nil, nil, err
	}
	conf := oauth2.Config{
		ClientID:     c.ID,
		ClientSecret: c.Secret,
		Scopes:       c.Scopes,
		Endpoint:     *endpoint,
	}
	token, err := conf.PasswordCredentialsToken(ctx, c.UserName, c.Password)
	if err != nil {
		return nil, nil, err
	}
	ts := conf.TokenSource(ctx, token)
	return oauth2.NewClient(ctx, ts), ts, nil
}
