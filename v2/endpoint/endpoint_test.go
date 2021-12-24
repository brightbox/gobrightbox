package endpoint

import "testing"

func TestDefaultEndpoint(t *testing.T) {
	testConfig := &Config{}
	endpoint, err := testConfig.Endpoint()
	if err != nil {
		t.Errorf("Unexpected error raised")
	}
	if *endpoint != Brightbox {
		t.Errorf("Generated Endpoint from blank config should match default endpoint")
		t.Errorf("%v", *endpoint)
	}
}
