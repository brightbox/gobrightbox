/*
Package passwordcredentials implements the API Resource Owner Password
Credentials access method.

This access method uses a user ID and password that can access several
accounts.

Accounts are selected within the [endpoint.Config].

	conf := &passwordcredentials.Config{
		UserName: userName,
		Password: userPassword,
		ID:       applicationID,
		Secret:   applicationSecret,
		Config: endpoint.Config{
			Account: accountID,
			Scopes:  endpoint.InfrastructureScope,
		},
	}

The application ID and secret are obtained by creating an [Oauth application ID] for your program.

[Oauth application ID]: https://www.brightbox.com/docs/guides/manager/oauth-applications/

*/
package passwordcredentials

import (
	"context"
	"net/http"

	"github.com/brightbox/gobrightbox/v2/endpoint"
	"golang.org/x/oauth2"
)

// Config includes the items necessary to authenticate using password
// credentials.
//
type Config struct {
	UserName string // The email address used to sign up to Brightbox
	Password string // Password, password and 2fa code, or temporary access token
	ID       string // Oauth application ID
	Secret   string // Oauth application secret
	endpoint.Config
}

// Client implements the [brightbox.Oauth2] access interface.
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
