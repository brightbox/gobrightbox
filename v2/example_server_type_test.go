package brightbox_test

import (
	"context"
	"fmt"
	"log"

	brightbox "github.com/brightbox/gobrightbox/v2"
	"github.com/brightbox/gobrightbox/v2/clientcredentials"
)

// Authenticate using an API Client identifier and secret, and get a list of
// configMaps
func ExampleConnect_fourth() {
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

	// Get a server type by handle
	serverType, err := client.ServerTypeByHandle(ctx, "2gb.ssd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id:%s name:%s\n", serverType.ID, serverType.Name)
}
