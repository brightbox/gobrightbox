package brightbox

import (
	"errors"
	"fmt"
	"bytes"
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
)

//go:generate ./generate_default_functions paths.yaml

// Client represents a connection to the Brightbox API. You should use NewConnect
// to allocate and configure Clients, and pass in either a
// clientcredentials or password configuration.
type Client struct {
	UserAgent      string
	baseURL        *url.URL
	client         *http.Client
	tokensource    oauth2.TokenSource
	hardcoreDecode bool
}

// ResourceBaseURL returns the base URL within the client
func (q *Client) ResourceBaseURL() *url.URL {
	return q.baseURL
}

// HTTPClient returns the current HTTP structure within the client
func (q *Client) HTTPClient() *http.Client {
	return q.client
}

// ExtractTokenID implements the AuthResult interface for gophercloud clients
func (q *Client) ExtractTokenID() (string, error) {
	token, err := q.tokensource.Token()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

// AllowUnknownFields stops the Client generating an error is an unsupported field is
// returned by the API.
func (q *Client) AllowUnknownFields() {
	q.hardcoreDecode = false
}

// DisallowUnknownFields causes the Client to generate an error if an unsupported field is
// returned by the API.
func (q *Client) DisallowUnknownFields() {
	q.hardcoreDecode = true
}

// APIGet makes a GET request to the API
// and decoding any JSON response.
//
// relUrl is the relative path of the endpoint to the base URL, e.g. "servers".
func APIGet[O any](
	ctx context.Context,
	q *Client,
	relUrl string,
) (*O, error) {
	return apiCommand[O](ctx, q, "GET", relUrl)
}

// APIGetCollection makes a GET request to the API
// and decoding any JSON response into an appropriate slice
//
// relUrl is the relative path of the endpoint to the base URL, e.g. "servers".
func APIGetCollection[O any](
	ctx context.Context,
	q *Client,
	relUrl string,
) ([]O, error) {
	collection, err := APIGet[[]O](ctx, q, relUrl)
	if collection == nil {
		return nil, err
	}
	return *collection, err
}

// APIPost makes a POST request to the API, JSON encoding any given data
// and decoding any JSON response.
//
// relUrl is the relative path of the endpoint to the base URL, e.g. "servers".
//
// if reqBody is non-nil, it will be Marshaled to JSON and set as the request
// body.
func APIPost[O any](
	ctx context.Context,
	q *Client,
	relUrl string,
	reqBody interface{},
) (*O, error) {
	return apiObject[O](ctx, q, "POST", relUrl, reqBody)
}

// APIPut makes a PUT request to the API, JSON encoding any given data
// and decoding any JSON response.
//
// relUrl is the relative path of the endpoint to the base URL, e.g. "servers".
//
// if reqBody is non-nil, it will be Marshaled to JSON and set as the request
// body.
func APIPut[O any](
	ctx context.Context,
	q *Client,
	relUrl string,
	reqBody interface{},
) (*O, error) {
	return apiObject[O](ctx, q, "PUT", relUrl, reqBody)
}

// APIDelete makes a DELETE request to the API
//
// relUrl is the relative path of the endpoint to the base URL, e.g. "servers".
func APIDelete[O any](
	ctx context.Context,
	q *Client,
	relUrl string,
) (*O, error) {
	return apiCommand[O](ctx, q, "DELETE", relUrl)
}

func apiObject[O any](
	ctx context.Context,
	q *Client,
	method string,
	relUrl string,
	reqBody interface{},
) (*O, error) {
	req, err := jsonRequest(ctx, q, method, relUrl, reqBody)
	if err != nil {
		return nil, err
	}
	res, err := q.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return jsonResponse[O](res, q.hardcoreDecode)
}

func apiCommand[O any](
	ctx context.Context,
	q *Client,
	method string,
	relUrl string,
) (*O, error) {
	return apiObject[O](ctx, q, method, relUrl, nil)
}

func jsonResponse[O any](res *http.Response, hardcoreDecode bool) (*O, error) {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		decode := json.NewDecoder(res.Body)
		if hardcoreDecode {
			decode.DisallowUnknownFields()
		}
		result := new(O)
		err := decode.Decode(result)
		if err != nil {
			unmarshalError := new(json.UnmarshalTypeError)
			if errors.As(err, &unmarshalError) {
				unmarshalError.Offset = decode.InputOffset()
			}
			return nil, &APIError{
				RequestURL: res.Request.URL,
				StatusCode: res.StatusCode,
				Status:     res.Status,
				ParseError: err,
			}
		}
		if decode.More() {
			return nil, &APIError{
				RequestURL: res.Request.URL,
				StatusCode: res.StatusCode,
				Status:     res.Status,
				ParseError: fmt.Errorf("Response body has additional unparsed data at position %d", decode.InputOffset()+1),
			}
		}
		return result, err
	}
	return nil, newAPIError(res)
}

func newAPIError(res *http.Response) *APIError {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	apierr := APIError{
		RequestURL: res.Request.URL,
		StatusCode: res.StatusCode,
		Status:     res.Status,
	}
	body, err := io.ReadAll(res.Body)
	if err == nil {
		err = json.Unmarshal(body, &apierr)
	}
	apierr.ParseError = err
	apierr.ResponseBody = body
	return &apierr
}

func jsonRequest(ctx context.Context, q *Client, method string, relURL string, body interface{}) (*http.Request, error) {
	absUrl, err := q.baseURL.Parse(relURL)
	if err != nil {
		return nil, err
	}
	buf, err := jsonReader(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, absUrl.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if q.UserAgent != "" {
		req.Header.Add("User-Agent", q.UserAgent)
	}
	return req, nil
}

func jsonReader(from interface{}) (io.Reader, error) {
	var buf bytes.Buffer
	if from == nil {
		return &buf, nil
	}
	err := json.NewEncoder(&buf).Encode(from)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
