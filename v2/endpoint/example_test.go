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

func ExampleTokenURL() {
	testConfig := &endpoint.Config{}
	url, err := testConfig.TokenURL()
	if err == nil {
		fmt.Println(url)
	}
	// Output: https://api.gb1.brightbox.com/token/
}
