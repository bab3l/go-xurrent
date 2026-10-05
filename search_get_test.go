package xurrent

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSearchJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v1/search", r.URL.Path)
		require.Equal(t, "windows", r.URL.Query().Get("q"))
		require.Equal(t, "tok-next", r.URL.Query().Get("page"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[]`))
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: srv.URL, Description: "t"}}
	cfg.AddDefaultHeader("Authorization", "Bearer t")
	c := NewAPIClient(cfg)
	q := url.Values{}
	q.Set("q", "windows")
	q.Set("per_page", "25")
	q.Set("page", "tok-next")
	body, resp, err := c.GetSearchJSON(context.Background(), q)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, "[]", string(body))
}
