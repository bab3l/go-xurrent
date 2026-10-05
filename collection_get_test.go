package xurrent

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCollectionJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v1/teams", r.URL.Path)
		require.Equal(t, "foo", r.URL.Query().Get("name"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[]`))
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: srv.URL, Description: "t"}}
	cfg.AddDefaultHeader("Authorization", "Bearer t")
	c := NewAPIClient(cfg)
	q := url.Values{}
	q.Set("name", "foo")
	body, resp, err := c.GetCollectionJSON(context.Background(), "/v1/teams", q)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, "[]", string(body))
}
