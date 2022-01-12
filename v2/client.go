package brightbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client represents a connection to the Brightbox API. You should use NewConnect
// to allocate and configure Clients, and pass in either a
// clientcredentials or password configuration.
type Client struct {
	UserAgent string
	baseURL   *url.URL
	client    *http.Client
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
				return res, APIError{
					RequestURL: res.Request.URL,
					StatusCode: res.StatusCode,
					Status:     res.Status,
					ParseError: &err,
				}
			}
		}
		return res, nil
	}
	apierr := APIError{
		RequestURL: res.Request.URL,
		StatusCode: res.StatusCode,
		Status:     res.Status,
	}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &apierr)
	apierr.ResponseBody = body
	return res, apierr
}

// APIError can be returned when an API request fails. It provides any error
// messages provided by the API, along with other details about the response.
type APIError struct {
	// StatusCode will hold the HTTP status code from the request that errored
	StatusCode int
	// Status will hold the HTTP status line from the request that errored
	Status string
	// AuthError will hold any available OAuth "error" field contents. See
	// https://api.gb1.brightbox.com/1.0/#errors
	AuthError string `json:"error"`
	// AuthErrorDescription will hold any available OAuth "error_description"
	// field contents. See https://api.gb1.brightbox.com/1.0/#errors
	AuthErrorDescription string `json:"error_description"`
	// ErrorName will hold any available Brightbox API "error_name" field
	// contents. See https://api.gb1.brightbox.com/1.0/#request_errors
	ErrorName string `json:"error_name"`
	// Errors will hold any available Brightbox API "errors" field contents. See
	// https://api.gb1.brightbox.com/1.0/#request_errors
	Errors []string `json:"errors"`
	// ParseError will hold any errors from the JSON parser whilst parsing an
	// API response
	ParseError *error
	// RequestURL will hold the full URL used to make the request that errored,
	// if available
	RequestURL *url.URL
	// ResponseBody will hold the raw respose body of the request that errored,
	// if available
	ResponseBody []byte
}

func (e APIError) Error() string {
	var url string
	if e.RequestURL != nil {
		url = e.RequestURL.String()
	}
	if e.ParseError != nil {
		return fmt.Sprintf("%d: %s: %s", e.StatusCode, url, *e.ParseError)
	}

	var msg string
	if e.AuthError != "" {
		msg = fmt.Sprintf("%s, %s", e.AuthError, e.AuthErrorDescription)
	}
	if e.ErrorName != "" {
		msg = e.ErrorName
		if len(e.Errors) == 1 {
			msg = msg + ": " + e.Errors[0]
		} else if len(e.Errors) > 1 {
			msg = fmt.Sprintf("%s: %s", msg, e.Errors)
		}

	}
	if msg == "" {
		msg = fmt.Sprintf("%s: %s", e.Status, url)
	}
	return msg
}

// NewRequest allocates and configures a http.Request ready to make an API call.
//
// method should be the desired http method, e.g: "GET", "POST", "PUT" etc.
//
// urlStr should be the url path, relative to the api url e.g: "/1.0/servers"
//
// if body is non-nil, it will be Marshaled to JSON and set as the request body
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
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
