package endpoint_test

import (
	"fmt"

	"github.com/brightbox/gobrightbox/v2/endpoint"
)

func ExampleConfig() {
	testConfig := &endpoint.Config{
		BaseURL: "https://api.gb2.brightbox.com",
		Version: "2.0",
	}
	u, err := testConfig.APIURL()
	if err == nil {
		fmt.Println(u.String())
	}
	// Output: https://api.gb2.brightbox.com/2.0/
}

func ExampleConfig_defaults() {
	testConfig := &endpoint.Config{}
	u, err := testConfig.APIURL()
	if err == nil {
		fmt.Println(u.String())
	}
	// Output: https://api.gb1.brightbox.com/1.0/
}

func ExampleConfig_account() {
	testConfig := &endpoint.Config{
		BaseURL: "https://api.gb2.brightbox.com",
		Version: "2.0",
		Account: "acc-testy",
	}
	u, err := testConfig.APIURL()
	if err == nil {
		fmt.Println(u.String())
	}
	// Output: https://api.gb2.brightbox.com/2.0/?account_id=acc-testy
}

func ExampleConfig_token() {
	testConfig := &endpoint.Config{}
	url, err := testConfig.TokenURL()
	if err == nil {
		fmt.Println(url)
	}
	// Output: https://api.gb1.brightbox.com/token/
}

func ExampleConfig_storage() {
	testConfig := &endpoint.Config{
		BaseURL: "https://files.gb2.brightbox.com",
		Version: "v2",
	}
	u, err := testConfig.StorageURL()
	if err == nil {
		fmt.Println(u)
	}
	// Output: https://files.gb2.brightbox.com/v2/
}

func ExampleConfig_storageDefaults() {
	testConfig := &endpoint.Config{}
	u, err := testConfig.StorageURL()
	if err == nil {
		fmt.Println(u)
	}
	// Output: https://orbit.brightbox.com/v1/
}

func ExampleConfig_storageAccount() {
	testConfig := &endpoint.Config{
		Account: "acc-testy",
	}
	u, err := testConfig.StorageURL()
	if err == nil {
		fmt.Println(u)
	}
	// Output: https://orbit.brightbox.com/v1/acc-testy/
}
