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
func ExampleConnect_api_client_auth() {
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

	// Get a list of configMaps
	configMaps, err := client.ConfigMaps(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, configMap := range configMaps {
		fmt.Printf("id:%s name:%s\n", configMap.ID, configMap.Name)
	}

	// Create a new configMap
	name := "new_map"
	data := map[string]interface{}{
		"attribute": 42,
	}
	configMap, err := client.CreateConfigMap(
		ctx,
		brightbox.ConfigMapOptions{
			Name: &name,
			Data: &data,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id:%s name:%s\n", configMap.ID, configMap.Name)
}
