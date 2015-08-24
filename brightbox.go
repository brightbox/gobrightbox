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

type Client struct {
	BaseURL   *url.URL
	client    *http.Client
	UserAgent string
	AccountId string
}

type ApiError struct {
	StatusCode           int
	Status               string
	AuthError            string   `json:"error"`
	AuthErrorDescription string   `json:"error_description"`
	ErrorName            string   `json:"error_name"`
	Errors               []string `json:"errors"`
	ParseError           *error
	RequestUrl           *url.URL
	ResponseBody         *[]byte
}

func (e ApiError) Error() string {
	var url string
	if e.RequestUrl != nil {
		url = e.RequestUrl.String()
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

type Resource struct {
	Id string
}

func NewClient(apiUrl url.URL, accountId *string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{
		client:  httpClient,
		BaseURL: &apiUrl,
	}
	if accountId != nil {
		c.AccountId = *accountId
	}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	if c.AccountId != "" {
		q := u.Query()
		q.Set("account_id", c.AccountId)
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

func (c *Client) MakeApiRequest(method string, path string, reqBody interface{}, resBody interface{}) (*http.Response, error) {
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
				return res, ApiError{
					RequestUrl: res.Request.URL,
					StatusCode: res.StatusCode,
					Status:     res.Status,
					ParseError: &err,
				}
			}
		}
		return res, nil
	} else {
		apierr := ApiError{
			RequestUrl: res.Request.URL,
			StatusCode: res.StatusCode,
			Status:     res.Status,
		}
		body, _ := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &apierr)
		apierr.ResponseBody = &body
		return res, apierr
	}
}
