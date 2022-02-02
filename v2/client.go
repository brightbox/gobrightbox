package brightbox

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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

// APIGet makes a GET request to the API
// and decoding any JSON response.
//
// relUrl is the relative path of the endpoint to the base URL, e.g. "servers".
func APIGet[O any](
	ctx context.Context,
	q *Client,
	relUrl string,
) (*O, error) {
	return apiObject[O](ctx, q, "GET", relUrl, nil)
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

// APIPutCommand makes a PUT request to the API
//
// relUrl is the relative path of the endpoint to the base URL, e.g. "servers".
func APIPutCommand(
	ctx context.Context,
	q *Client,
	relUrl string,
) error {
	return apiCommand(ctx, q, "PUT", relUrl)
}

// APIDelete makes a DELETE request to the API
//
// relUrl is the relative path of the endpoint to the base URL, e.g. "servers".
func APIDelete(
	ctx context.Context,
	q *Client,
	relUrl string,
) error {
	return apiCommand(ctx, q, "DELETE", relUrl)
}

func apiObject[O any](
	ctx context.Context,
	q *Client,
	method string,
	relUrl string,
	reqBody interface{},
) (*O, error) {
	req, err := q.jsonRequest(ctx, method, relUrl, reqBody)
	if err != nil {
		return nil, err
	}
	res, err := q.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return jsonResponse[O](res)
}

func apiCommand(
	ctx context.Context,
	q *Client,
	method string,
	relUrl string,
) error {
	req, err := q.jsonRequest(ctx, method, relUrl, nil)
	if err != nil {
		return err
	}
	res, err := q.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	return newAPIError(res)
}

func jsonResponse[O any](res *http.Response) (*O, error) {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		result := new(O)
		decode := json.NewDecoder(res.Body)
		decode.DisallowUnknownFields()
		err := decode.Decode(result)
		if err != nil {
			return nil, &APIError{
				RequestURL: res.Request.URL,
				StatusCode: res.StatusCode,
				Status:     res.Status,
				ParseError: err,
			}
		}
		return result, err
	}
	return nil, newAPIError(res)
}

func newAPIError(res *http.Response) *APIError {
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

func (q *Client) jsonRequest(ctx context.Context, method string, relURL string, body interface{}) (*http.Request, error) {
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

func jsonReader(from interface{}) (io.Reader, error){
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
