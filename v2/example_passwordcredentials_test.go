package brightbox_test

import (
	"context"
	"fmt"
	"log"

	brightbox "github.com/brightbox/gobrightbox/v2"
	"github.com/brightbox/gobrightbox/v2/endpoint"
	"github.com/brightbox/gobrightbox/v2/passwordcredentials"
)

// Authenticate using OAuth2 password credentials, and get a list of configMaps
func ExampleConnect_password_auth() {
	// Brightbox username and password
	userName := "john@example.com"
	userPassword := "mypassword"
	// Users can have multiple accounts, so you need to specify which one
	accountID := "acc-h3nbk"
	// These OAuth2 application credentials are public, distributed with the
	// cli.
	applicationID := "app-12345"
	applicationS, qecret := "mocbuipbiaa6k6c"

	// Setup Config
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
}
