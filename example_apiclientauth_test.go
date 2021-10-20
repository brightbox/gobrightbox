package gobrightbox_test

import (
	"fmt"

	brightbox "github.com/brightbox/gobrightbox"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Authenticate using an API Client identifier and secret, and get a list of
// servers
func Example() {
	apiURL := "https://api.gb1.brightbox.com"
	clientID := "cli-xxxxx"
	clientSecret := "somesecret"

	// Setup OAuth2 authentication
	conf := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{},
		TokenURL:     apiURL + "/token",
	}
	oc := conf.Client(oauth2.NoContext)

	// Setup connection to API
	client, err := brightbox.NewClient(apiURL, "", oc)
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
		fmt.Printf("id:%s name:%s\n", server.ID, server.Name)
	}
}
