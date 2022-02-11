// Package gobrightbox is for interacting with the Brightbox Cloud API
//
// Brightbox Cloud is a UK-based infrastructure-as-a-service
// provider. More details available at https://www.brightbox.com
//
// The Brightbox Cloud API documentation is available at
// https://api.gb1.brightbox.com/1.0/
package gobrightbox

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/tomnomnom/linkheader"
)

const (
	// DefaultRegionAPIURL is the default API URL for the region. Use with NewClient.
	DefaultRegionAPIURL = "https://api.gb1.brightbox.com/"
	// DefaultOrbitAuthURL is the default Auth URL for Orbit.
	DefaultOrbitAuthURL = "https://orbit.brightbox.com/v1/"
)

// Client represents a connection to the Brightbox API. You should use NewClient
// to allocate and configure Clients. Authentication is handled externally by a
// http.Client with the appropriate Transport, such as those provided by
// https://github.com/golang/oauth2/
type Client struct {
	BaseURL   *url.URL
	client    *http.Client
	UserAgent string
	// The identifier of the account to use by default with this Client.
	AccountID string
}

// NewClient allocates and configures a Client for interacting with the API.
//
// apiURL should be an url of the form https://api.region.brightbox.com,
// e.g: https://api.gb1.brightbox.com. You can use the default defined in
// this package instead, i.e. brightbox.DefaultRegionAPIURL
//
// accountId should be the identifier of the default account to be used with
// this Client. Clients authenticated with Brightbox APIClient credentials are
// only ever associated with one single Account, so you can leave this empty for
// those. Client's authenticated with Brightbox User credentials can have access
// to multiple accounts, so this parameter should be provided.
//
// httpClient should be a http.Client with a transport that will handle the
// OAuth token authentication, such as those provided by
// https://github.com/golang/oauth2/
func NewClient(apiURL string, accountID string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	au, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:    httpClient,
		BaseURL:   au,
		AccountID: accountID,
	}
	return c, nil
}

// NewRequest allocates and configures a http.Request ready to make an API call.
//
// method should be the desired http method, e.g: "GET", "POST", "PUT" etc.
//
// urlStr should be the url path, relative to the api url e.g: "/1.0/servers"
//
// if body is non-nil, it will be Marshaled to JSON and set as the request body
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	if c.AccountID != "" {
		q := u.Query()
		q.Set("account_id", c.AccountID)
		u.RawQuery = q.Encode()
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

// MakeAPIRequest makes a http request to the API, JSON encoding any given data
// and decoding any JSON response.
//
// method should be the desired http method, e.g: "GET", "POST", "PUT" etc.
//
// urlStr should be the url path, relative to the api url e.g: "/1.0/servers"
//
// if reqBody is non-nil, it will be Marshaled to JSON and set as the request
// body.
//
// Optionally, the response body will be Unmarshaled from JSON into whatever
// resBody is a pointer to. Leave nil to skip.
//
// If the response is non-2xx, MakeAPIRequest will try to parse the error
// message and return an APIError struct.
func (c *Client) MakeAPIRequest(method string, path string, reqBody interface{}, resBody interface{}) (*http.Response, error) {
	req, err := c.NewRequest(method, path, reqBody)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		if resBody != nil {
			err := json.NewDecoder(res.Body).Decode(resBody)
			if err != nil {
				return res, &APIError{
					RequestURL: res.Request.URL,
					StatusCode: res.StatusCode,
					Status:     res.Status,
					ParseError: err,
				}
			}
		}
		return res, nil
	}
	apierr := &APIError{
		RequestURL: res.Request.URL,
		StatusCode: res.StatusCode,
		Status:     res.Status,
	}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &apierr)
	apierr.ResponseBody = body
	return res, apierr
}

func getLinkRel(header string, prefix string, rel string) *string {
	links := linkheader.Parse(header)
	re := regexp.MustCompile(prefix + "-[^/]+")
	for _, link := range links {
		id := re.FindString(link.URL)
		if id != "" && link.Rel == rel {
			return &id
		}
	}
	return nil
}
