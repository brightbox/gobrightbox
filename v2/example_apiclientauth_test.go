package brightbox_test

import (
	"context"
	"fmt"
	"log"

	brightbox "github.com/brightbox/gobrightbox/v2"
	"github.com/brightbox/gobrightbox/v2/clientcredentials"
)

// Authenticate using an API Client identifier and secret, and get a list of
// servers
func ExampleClientCredentials() {
	// Brightbox client details issued on a specific account
	clientID := "cli-xxxxx"
	clientSecret := "somesecret"

	// Setup Config
	conf := &clientcredentials.Config{
		ID:     clientID,
		Secret: clientSecret,
	}

	// Underlying network connection context.
	ctx := context.Background()

	// Setup connection to API
	client, err := brightbox.Connect(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of servers
	serverQuerier := brightbox.NewQuerier[brightbox.Server](client)
	servers, err := serverQuerier.All()
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Printf("id:%s name:%s\n", server.ID, server.Name)
	}
}
