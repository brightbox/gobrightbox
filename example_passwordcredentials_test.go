package brightbox_test

import (
	"fmt"
	"github.com/brightbox/gobrightbox"
	"golang.org/x/oauth2"
)

// Authenticate using OAuth2 password credentials
func ExamplePasswordCredentials() {
	apiUrl := "https://api.gb1.brightbox.com"
	// Brightbox username and password
	userName := "john@example.com"
	password := "mypassword"
	// Users can have multiple accounts, so you need to specify which one
	accountId := "acc-h3nbk"
	// These OAuth2 application credentials are public, distributed with the
	// cli.
	applicationId := "app-12345"
	applicationSecret := "mocbuipbiaa6k6c"

	// Setup OAuth2 authentication.
	conf := oauth2.Config{
		ClientID:     applicationId,
		ClientSecret: applicationSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: apiUrl + "/token",
		},
	}
	token, err := conf.PasswordCredentialsToken(oauth2.NoContext, userName, password)
	if err != nil {
		fmt.Println(err)
	}
	oc := oauth2.NewClient(oauth2.NoContext, conf.TokenSource(oauth2.NoContext, token))

	// Setup connection to API
	client, err := brightbox.NewClient(apiUrl, accountId, oc)
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
