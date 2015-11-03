package brightbox_test

import (
	"fmt"
	"github.com/brightbox/gobrightbox"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Authenticate using an API Client identifier and secret, and get a list of
// servers
func Example() {
	apiUrl := "https://api.gb1.brightbox.com"
	clientId := "cli-xxxxx"
	clientSecret := "somesecret"

	// Setup OAuth2 authentication
	conf := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{},
		TokenURL:     apiUrl + "/token",
	}
	oc := conf.Client(oauth2.NoContext)

	// Setup connection to API
	client, err := brightbox.NewClient(apiUrl, "", oc)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get a list of servers
	servers, err := client.Servers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, server := range servers {
		fmt.Printf("id:%s name:%s\n", server.Id, server.Name)
	}
}
