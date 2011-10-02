package brightbox

import (
	"testing"
	"os"
)

func TestNewApiClientAuth(t *testing.T) {
	c := NewApiClientAuth("auth", "cli-xxxxx", "asdf1234")
	if c.url != "auth" {
		t.Error("url not set correctly")
	}
	if c.id != "cli-xxxxx" {
		t.Error("id not set correcty")
	}
	if c.secret != "asdf1234" {
		t.Error("secret not set correctly")
	}
	if c.token != "" {
		t.Error("token was not default empty")
	}
}

func TestRequestTokenWithInvalidDetails(t *testing.T) {
	c := NewApiClientAuth("https://api.gb1.brightbox.com", "test", "asdf1234")
	token, err := RequestToken(c)
	if token != "" {
		t.Errorf("token should be empty")
	}
	if err == nil || err.String() != "Token not granted" {
		t.Errorf("err should be 'Token not granted'")
	}
}

func TestRequestToken(t *testing.T) {
	if os.Getenv("CLIENT") == "" || os.Getenv("SECRET") == "" {
		t.Fatal("Test requires CLIENT and SECRET env variables set")
	}
	c := NewApiClientAuth("https://api.gb1.brightbox.com", os.Getenv("CLIENT"), os.Getenv("SECRET"))
	token, err := RequestToken(c)
	if token == "" {
		t.Errorf("token should not be nil")
	}
	if err != nil {
		t.Errorf("err should be nil")
	}
}