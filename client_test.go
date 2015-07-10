// +build !integration

package linode

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

func TestClientSerialization(t *testing.T) {
	// Just check the paths that aren't exercised during normal usage.

	// Non-nil bool
	capturePost := func(_ string, v url.Values) (*http.Response, error) {
		assert.Equal(t, "true", v.Get("foo"))
		assert.Equal(t, "false", v.Get("bar"))

		return nil, errors.New("bail")
	}

	c := NewClient("")
	c.post = capturePost

	args := make(map[string]interface{})
	args["foo"] = true
	args["bar"] = false

	_, _ = c.apiCall("testing", args)

	// Unsupported
	args = make(map[string]interface{})
	args["foo"] = []struct{}{}

	_, err := c.apiCall("testing", args)
	assert.Error(t, err)
}

type errReadCloser struct{}

func (e errReadCloser) Read(_ []byte) (int, error) {
	return 0, errors.New("foo")
}

func (e errReadCloser) Close() error {
	return nil
}

func TestClientReadError(t *testing.T) {
	bodyErrPost := func(_ string, _ url.Values) (*http.Response, error) {
		return &http.Response{Body: errReadCloser{}}, nil
	}

	c := NewClient("")
	c.post = bodyErrPost

	_, err := c.apiCall("foo", map[string]interface{}{})
	assert.Error(t, err)
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

func TestClientJSONUnmarshalError(t *testing.T) {
	badJSONPost := func(_ string, _ url.Values) (*http.Response, error) {
		return &http.Response{Body: nopCloser{bytes.NewBufferString("<")}}, nil
	}

	c := NewClient("")
	c.post = badJSONPost

	_, err := c.apiCall("foo", map[string]interface{}{})
	assert.Error(t, err)
}

func mockClientAPIError() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[{"ERRORCODE":4,"ERRORMESSAGE":"Authentication failed"}],"DATA":{},"ACTION":"test.echo"}`
	params = map[string]string{}
	responses = append(responses, newMockAPIResponse("test.echo", params, output))

	return responses
}

func TestClientAPIError(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockClientAPIError()))
	defer ts.Close()

	err := c.TestEcho()
	require.Error(t, err)
	assert.Equal(t, "api: 4: Authentication failed", err.Error())
}
