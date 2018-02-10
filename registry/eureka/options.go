package eureka

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"micro/go-micro/registry"
)

type contextHttpClient struct{}

var newOAuthClient = func(c clientcredentials.Config) *http.Client {
	return c.Client(oauth2.NoContext)
}

// Enable OAuth 2.0 Client Credentials Grant Flow
func OAuth2ClientCredentials(clientID, clientSecret, tokenURL string) registry.Option {
	return func(o *registry.Options) {
		c := clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     tokenURL,
		}

		o.Context = context.WithValue(o.Context, contextHttpClient{}, newOAuthClient(c))
	}
}
