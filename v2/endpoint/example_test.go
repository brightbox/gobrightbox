package endpoint_test

import (
	"fmt"

	"github.com/brightbox/gobrightbox/v2/endpoint"
)

func ExampleAPIURL() {
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

func ExampleAPIURL_defaults() {
	testConfig := &endpoint.Config{}
	u, err := testConfig.APIURL()
	if err == nil {
		fmt.Println(u.String())
	}
	// Output: https://api.gb1.brightbox.com/1.0/
}

func ExampleAPIURL_account() {
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

func ExampleTokenURL() {
	testConfig := &endpoint.Config{}
	url, err := testConfig.TokenURL()
	if err == nil {
		fmt.Println(url)
	}
	// Output: https://api.gb1.brightbox.com/token/
}

func ExampleStorageURL() {
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

func ExampleStorageURL_defaults() {
	testConfig := &endpoint.Config{}
	u, err := testConfig.StorageURL()
	if err == nil {
		fmt.Println(u)
	}
	// Output: https://orbit.brightbox.com/v1/
}

func ExampleStorageURL_account() {
	testConfig := &endpoint.Config{
		Account: "acc-testy",
	}
	u, err := testConfig.StorageURL()
	if err == nil {
		fmt.Println(u)
	}
	// Output: https://orbit.brightbox.com/v1/acc-testy/
}
