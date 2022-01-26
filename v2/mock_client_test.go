package brightbox

import (
	"context"
	"net/http"
	"net/url"
	
	"github.com/brightbox/gobrightbox/v2/endpoint"
)

type MockAuth struct {
	url string
}

func (a *MockAuth) APIURL() (*url.URL, error) {
	conf := endpoint.Config{
		BaseURL: a.url,
	}
	return conf.APIURL()
}

// HTTPClient is the context key to use with golang.org/x/net/context's
// WithValue function to associate an *http.Client value with a context.
var HTTPClient contextKey

// contextKey is just an empty struct. It exists so HTTPClient can be
// an immutable public variable with a unique type. It's immutable
// because nobody else can create a ContextKey, being unexported.
type contextKey struct{}

// Client creates a new http Client from context
func (_a *MockAuth) Client(ctx context.Context) (*http.Client, error) {
	if ctx != nil {
		if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok {
			return hc, nil
		}
	}
	return http.DefaultClient, nil
}
