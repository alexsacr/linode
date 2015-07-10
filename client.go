package linode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// Linode API URL.
	BaseURL = "https://api.linode.com/"
)

type httpPoster func(string, url.Values) (*http.Response, error)
type apiCaller func(string, map[string]interface{}) (json.RawMessage, error)
type argMarshaler func(interface{}) (map[string]interface{}, error)

// Client is the API client.  It should be created by a call to
// NewClient().
type Client struct {
	URL        string
	key        string
	post       httpPoster
	apiCall    apiCaller
	argMarshal argMarshaler
}

// NewClient returns a new client configured with the passed API key.
func NewClient(apiKey string) *Client {
	c := &Client{
		URL:        BaseURL,
		key:        apiKey,
		post:       http.PostForm,
		argMarshal: marshallArgs,
	}
	c.apiCall = c.liveAPICaller
	return c
}

type apiError struct {
	Code int    `json:"ERRORCODE"`
	Msg  string `json:"ERRORMESSAGE"`
}

type apiResponse struct {
	APIErrors []apiError      `json:"ERRORARRAY"`
	Data      json.RawMessage `json:"DATA"`
}

func (c *Client) liveAPICaller(method string, args map[string]interface{}) (json.RawMessage, error) {
	vals := url.Values{}
	vals.Set("api_action", method)
	vals.Set("api_key", c.key)

	for k, t := range args {
		switch v := t.(type) {
		case string:
			vals.Set(k, v)
		case *string:
			if v != nil {
				vals.Set(k, *v)
			}
		case int:
			vals.Set(k, strconv.Itoa(v))
		case *int:
			if v != nil {
				vals.Set(k, strconv.Itoa(*v))
			}
		case bool:
			vals.Set(k, fmt.Sprintf("%t", v))
		case *bool:
			if v != nil {
				vals.Set(k, fmt.Sprintf("%t", *v))
			}
		default:
			return nil, fmt.Errorf("cannot convert %s to string", k)
		}
	}

	resp, err := c.post(c.URL, vals)
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	record(method, vals, body)

	ret := apiResponse{}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	if len(ret.APIErrors) != 0 && ret.APIErrors[0].Code != 0 {
		// Only return a single error for now.
		return nil, fmt.Errorf("api: %d: %s", ret.APIErrors[0].Code, ret.APIErrors[0].Msg)
	}

	return ret.Data, nil
}
