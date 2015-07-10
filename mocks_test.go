package linode

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

type mockAPIResponse struct {
	action string
	output func() string
	params map[string]string
}

func newMockAPIResponse(action string, params map[string]string, output interface{}) mockAPIResponse {
	resp := mockAPIResponse{
		action: action,
		params: params,
	}

	switch v := output.(type) {
	case string:
		resp.output = func() string {
			return v
		}
	case func() string:
		resp.output = v
	default:
		panic("Unknown output type passed to newMockResponse.")
	}

	return resp
}

func newMockAPIServer(t *testing.T, responses []mockAPIResponse) *httptest.Server {
	var reqCount int

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(responses) < reqCount {
			msg := fmt.Sprintf("Request count: %d len(responses): %d", reqCount+1, len(responses))
			require.FailNow(t, msg)
		}

		resp := responses[reqCount]

		action := r.FormValue("api_action")
		require.NotEmpty(t, action, fmt.Sprintf("%d", reqCount+1))

		require.Equal(t, resp.action, action, fmt.Sprintf("%d", reqCount+1))

		for k, v := range resp.params {
			assert.Equal(t, v, r.FormValue(k), fmt.Sprintf("params: %s - %d", k, reqCount+1))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(resp.output()))

		reqCount++
	}))

	return ts
}

func clientFor(ts *httptest.Server) (*Client, *httptest.Server) {
	c := NewClient("foo")
	c.URL = ts.URL
	return c, ts
}
