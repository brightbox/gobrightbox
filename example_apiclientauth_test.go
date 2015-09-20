package brightbox_test

import (
	"fmt"
	"github.com/brightbox/gobrightbox"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Authenticate using a api client identifier and secret
func Example() {
	apiUrl := "https://api.gb1.brightbox.com"
	applicationId := "cli-xxxxx"
	applicationSecret := "somesecret"

	// Setup OAuth2 authentication
	conf := clientcredentials.Config{
		ClientID:     applicationId,
		ClientSecret: applicationSecret,
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
	for _, server := range *servers {
		fmt.Printf("id:%s name:%s\n", server.Id, server.Name)
	}
}
